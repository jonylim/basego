package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// CstAccountTOSStore manages Redis operations for customer account's Terms of Service status.
type CstAccountTOSStore struct {
	redisStore
	ttl     int
	byID    string
	byAccID string
}

// NewCstAccountTOSStore returns new instance of CstAccountTOSStore.
func NewCstAccountTOSStore(conn redis.Conn) *CstAccountTOSStore {
	return &CstAccountTOSStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "cstAccTOS",
		},
		ttl:     86400, // 1 day
		byID:    "id",
		byAccID: "accID",
	}
}

// GetByID returns a customer account's TOS status by ID.
func (store *CstAccountTOSStore) GetByID(id int64) (model.CstAccountTOS, error) {
	return store.getByKey(store.generateStoreKeyByID(id))
}

// GetByAccountID returns a customer account's TOS status by account ID.
func (store *CstAccountTOSStore) GetByAccountID(accountID int64) (model.CstAccountTOS, error) {
	return store.getByKey(store.generateStoreKeyByAccountID(accountID))
}

// Save saves a customer account's TOS status.
func (store *CstAccountTOSStore) Save(tos model.CstAccountTOS) error {
	// Delete old data.
	store.Delete(tos)

	// Save new data.
	keys := store.generateStoreKeys(tos)
	for _, key := range keys {
		if err := store.DoHMSET(key, &tos, store.ttl); err != nil {
			logger.Fatal("CstAccountTOSStore", logger.FromError(err))
			return err
		}
	}
	return nil
}

// SaveNilByID saves an empty TOS status for the ID.
func (store *CstAccountTOSStore) SaveNilByID(id int64) error {
	if err := store.DoHMSET(store.generateStoreKeyByID(id), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountTOSStore", logger.FromError(err))
		return err
	}
	return nil
}

// SaveNilByAccountID saves an empty TOS status for the account ID.
func (store *CstAccountTOSStore) SaveNilByAccountID(accountID int64) error {
	if err := store.DoHMSET(store.generateStoreKeyByAccountID(accountID), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountTOSStore", logger.FromError(err))
		return err
	}
	return nil
}

// Delete deletes a customer account's TOS status by all possible keys.
func (store *CstAccountTOSStore) Delete(tos model.CstAccountTOS) (bool, error) {
	checkKeys := store.generateStoreKeys(tos)
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
		logger.Fatal("CstAccountTOSStore", logger.FromError(err))
		return false, err
	}
	return count != 0, nil
}

// DeleteByID deletes a customer account's TOS status by account ID.
func (store *CstAccountTOSStore) DeleteByAccountID(accountID int64) (deleted bool, err error) {
	var count int
	key := store.generateStoreKeyByAccountID(accountID)
	if existing, err1 := store.getByKey(key); err1 == nil {
		deleteKeys := store.generateStoreKeys(existing)
		deleteKeys = append(deleteKeys, key)
		count, err = store.DoDEL(deleteKeys...)
	} else {
		count, err = store.DoDEL(key)
	}
	if err != nil && err != redis.ErrNil {
		logger.Fatal("CstAccountTOSStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

func (store *CstAccountTOSStore) getByKey(key string) (res model.CstAccountTOS, err error) {
	err = store.DoHGETALL(key, &res)
	if err != nil && err != redis.ErrNil {
		logger.Error("CstAccountTOSStore", logger.FromError(err))
	}
	return
}

func (store *CstAccountTOSStore) generateStoreKeys(tos model.CstAccountTOS) []string {
	return []string{
		store.generateStoreKeyByID(tos.ID),
		store.generateStoreKeyByAccountID(tos.AccountID),
	}
}

func (store *CstAccountTOSStore) generateStoreKeyByID(id int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, id)
}

func (store *CstAccountTOSStore) generateStoreKeyByAccountID(accountID int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byAccID, accountID)
}
