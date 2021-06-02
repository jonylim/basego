package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// FileRepo manages data operations for files, especially cache operations.
type FileRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	store     *redisstore.FileStore
}

// NewFileRepo returns new instance of FileRepo.
func NewFileRepo(redisConn redis.Conn) *FileRepo {
	return &FileRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewFileStore(redisConn),
	}
}

// RedisStore returns Redis store used by the repository.
func (instance *FileRepo) RedisStore() *redisstore.FileStore {
	return instance.store
}

// GetByID returns a file's details by ID.
func (instance *FileRepo) GetByID(id int64) (model.File, error) {
	// Get from Redis.
	res, err := instance.store.GetByID(id)
	if err != nil {
		// Get from database.
		da := dao.NewFileDAO()
		da.WithDeleted()
		res, err = da.GetByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByID(id)
				return res, instance.ErrNotFound
			}
			return res, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(res)
	}
	if instance.exists(res) {
		return res, nil
	}
	return res, instance.ErrNotFound
}

// GetByOwnerAndCategory returns a file's details by owner type, owner ID, and category.
func (instance *FileRepo) GetByOwnerAndCategory(ownerType string, ownerID int64, category string) (model.File, error) {
	// Get from Redis.
	res, err := instance.store.GetByOwnerAndCategory(ownerType, ownerID, category)
	if err != nil {
		// Get from database.
		da := dao.NewFileDAO()
		da.WithDeleted()
		res, err = da.GetByOwnerAndCategory(ownerType, ownerID, category)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByOwnerAndCategory(ownerType, ownerID, category)
				return res, instance.ErrNotFound
			}
			return res, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(res)
	}
	if instance.exists(res) {
		return res, nil
	}
	return res, instance.ErrNotFound
}

// GetByOwnerTypeAndCategoryAndFilename returns a file's details by owner type, category, and filename.
func (instance *FileRepo) GetByOwnerTypeAndCategoryAndFilename(ownerType, category, filename string) (model.File, error) {
	// Get from Redis.
	res, err := instance.store.GetByOwnerTypeAndCategoryAndFilename(ownerType, category, filename)
	if err != nil {
		// Get from database.
		da := dao.NewFileDAO()
		da.WithDeleted()
		res, err = da.GetByOwnerTypeAndCategoryAndFilename(ownerType, category, filename)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByOwnerTypeAndCategoryAndFilename(ownerType, category, filename)
				return res, instance.ErrNotFound
			}
			return res, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(res)
	}
	if instance.exists(res) {
		return res, nil
	}
	return res, instance.ErrNotFound
}

func (instance *FileRepo) exists(data model.File) bool {
	return !data.RedisNil && data.ID != 0 && data.DeletedTime == 0
}
