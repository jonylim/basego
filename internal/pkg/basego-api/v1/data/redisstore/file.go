package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// FileStore manages Redis operations for files.
type FileStore struct {
	redisStore
	byID                        string
	byOwnerTypeCategoryFilename string
	byOwnerAndCategory          string
	ttl                         int
}

// NewFileStore returns new instance to manage files.
func NewFileStore(conn redis.Conn) *FileStore {
	return &FileStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "file",
		},
		byID:                        "id",
		byOwnerTypeCategoryFilename: "owtCatFile",
		byOwnerAndCategory:          "ownAndCat",
		ttl:                         180 * 86400, // 180 days
	}
}

// GetByID returns a file's details by ID.
func (store *FileStore) GetByID(id int64) (model.File, error) {
	return store.getByKey(store.generateStoreKeyByID(id))
}

// GetByOwnerAndCategory returns a file's details by owner type, owner ID, and category.
func (store *FileStore) GetByOwnerAndCategory(ownerType string, ownerID int64, category string) (model.File, error) {
	return store.getByKey(store.generateStoreKeyByOwnerAndCategory(ownerType, ownerID, category))
}

// GetByOwnerTypeAndCategoryAndFilename returns a file's details by owner type, category, and filename.
func (store *FileStore) GetByOwnerTypeAndCategoryAndFilename(ownerType, category, filename string) (model.File, error) {
	return store.getByKey(store.generateStoreKeyByOwnerTypeCategoryFilename(ownerType, category, filename))
}

func (store *FileStore) getByKey(key string) (model.File, error) {
	var tmp, res model.File
	err := store.DoHGETALL(key, &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("FileStore", logger.FromError(err))
	}
	return res, err
}

// Save saves a file.
func (store *FileStore) Save(item model.File) error {
	// Delete old data.
	store.Delete(item)

	// Save new data.
	keys := store.generateStoreKeys(item)
	for _, key := range keys {
		if err := store.DoHMSET(key, &item, store.ttl); err != nil {
			logger.Fatal("FileStore", logger.FromError(err))
			return err
		}
	}
	return nil
}

// SaveNilByID saves an empty file's details by ID.
func (store *FileStore) SaveNilByID(id int64) error {
	return store.saveNilByKey(store.generateStoreKeyByID(id))
}

// SaveNilByOwnerAndCategory saves an empty file's details by owner type, owner ID, and category.
func (store *FileStore) SaveNilByOwnerAndCategory(ownerType string, ownerID int64, category string) error {
	return store.saveNilByKey(store.generateStoreKeyByOwnerAndCategory(ownerType, ownerID, category))
}

// SaveNilByOwnerTypeAndCategoryAndFilename saves an empty file's details by owner type, category, and filename.
func (store *FileStore) SaveNilByOwnerTypeAndCategoryAndFilename(ownerType, category, filename string) error {
	return store.saveNilByKey(store.generateStoreKeyByOwnerTypeCategoryFilename(ownerType, category, filename))
}

func (store *FileStore) saveNilByKey(key string) error {
	if err := store.DoHMSET(key, emptyItem, store.ttl); err != nil {
		logger.Fatal("FileStore", logger.FromError(err))
		return err
	}
	return nil
}

// Delete deletes a file's details by all possible keys.
func (store *FileStore) Delete(item model.File) (bool, error) {
	checkKeys := store.generateStoreKeys(item)
	deleteKeys := make([]string, 0, len(checkKeys))
	deleteKeys = append(deleteKeys, checkKeys...)
	for _, key := range checkKeys {
		if existing, err := store.getByKey(key); err == nil {
			existingKeys := store.generateStoreKeys(existing)
			deleteKeys = append(deleteKeys, existingKeys...)
		}
	}
	if len(deleteKeys) == 0 {
		return false, nil
	}
	count, err := store.DoDEL(deleteKeys...)
	if err != nil && err != redis.ErrNil {
		logger.Fatal("FileStore", logger.FromError(err))
		return false, err
	}
	return count != 0, nil
}

// DeleteByID deletes a file's details by ID.
func (store *FileStore) DeleteByID(id int64) (bool, error) {
	return store.deleteByKey(store.generateStoreKeyByID(id))
}

// DeleteByOwnerAndCategory deletes a file's details by owner type, owner ID, and category.
func (store *FileStore) DeleteByOwnerAndCategory(ownerType string, ownerID int64, category string) (bool, error) {
	return store.deleteByKey(store.generateStoreKeyByOwnerAndCategory(ownerType, ownerID, category))
}

func (store *FileStore) deleteByKey(key string) (deleted bool, err error) {
	var count int
	if existing, err1 := store.getByKey(key); err1 == nil {
		deleteKeys := store.generateStoreKeys(existing)
		deleteKeys = append(deleteKeys, key)
		count, err = store.DoDEL(deleteKeys...)
	} else {
		count, err = store.DoDEL(key)
	}
	if err != nil && err != redis.ErrNil {
		logger.Fatal("FileStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

func (store *FileStore) generateStoreKeys(item model.File) []string {
	keys := make([]string, 0)
	if item.ID != 0 {
		keys = append(keys, store.generateStoreKeyByID(item.ID))
	}
	if item.OwnerType != "" && item.OwnerID != 0 && item.Category != "" {
		keys = append(keys, store.generateStoreKeyByOwnerAndCategory(item.OwnerType, item.OwnerID, item.Category))
	}
	if item.OwnerType != "" && item.Category != "" && item.Filename != "" {
		keys = append(keys, store.generateStoreKeyByOwnerTypeCategoryFilename(item.OwnerType, item.Category, item.Filename))
	}
	return keys
}

func (store *FileStore) generateStoreKeyByID(id int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, id)
}

func (store *FileStore) generateStoreKeyByOwnerAndCategory(ownerType string, ownerID int64, category string) string {
	return fmt.Sprintf("%s:%s:%s:%v:%s", store.baseKey, store.byOwnerAndCategory, ownerType, ownerID, category)
}

func (store *FileStore) generateStoreKeyByOwnerTypeCategoryFilename(ownerType, category, filename string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", store.baseKey, store.byOwnerTypeCategoryFilename, ownerType, category, filename)
}
