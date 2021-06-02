package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// FileDAO manages database operations for uploaded file data.
type FileDAO struct {
	dao
	selectColumns string
}

// NewFileDAO returns new instance of FileDAO.
func NewFileDAO() *FileDAO {
	return &FileDAO{
		dao: dao{db.Get(), false},
		selectColumns: `
				id, owner_type, owner_id, category, filename, original_filename,
				media_type, file_ext, file_size, width, height,
				thumbnail_media_type, thumbnail_file_ext, thumbnail_file_size, thumbnail_width, thumbnail_height,
				storage, is_encrypted, encrypt_key, uploader,
				` + sqlTimestampToUnixMilliseconds("created_at") + ` AS created_time,
				` + sqlTimestampToUnixMilliseconds("updated_at") + ` AS updated_time,
				` + sqlTimestampToUnixMilliseconds("deleted_at") + ` AS deleted_time`,
	}
}

func (instance *FileDAO) scanRow(row *sql.Row) (res model.File, err error) {
	err = row.Scan(
		&res.ID, &res.OwnerType, &res.OwnerID, &res.Category, &res.Filename, &res.OriginalFilename,
		&res.MediaType, &res.FileExt, &res.FileSize, &res.Width, &res.Height,
		&res.ThumbMediaType, &res.ThumbFileExt, &res.ThumbFileSize, &res.ThumbWidth, &res.ThumbHeight,
		&res.Storage, &res.IsEncrypted, &res.EncryptKey, &res.Uploader,
		&res.CreatedTime, &res.UpdatedTime, &res.DeletedTime)
	return
}

func (instance *FileDAO) getWhere(sqlWhere string, params ...interface{}) (res model.File, err error) {
	row := instance.db.QueryRow(`SELECT `+instance.selectColumns+`
			FROM tb_m_file
			`+sqlWhere, params...)
	res, err = instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("FileDAO", logger.FromError(err))
	}
	return
}

// GetByID returns a file's details by ID.
func (instance *FileDAO) GetByID(id int64) (model.File, error) {
	where := `WHERE id = $1 `
	if !instance.withDeleted {
		where += `AND deleted_at IS NULL `
	}
	return instance.getWhere(where, id)
}

// GetByOwnerAndCategory returns a file's details by owner type, owner ID, and category.
func (instance *FileDAO) GetByOwnerAndCategory(ownerType string, ownerID int64, category string) (model.File, error) {
	where := `WHERE owner_type = $1
				AND owner_id = $2
				AND category = $3 `
	if !instance.withDeleted {
		where += `
				AND deleted_at IS NULL `
	}
	return instance.getWhere(where, ownerType, ownerID, category)
}

// GetByOwnerTypeAndCategoryAndFilename returns a file's details by owner type, category, and filename.
func (instance *FileDAO) GetByOwnerTypeAndCategoryAndFilename(ownerType, category, filename string) (model.File, error) {
	where := `WHERE owner_type = $1
				AND category = $2
				AND filename = $3 `
	if !instance.withDeleted {
		where += `
				AND deleted_at IS NULL `
	}
	return instance.getWhere(where, ownerType, category, filename)
}

func (instance *FileDAO) existsWhere(tx *sql.Tx, sqlWhere string, params ...interface{}) (bool, error) {
	rows, err := tx.Query(`SELECT 1 AS exists FROM tb_m_file `+sqlWhere, params...)
	if err != nil {
		logger.Fatal("FileDAO", logger.FromError(err))
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, rows.Err()
}

// ExistsByOwnerAndCategory checks if a file exists by owner type, owner ID, and category.
func (instance *FileDAO) ExistsByOwnerAndCategory(tx *sql.Tx, ownerType string, ownerID int64, category string) (bool, error) {
	return instance.existsWhere(tx, `
			WHERE owner_type = $1
				AND owner_id = $2
				AND category = $3
				AND deleted_at IS NULL
		`, ownerType, ownerID, category)
}

// ExistsByOwnerTypeAndCategoryAndFilename checks if a file exists by owner type, category, and filename.
func (instance *FileDAO) ExistsByOwnerTypeAndCategoryAndFilename(tx *sql.Tx, ownerType, category, filename string) (bool, error) {
	return instance.existsWhere(tx, `
			WHERE owner_type = $1
				AND category = $2
				AND filename = $3
				AND deleted_at IS NULL
		`, ownerType, category, filename)
}

// Insert inserts new record of a file to database. This method requires database transaction to be passed.
func (instance *FileDAO) Insert(tx *sql.Tx, item model.File) (id, createdMillis int64, err error) {
	err = tx.QueryRow(`INSERT INTO tb_m_file (
				owner_type, owner_id, category, filename, original_filename,
				media_type, file_ext, file_size, width, height,
				thumbnail_media_type, thumbnail_file_ext, thumbnail_file_size, thumbnail_width, thumbnail_height,
				storage, is_encrypted, encrypt_key, uploader
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
			RETURNING id, `+sqlTimestampToUnixMilliseconds("created_at"),
		item.OwnerType, item.OwnerID, item.Category, item.Filename, item.OriginalFilename,
		item.MediaType, item.FileExt, item.FileSize, item.Width, item.Height,
		item.ThumbMediaType, item.ThumbFileExt, item.ThumbFileSize, item.ThumbWidth, item.ThumbHeight,
		item.Storage, item.IsEncrypted, item.EncryptKey, item.Uploader,
	).Scan(&id, &createdMillis)
	if err != nil {
		logger.Fatal("FileDAO", logger.FromError(err))
	}
	return
}

func (instance *FileDAO) deleteFileWhere(tx *sql.Tx, sqlWhere string, params ...interface{}) (deletedFile model.File, err error) {
	row := tx.QueryRow(`UPDATE tb_m_file
			SET deleted_at = CURRENT_TIMESTAMP
			`+sqlWhere+`
			RETURNING `+instance.selectColumns,
		params...)
	deletedFile, err = instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("FileDAO", logger.FromError(err))
	}
	return
}

// DeleteFileByOwnerAndCategory deletes a file's details by owner type, owner ID, and category.
// If successful, returns the deleted file's details.
func (instance *FileDAO) DeleteFileByOwnerAndCategory(tx *sql.Tx, ownerType string, ownerID int64, category string) (model.File, error) {
	return instance.deleteFileWhere(tx, `
			WHERE owner_type = $1
				AND owner_id = $2
				AND category = $3
				AND deleted_at IS NULL
		`, ownerType, ownerID, category)
}

// DeleteFileByOwnerTypeAndCategoryAndFilename deletes a file's details by owner type, category, and filename.
// If successful, returns the deleted file's details.
func (instance *FileDAO) DeleteFileByOwnerTypeAndCategoryAndFilename(tx *sql.Tx, ownerType, category, filename string) (model.File, error) {
	return instance.deleteFileWhere(tx, `
			WHERE owner_type = $1
				AND category = $2
				AND filename = $3
				AND deleted_at IS NULL
		`, ownerType, category, filename)
}
