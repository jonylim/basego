package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// XAPIKeyStore manages Redis operations for API key data.
type XAPIKeyStore struct {
	redisStore
	domain string
	ttl    int
}

// NewXAPIKeyStore returns new instance to manage API keys for a domain ("customer" or "internal").
func NewXAPIKeyStore(conn redis.Conn, domain string) *XAPIKeyStore {
	if domain == "" {
		logger.Warn("XAPIKeyStore", `Domain is empty`)
	}
	return &XAPIKeyStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "apiKey",
		},
		domain: domain,
		ttl:    30 * 86400, // 30 days
	}
}

// GetByAPIKeyID returns an API key's details by API key ID.
func (store *XAPIKeyStore) GetByAPIKeyID(apiKeyID string) (model.XAPIKey, error) {
	var data model.XAPIKey
	err := store.DoHGETALL(store.generateStoreKey(apiKeyID), &data)
	if err != nil && err != redis.ErrNil {
		logger.Error("XAPIKeyStore", logger.FromError(err))
	}
	return data, err
}

// Save saves an API key's details.
func (store *XAPIKeyStore) Save(item model.XAPIKey) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKey(item.APIKeyID), &item, store.ttl); err != nil {
		logger.Fatal("XAPIKeyStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// Delete deletes an API key's details by API key ID.
func (store *XAPIKeyStore) Delete(apiKeyID string) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKey(apiKeyID))
	if err != nil {
		if err != redis.ErrNil {
			logger.Error("XAPIKeyStore", logger.FromError(err))
		}
		return false, err
	}
	return (count != 0), nil
}

func (store *XAPIKeyStore) generateStoreKey(apiKeyID string) string {
	return fmt.Sprintf("%s:%s:%s", store.baseKey, store.domain, apiKeyID)
}
