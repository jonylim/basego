package storage

import (
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// GeneratePublicFileURL returns public URL for a file.
func GeneratePublicFileURL(filepath, storage string, isEncrypted bool) (url string) {
	if filepath != "" {
		if isEncrypted {
			// TODO: Not supported yet.
			logger.Warn("storage", "Public file URL for encrypted file is not supported yet")
		} else {
			instance, err := getPrivateInstance(storage)
			if err != nil {
				logger.Error("storage", logger.FromError(err))
				return
			}
			url = instance.GetPublicFileURL(filepath)
		}
	}
	return
}

// GenerateCstAccountPhotoURL returns image URL for a customer account's photo.
func GenerateCstAccountPhotoURL(filename, storage string, isEncrypted bool) model.ImageURL {
	if filename != "" {
		fullFilepath, thumbFilepath := GetCstAccountPhotoFilepath(filename)
		if isEncrypted {
			// TODO: Not supported yet.
			logger.Warn("storage", "Public file URL for encrypted file is not supported yet")
		} else {
			instance, err := getPrivateInstance(storage)
			if err != nil {
				logger.Error("storage", logger.FromError(err))
				return model.ImageURL{}
			}
			return model.ImageURL{
				Fullsize:  instance.GetPublicFileURL(fullFilepath),
				Thumbnail: instance.GetPublicFileURL(thumbFilepath),
			}
		}
	}
	return model.ImageURL{}
}
