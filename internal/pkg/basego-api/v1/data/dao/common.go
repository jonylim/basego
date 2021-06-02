package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/storage"
)

type cstAccountPhoto struct {
	Filename, Storage sql.NullString
	IsEncrypted       sql.NullBool
}

// ImageURL generates fullsize & thumbnail image URLs.
func (img *cstAccountPhoto) ImageURL() model.ImageURL {
	return storage.GenerateCstAccountPhotoURL(img.Filename.String, img.Storage.String, img.IsEncrypted.Bool)
}
