package redisstore

import (
	"fmt"
	"time"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// CstAccountOTPStore manages Redis operations for customer account OTP data.
type CstAccountOTPStore struct {
	redisStore
	byID               string
	byAccountAndAction string
}

// NewCstAccountOTPStore returns new instance to manage customer account OTP data.
func NewCstAccountOTPStore(conn redis.Conn) *CstAccountOTPStore {
	return &CstAccountOTPStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "cstAccOTP",
		},
		byID:               "id",
		byAccountAndAction: "accAct",
	}
}

// GetOTPByID returns an OTP's details by OTP ID.
func (store *CstAccountOTPStore) GetOTPByID(id int64) (model.CstAccountOTP, error) {
	return store.getOTPByKey(store.generateStoreKeyByID(id))
}

// GetActiveOTPByAccountAndAction returns an active OTP's details by account ID and action.
func (store *CstAccountOTPStore) GetActiveOTPByAccountAndAction(accountID int64, action string) (model.CstAccountOTP, error) {
	data, err := store.getOTPByKey(store.generateStoreKeyByAccountAndAction(accountID, action))
	if err == nil {
		if data.ExpiryTime <= helper.UnixMillisecond(time.Now()) || data.IsVerified {
			return model.CstAccountOTP{
				RedisNil: true,
			}, nil
		}
	}
	return data, err
}

// SaveOTPByID saves an OTP's details by OTP ID.
func (store *CstAccountOTPStore) SaveOTPByID(data model.CstAccountOTP, ttlSeconds int) error {
	return store.saveOTPByKey(store.generateStoreKeyByID(data.ID), data, ttlSeconds)
}

// SaveOTPByAccountAndAction saves an OTP's details by account ID and action.
func (store *CstAccountOTPStore) SaveOTPByAccountAndAction(data model.CstAccountOTP, ttlSeconds int) error {
	store.saveOTPByKey(store.generateStoreKeyByID(data.ID), data, ttlSeconds)
	return store.saveOTPByKey(store.generateStoreKeyByAccountAndAction(data.AccountID, data.Action), data, ttlSeconds)
}

// SaveNilByID saves an empty OTP's details by OTP ID.
func (store *CstAccountOTPStore) SaveNilByID(id int64, ttlSeconds int) error {
	return store.saveNilByKey(store.generateStoreKeyByID(id), ttlSeconds)
}

// SaveNilByAccountAndAction saves an empty OTP's details by account ID and action.
func (store *CstAccountOTPStore) SaveNilByAccountAndAction(accountID int64, action string, ttlSeconds int) error {
	return store.saveNilByKey(store.generateStoreKeyByAccountAndAction(accountID, action), ttlSeconds)
}

// DeleteOTPByID deletes an OTP's details by OTP ID.
func (store *CstAccountOTPStore) DeleteOTPByID(id int64) (deleted bool, err error) {
	var count int
	count, err = store.DoDEL(store.generateStoreKeyByID(id))
	if err != nil && err != redis.ErrNil {
		logger.Error("CstAccountOTPStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

// DeleteOTPByAccountAndAction deletes an OTP's details by account ID and action.
func (store *CstAccountOTPStore) DeleteOTPByAccountAndAction(accountID int64, action string) (deleted bool, err error) {
	key := store.generateStoreKeyByAccountAndAction(accountID, action)
	deleteKeys := make([]string, 0, 2)
	deleteKeys = append(deleteKeys, key)
	if existing, err1 := store.getOTPByKey(key); err1 == nil {
		deleteKeys = append(deleteKeys, store.generateStoreKeyByID(existing.ID))
	}
	var count int
	count, err = store.DoDEL(deleteKeys...)
	if err != nil && err != redis.ErrNil {
		logger.Error("CstAccountOTPStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

func (store *CstAccountOTPStore) getOTPByKey(key string) (res model.CstAccountOTP, err error) {
	var tmp model.CstAccountOTP
	err = store.DoHGETALL(key, &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("CstAccountOTPStore", logger.FromError(err))
	}
	return
}

func (store *CstAccountOTPStore) saveOTPByKey(key string, data model.CstAccountOTP, ttlSeconds int) error {
	if err := store.DoHMSET(key, &data, ttlSeconds); err != nil {
		logger.Fatal("CstAccountOTPStore", logger.FromError(err))
		return err
	}
	return nil
}

func (store *CstAccountOTPStore) saveNilByKey(key string, ttlSeconds int) error {
	if err := store.DoHMSET(key, emptyItem, ttlSeconds); err != nil {
		logger.Fatal("CstAccountOTPStore", logger.FromError(err))
		return err
	}
	return nil
}

func (store *CstAccountOTPStore) generateStoreKeyByID(id int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, id)
}

func (store *CstAccountOTPStore) generateStoreKeyByAccountAndAction(accountID int64, action string) string {
	return fmt.Sprintf("%s:%s:%v:%s", store.baseKey, store.byAccountAndAction, accountID, action)
}
