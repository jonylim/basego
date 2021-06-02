package storage

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// Storage defines interface for object storage operations.
type Storage interface {
	// GetPublicFileURL creates public file URL of the specified filepath.
	GetPublicFileURL(objFilepath string) string

	// FetchObject creates a new reader to read the contents of the object specified by the filepath.
	// The caller must call Close on the returned reader when done reading.
	FetchObject(srcFilepath string) (reader io.ReadCloser, code int, err error)

	// DeleteObject deletes the object specified by the filepath.
	DeleteObject(srcFilepath string) (code int, err error)

	// StoreObject saves a file to the destination filepath in the selected storage.
	// The `path` should not starts with file separators or dots.
	StoreObject(dstFilepath string, src io.Reader, contentType string) error

	// ShouldMakePublic makes the next files stored to the storage either public or not. Only supported for certain storages.
	ShouldMakePublic(makePublic bool)
}

// Constants for available storage types.
const (
	LocalStorage       = "local"
	GoogleCloudStorage = "gcs"
)

// Constants for error code.
const (
	ErrNotFound = 1
	ErrOther    = 9
)

var defaultStorage string
var isInitialized = false

// Init initializes storage configurations.
func Init() {
	// Set the default storage type.
	SetDefaultStorage(os.Getenv(envvar.Storage.DefaultStorage))

	// Initializes storage configurations.
	initLocalStorageConfig()
	initGoogleCloudStorageConfig()

	// Mark as initialized.
	isInitialized = true
}

// SetDefaultStorage sets default storage type.
func SetDefaultStorage(storage string) {
	if storage == "" {
		logger.Println("storage", "ERROR: Default storage is undefined")
		os.Exit(1)
	} else if storage != LocalStorage && storage != GoogleCloudStorage {
		logger.Println("storage", fmt.Sprintf("ERROR: Default storage '%v' is invalid", storage))
		os.Exit(1)
	} else {
		logger.Println("storage", fmt.Sprintf("Default storage set to '%v'", storage))
	}
	defaultStorage = storage
}

// GetDefaultStorage returns default storage type.
func GetDefaultStorage() string {
	return defaultStorage
}

// GetStorageInstance returns storage instance of specific type.
func GetStorageInstance(storage string) (Storage, error) {
	if !isInitialized {
		return nil, errors.New("Storage is not initialized")
	}
	switch storage {
	case GoogleCloudStorage:
		return newGoogleCloudStorage()
	case LocalStorage:
		return newLocalStorage()
	default:
		return nil, fmt.Errorf("Storage '%v' is invalid", storage)
	}
}

// DeleteObjects deletes objects from the specified storage.
// The returned map contains error for each filepath.
func DeleteObjects(storage string, filepaths ...string) map[string]error {
	errs := make(map[string]error, 0)
	st, err := GetStorageInstance(storage)
	if err != nil {
		for _, f := range filepaths {
			errs[f] = err
		}
	} else {
		for _, f := range filepaths {
			_, err = st.DeleteObject(f)
			errs[f] = err
		}
	}
	return errs
}

func getPrivateInstance(storage string) (Storage, error) {
	switch storage {
	case GoogleCloudStorage:
		return getGoogleCloudStorageInstance()
	case LocalStorage:
		return getLocalStorageInstance()
	default:
		return nil, fmt.Errorf("Storage '%v' is invalid", storage)
	}
}

func trimStartingSlashes(s string) string {
	if s != "" {
		for s[0] == '/' || s[0] == '\\' {
			s = s[1:]
		}
	}
	return s
}
