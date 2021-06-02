package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// CstAccountSessionTokenStore manages Redis operations for customer account session tokens.
type CstAccountSessionTokenStore struct {
	redisStore
	bySessionID string
	ttl         int
}

// NewCstAccountSessionTokenStore returns new instance to manage customer account session tokens.
func NewCstAccountSessionTokenStore(conn redis.Conn) *CstAccountSessionTokenStore {
	return &CstAccountSessionTokenStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "cstAccSessToken",
		},
		bySessionID: "sid",
		ttl:         7 * 86400,
	}
}

// GetTokenBySessionID returns a customer account session's access & refresh tokens by session ID.
func (store *CstAccountSessionTokenStore) GetTokenBySessionID(sessionID int64) (model.CstAccountSessionToken, error) {
	var tmp, res model.CstAccountSessionToken
	err := store.DoHGETALL(store.generateStoreKeyBySessionID(sessionID), &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("CstAccountSessionTokenStore", logger.FromError(err))
	}
	return res, err
}

// SaveToken saves a customer account session's access token and refresh token.
func (store *CstAccountSessionTokenStore) SaveToken(item model.CstAccountSessionToken) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyBySessionID(item.SessionID), &item, store.ttl); err != nil {
		logger.Fatal("CstAccountSessionTokenStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// SaveNilBySessionID saves an empty customer account session's access token and refresh token by session ID.
func (store *CstAccountSessionTokenStore) SaveNilBySessionID(sessionID int64) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyBySessionID(sessionID), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountSessionTokenStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// SaveNilBySessionIDs saves empty customer account session's access tokens and refresh tokens by session IDs.
func (store *CstAccountSessionTokenStore) SaveNilBySessionIDs(sessionIDs []int64) (countSuccess int, errs []error) {
	errs = make([]error, len(sessionIDs))
	for i, sid := range sessionIDs {
		if errs[i] = store.DoHMSET(store.generateStoreKeyBySessionID(sid), emptyItem, store.ttl); errs[i] != nil {
			logger.Fatal("CstAccountSessionTokenStore", logger.FromError(errs[i]))
		} else {
			countSuccess++
		}
	}
	return
}

// DeleteTokenBySessionID deletes a customer account session's token details by session ID.
func (store *CstAccountSessionTokenStore) DeleteTokenBySessionID(sessionID int64) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKeyBySessionID(sessionID))
	if err != nil {
		logger.Error("CstAccountSessionTokenStore", logger.FromError(err))
		return false, err
	}
	return (count != 0), nil
}

// DeleteTokensBySessionIDs deletes customer account session tokens' token details by session IDs.
func (store *CstAccountSessionTokenStore) DeleteTokensBySessionIDs(sessionIDs []int64) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKeysBySessionIDs(sessionIDs)...)
	if err != nil {
		logger.Error("CstAccountSessionTokenStore", logger.FromError(err))
		return false, err
	}
	return (count != 0), err
}

func (store *CstAccountSessionTokenStore) generateStoreKeyBySessionID(sessionID int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.bySessionID, sessionID)
}

func (store *CstAccountSessionTokenStore) generateStoreKeysBySessionIDs(sessionIDs []int64) []string {
	keys := make([]string, len(sessionIDs))
	for i, sid := range sessionIDs {
		keys[i] = store.generateStoreKeyBySessionID(sid)
	}
	return keys
}
