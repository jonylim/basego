package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// CstAccountSessionStore manages Redis operations for customer account sessions.
type CstAccountSessionStore struct {
	redisStore
	byID string
	ttl  int
}

// NewCstAccountSessionStore returns new instance to manage customer account sessions.
func NewCstAccountSessionStore(conn redis.Conn) *CstAccountSessionStore {
	return &CstAccountSessionStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "cstAccSess",
		},
		byID: "id",
		ttl:  7 * 86400, // 7 days
	}
}

// GetSessionByID returns a customer account session's details by session ID.
func (store *CstAccountSessionStore) GetSessionByID(sessionID int64) (model.CstAccountSession, error) {
	var tmp, res model.CstAccountSession
	err := store.DoHGETALL(store.generateStoreKeyByID(sessionID), &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("CstAccountSessionStore", logger.FromError(err))
	}
	return res, err
}

// SaveSession saves a customer account session's details.
func (store *CstAccountSessionStore) SaveSession(item model.CstAccountSession) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyByID(item.ID), &item, store.ttl); err != nil {
		logger.Fatal("CstAccountSessionStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// SaveNilByID saves an empty customer account session's details by session ID.
func (store *CstAccountSessionStore) SaveNilByID(sessionID int64) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyByID(sessionID), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountSessionStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// SaveNilByIDs saves empty customer account session's details by session IDs.
func (store *CstAccountSessionStore) SaveNilByIDs(sessionIDs []int64) (countSuccess int, errs []error) {
	errs = make([]error, len(sessionIDs))
	for i, sid := range sessionIDs {
		if errs[i] = store.DoHMSET(store.generateStoreKeyByID(sid), emptyItem, store.ttl); errs[i] != nil {
			logger.Fatal("CstAccountSessionStore", logger.FromError(errs[i]))
		} else {
			countSuccess++
		}
	}
	return
}

// ResetTTLByID resets a store key's TTL by session ID.
func (store *CstAccountSessionStore) ResetTTLByID(sessionID int64) {
	store.conn.Do("EXPIRE", store.generateStoreKeyByID(sessionID), store.ttl)
}

// DeleteByID deletes a customer account session's details by session ID.
func (store *CstAccountSessionStore) DeleteByID(sessionID int64) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKeyByID(sessionID))
	if err != nil {
		logger.Error("CstAccountSessionStore", logger.FromError(err))
		return false, err
	}
	return (count != 0), nil
}

// DeleteByIDs deletes multiple customer account sessions' details by session IDs.
func (store *CstAccountSessionStore) DeleteByIDs(sessionIDs []int64) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKeysByIDs(sessionIDs)...)
	if err != nil {
		logger.Error("CstAccountSessionStore", logger.FromError(err))
		return false, err
	}
	return (count != 0), err
}

func (store *CstAccountSessionStore) generateStoreKeyByID(sessionID int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, sessionID)
}

func (store *CstAccountSessionStore) generateStoreKeysByIDs(sessionIDs []int64) []string {
	keys := make([]string, len(sessionIDs))
	for i, sid := range sessionIDs {
		keys[i] = store.generateStoreKeyByID(sid)
	}
	return keys
}
