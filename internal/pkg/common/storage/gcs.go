package storage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"cloud.google.com/go/storage"
	context "golang.org/x/net/context"
)

// GoogleCloudStorageConfig defines configurations for using Google Cloud Storage.
type GoogleCloudStorageConfig struct {
	ProjectID, BucketName, StorageClass, Location string
}

type tGoogleCloudStorage struct {
	projectID, bucketName, storageClass, bucketLocation string
	makePublic                                          bool
}

var gcsConfig GoogleCloudStorageConfig
var gcsInstance *tGoogleCloudStorage
var errGCS = errors.New("Storage is not initialized")

func initGoogleCloudStorageConfig() {
	gcsConfig = GoogleCloudStorageConfig{
		ProjectID:    os.Getenv(envvar.Google.ProjectID),
		BucketName:   os.Getenv(envvar.Google.StorageBucket),
		StorageClass: os.Getenv(envvar.Google.StorageClass),
		Location:     os.Getenv(envvar.Google.StorageLocation),
	}

	// Validate the configs.
	if gcsConfig.ProjectID == "" {
		logger.Println("storage", "WARN: GoogleCloudStorageConfig.ProjectID is empty")
	} else {
		logger.Println("storage", fmt.Sprintf("GoogleCloudStorageConfig.ProjectID set to '%s'", gcsConfig.ProjectID))
	}
	if gcsConfig.BucketName == "" {
		logger.Println("storage", "WARN: GoogleCloudStorageConfig.BucketName is empty")
	} else {
		logger.Println("storage", fmt.Sprintf("GoogleCloudStorageConfig.BucketName set to '%s'", gcsConfig.BucketName))
	}
	if gcsConfig.StorageClass == "" {
		gcsConfig.StorageClass = "REGIONAL"
		logger.Println("storage", fmt.Sprintf("WARN: GoogleCloudStorageConfig.StorageClass is empty, set to '%v' as default", gcsConfig.StorageClass))
	} else {
		logger.Println("storage", fmt.Sprintf("GoogleCloudStorageConfig.StorageClass set to '%s'", gcsConfig.StorageClass))
	}
	if gcsConfig.Location == "" {
		logger.Println("storage", "WARN: GoogleCloudStorageConfig.Location is empty")
	} else {
		logger.Println("storage", fmt.Sprintf("GoogleCloudStorageConfig.Location set to '%s'", gcsConfig.Location))
	}

	// Set default error.
	if gcsConfig.ProjectID == "" || gcsConfig.BucketName == "" || gcsConfig.Location == "" {
		errGCS = errors.New("Storage is not configured correctly")
	} else if os.Getenv(envvar.Google.AppCredentials) == "" {
		errGCS = errors.New("GOOGLE_APPLICATION_CREDENTIALS is undefined")
	} else {
		errGCS = nil
		gcsInstance, _ = newGoogleCloudStorage()
	}
}

func newGoogleCloudStorage() (*tGoogleCloudStorage, error) {
	if errGCS != nil {
		return nil, errGCS
	}
	return &tGoogleCloudStorage{
		projectID:      gcsConfig.ProjectID,
		bucketName:     gcsConfig.BucketName,
		storageClass:   gcsConfig.StorageClass,
		bucketLocation: gcsConfig.Location,
		makePublic:     false,
	}, nil
}

func getGoogleCloudStorageInstance() (*tGoogleCloudStorage, error) {
	return gcsInstance, errGCS
}

// GetPublicFileURL creates public file URL of the specified filepath.
func (instance *tGoogleCloudStorage) GetPublicFileURL(objFilepath string) string {
	return instance.getPublicFileURLInBucket(instance.bucketName, objFilepath)
}

// FetchObject creates a new reader to read the contents of the object specified by the filepath.
// The caller must call Close on the returned reader when done reading.
func (instance *tGoogleCloudStorage) FetchObject(srcFilepath string) (reader io.ReadCloser, code int, err error) {
	return instance.fetchObjectInBucket(instance.bucketName, srcFilepath)
}

// DeleteObject deletes the object specified by the filepath.
func (instance *tGoogleCloudStorage) DeleteObject(srcFilepath string) (code int, err error) {
	return instance.deleteObjectInBucket(instance.bucketName, srcFilepath)
}

// StoreObject saves a file to the destination filepath in a Google Cloud Storage's bucket.
// The `path` should not starts with file separators or dots.
func (instance *tGoogleCloudStorage) StoreObject(dstFilepath string, src io.Reader, contentType string) error {
	return instance.storeObjectInBucket(instance.bucketName, dstFilepath, src, contentType)
}

// ShouldMakePublic makes the next files stored to the storage either public or not. Only supported for certain storages.
func (instance *tGoogleCloudStorage) ShouldMakePublic(makePublic bool) {
	instance.makePublic = true
}

func (instance *tGoogleCloudStorage) getPublicFileURLInBucket(bucketName, objFilepath string) string {
	objFilepath = filepath.ToSlash(filepath.Clean(trimStartingSlashes(objFilepath)))
	if objFilepath == "" {
		return ""
	} else if objFilepath[:1] == "." {
		return ""
	}
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objFilepath)
}

func (instance *tGoogleCloudStorage) fetchObjectInBucket(bucketName, srcFilepath string) (reader io.ReadCloser, code int, err error) {
	srcFilepath = filepath.ToSlash(filepath.Clean(trimStartingSlashes(srcFilepath)))
	if srcFilepath == "" {
		code = ErrOther
		err = errors.New("Path can't be empty")
		return
	} else if srcFilepath[:1] == "." {
		code = ErrOther
		err = errors.New("Path can't start with dot")
		return
	}

	// Create a client.
	ctx := context.Background()
	client, errClient := storage.NewClient(ctx)
	if errClient != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		code = ErrOther
		err = errClient
		return
	}

	// Create object handle to the destination filepath.
	objHandle := client.Bucket(bucketName).Object(srcFilepath)

	// Get reader to the object.
	reader, err = objHandle.NewReader(ctx)
	if err != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		if err == storage.ErrObjectNotExist {
			code = ErrNotFound
		} else {
			code = ErrOther
		}
		return
	}
	return
}

func (instance *tGoogleCloudStorage) deleteObjectInBucket(bucketName, srcFilepath string) (code int, err error) {
	srcFilepath = filepath.ToSlash(filepath.Clean(trimStartingSlashes(srcFilepath)))
	if srcFilepath == "" {
		code = ErrOther
		err = errors.New("Path can't be empty")
		return
	} else if srcFilepath[:1] == "." {
		code = ErrOther
		err = errors.New("Path can't start with dot")
		return
	}

	// Create a client.
	ctx := context.Background()
	client, errClient := storage.NewClient(ctx)
	if errClient != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		code = ErrOther
		err = errClient
		return
	}

	// Create object handle to the destination filepath.
	objHandle := client.Bucket(bucketName).Object(srcFilepath)

	// Delete the object.
	err = objHandle.Delete(ctx)
	if err != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		if err == storage.ErrObjectNotExist {
			code = ErrNotFound
		} else {
			code = ErrOther
		}
		return
	}
	return
}

func (instance *tGoogleCloudStorage) storeObjectInBucket(bucketName, dstFilepath string, src io.Reader, contentType string) error {
	dstFilepath = filepath.ToSlash(filepath.Clean(trimStartingSlashes(dstFilepath)))
	if dstFilepath == "" {
		return errors.New("Path can't be empty")
	} else if dstFilepath[:1] == "." {
		return errors.New("Path can't start with dot")
	}

	// Create a client.
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		return err
	}

	// Create bucket handle.
	bucketHandle := client.Bucket(bucketName)

	// Check if the bucket already exists.
	if _, err := bucketHandle.Attrs(ctx); err != nil {
		if err != storage.ErrBucketNotExist {
			// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
			return err
		}

		// Create the bucket.
		bucketAttrs := &storage.BucketAttrs{
			StorageClass: instance.storageClass,
			Location:     instance.bucketLocation,
		}
		if err := bucketHandle.Create(ctx, instance.projectID, bucketAttrs); err != nil {
			// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
			return err
		}
	}

	// Create object handle to the destination filepath.
	objHandle := bucketHandle.Object(dstFilepath)

	// Get writer to the object.
	writer := objHandle.NewWriter(ctx)

	// Write the file the destination filepath.
	if _, err = io.Copy(writer, src); err != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		return err
	}

	// Close the writer.
	if err = writer.Close(); err != nil {
		// logger.Fatal("tGoogleCloudStorage", logger.FromError(err))
		return err
	}

	// Update the object attributes.
	objHandle.Update(ctx, storage.ObjectAttrsToUpdate{
		ContentType: contentType,
	})

	// Make public.
	if instance.makePublic {
		acl := objHandle.ACL()
		if err = acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			// Delete the object.
			objHandle.Delete(ctx)
		}
	}

	return nil
}
