package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// APIKeyRepo manages data operations for API keys, especially cache operations.
type APIKeyRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	domain    string
}

// NewAPIKeyRepo returns new instance of APIKeyRepo.
func NewAPIKeyRepo(redisConn redis.Conn, domain string) *APIKeyRepo {
	return &APIKeyRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		domain:    domain,
	}
}

// GetByAPIKeyID returns an API key's details by API key ID.
func (instance *APIKeyRepo) GetByAPIKeyID(apiKeyID string) (model.XAPIKey, error) {
	// Get from Redis.
	store := redisstore.NewXAPIKeyStore(instance.redisConn, instance.domain)
	apiKey, err := store.GetByAPIKeyID(apiKeyID)
	if err != nil {
		// Get from database.
		da := dao.NewXAPIKeyDAO(instance.domain)
		da.WithDeleted()
		apiKey, err = da.GetByAPIKeyID(apiKeyID)
		if err != nil {
			if err == sql.ErrNoRows {
				return apiKey, instance.ErrNotFound
			}
			return apiKey, instance.ErrDatabase
		}
		// Save to Redis.
		store.Save(apiKey)
	}
	// Check deleted.
	if apiKey.DeletedTime != 0 {
		return apiKey, instance.ErrNotFound
	}
	return apiKey, err
}
