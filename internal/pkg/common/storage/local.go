package storage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// LocalStorageConfig defines configurations for using local file storage.
type LocalStorageConfig struct {
	BaseDirPath string
}

type tLocalStorage struct {
	dirPath string
}

var localConfig LocalStorageConfig
var localInstance *tLocalStorage
var errLocal = errors.New("Storage is not initialized")

func initLocalStorageConfig() {
	localConfig = LocalStorageConfig{
		BaseDirPath: os.Getenv(envvar.Storage.LocalDirPath),
	}

	// Validate the configs.
	if localConfig.BaseDirPath == "" {
		logger.Println("storage", "WARN: LocalStorageConfig.BaseDirPath is empty")
	} else {
		if p, err := filepath.Abs(localConfig.BaseDirPath); err == nil {
			localConfig.BaseDirPath = p
		} else {
			localConfig.BaseDirPath = filepath.Clean(localConfig.BaseDirPath)
		}
		logger.Println("storage", fmt.Sprintf("LocalStorageConfig.BaseDirPath set to '%s'", localConfig.BaseDirPath))
	}

	// Set default error.
	if localConfig.BaseDirPath == "" {
		errLocal = errors.New("Storage is not configured correctly")
	} else {
		errLocal = nil
		localInstance, _ = newLocalStorage()
	}
}

func newLocalStorage() (*tLocalStorage, error) {
	if errLocal != nil {
		return nil, errLocal
	}
	return &tLocalStorage{
		dirPath: localConfig.BaseDirPath,
	}, nil
}

func getLocalStorageInstance() (*tLocalStorage, error) {
	return localInstance, errGCS
}

func (instance *tLocalStorage) completeFilepath(dir, file string) (fullpath string, code int, err error) {
	file = trimStartingSlashes(file)
	if file == "" {
		code = ErrOther
		err = errors.New("Path can't be empty")
	} else if file[:1] == "." {
		code = ErrOther
		err = errors.New("Path can't start with dot")
	} else {
		fullpath = filepath.Clean(dir + "/" + file)
	}
	return
}

// GetPublicFileURL creates public file URL of the specified filepath.
func (instance *tLocalStorage) GetPublicFileURL(objFilepath string) string {
	if _, _, err := instance.completeFilepath("", objFilepath); err != nil {
		return ""
	}
	// TODO: Set public file URL.
	logger.Warn("storage", "Public file URL for local storage is not supported yet")
	return ""
}

// FetchObject creates a new reader to read the contents of the object specified by the filepath.
// The caller must call Close on the returned reader when done reading.
func (instance *tLocalStorage) FetchObject(srcFilepath string) (reader io.ReadCloser, code int, err error) {
	return instance.fetchObjectInDir(instance.dirPath, srcFilepath)
}

// DeleteObject deletes the object specified by the filepath.
func (instance *tLocalStorage) DeleteObject(srcFilepath string) (code int, err error) {
	return instance.deleteObjectInDir(instance.dirPath, srcFilepath)
}

// StoreObject saves a file to the destination filepath in local storage.
// The `path` should not starts with file separators or dots.
func (instance *tLocalStorage) StoreObject(dstFilepath string, src io.Reader, contentType string) error {
	return instance.storeObjectInDir(instance.dirPath, dstFilepath, src, contentType)
}

// ShouldMakePublic do nothing in local storage.
func (instance *tLocalStorage) ShouldMakePublic(makePublic bool) {
}

func (instance *tLocalStorage) fetchObjectInDir(dirPath, srcFilepath string) (reader io.ReadCloser, code int, err error) {
	// Create the complete filepath.
	srcFilepath, code, err = instance.completeFilepath(dirPath, srcFilepath)
	if err != nil {
		return
	}

	// Open the file specified by the filepath.
	reader, err = os.Open(srcFilepath)
	if err != nil {
		// logger.Error("tLocalStorage", logger.FromError(err))
		if os.IsNotExist(err) {
			code = ErrNotFound
		} else {
			code = ErrOther
		}
	}
	return
}

func (instance *tLocalStorage) deleteObjectInDir(dirPath, srcFilepath string) (code int, err error) {
	// Create the complete filepath.
	srcFilepath, code, err = instance.completeFilepath(dirPath, srcFilepath)
	if err != nil {
		return
	}

	// Remove the file specified by the filepath.
	err = os.Remove(srcFilepath)
	if err != nil {
		// logger.Error("tLocalStorage", logger.FromError(err))
		if os.IsNotExist(err) {
			code = ErrNotFound
		} else {
			code = ErrOther
		}
	}
	return
}

func (instance *tLocalStorage) storeObjectInDir(dirPath, dstFilepath string, src io.Reader, contentType string) error {
	// Create the complete filepath.
	var err error
	dstFilepath, _, err = instance.completeFilepath(dirPath, dstFilepath)
	if err != nil {
		return err
	}

	// Check if the directory already exists.
	dir := filepath.Dir(dstFilepath)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	/*
		// Write the file the destination filepath
		data, err := ioutil.ReadAll(src)
		if err != nil {
			// logger.Error("tLocalStorage", logger.FromError(err))
			return err
		}
		err = ioutil.WriteFile(dstFilepath, data, 0666)
		if err != nil {
			// logger.Error("tLocalStorage", logger.FromError(err))
			return err
		}
	*/

	// Open the destination filepath.
	f, err := os.OpenFile(dstFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		// logger.Error("tLocalStorage", logger.FromError(err))
		return err
	}

	// Write the file the destination filepath.
	if _, err = io.Copy(f, src); err != nil {
		// logger.Fatal("tLocalStorage", logger.FromError(err))
		return err
	}

	// Close the destination file.
	if err = f.Close(); err != nil {
		// logger.Fatal("tLocalStorage", logger.FromError(err))
		return err
	}

	return nil
}
