package storage

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/common/constant"
	"github.com/jonylim/basego/internal/pkg/common/storage/basedir"
)

// GetCstAccountPhotoFilepath returns fullsize & thumbnail filepaths for a customer account's photo.
func GetCstAccountPhotoFilepath(filename string) (fullFilepath, thumbFilepath string) {
	dir := basedir.CstAccount(constant.FileCategoryPhoto)
	fullFilepath = fmt.Sprintf("%s/%s-%s", dir, filename, "full")
	thumbFilepath = fmt.Sprintf("%s/%s-%s", dir, filename, "thumb")
	return
}
