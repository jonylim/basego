package repository

import (
	"context"
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// PgTimeZoneRepo manages data operations for time zones, especially cache operations.
type PgTimeZoneRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	store     *redisstore.PgTimeZoneStore
}

// NewPgTimeZoneRepo returns new instance of PgTimeZoneRepo.
func NewPgTimeZoneRepo(redisConn redis.Conn) *PgTimeZoneRepo {
	return &PgTimeZoneRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewPgTimeZoneStore(redisConn),
	}
}

// RedisStore returns Redis store used by the repository.
func (instance *PgTimeZoneRepo) RedisStore() *redisstore.PgTimeZoneStore {
	return instance.store
}

// GetByName returns a time zone's details by ID.
func (instance *PgTimeZoneRepo) GetByName(ctx context.Context, name string) (model.PgTimeZone, error) {
	// Get from Redis.
	res, err := instance.store.GetByName(name)
	if err != nil {
		// Get from database.
		res, err = dao.NewPgTimeZoneDAO().GetByName(ctx, name)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByName(name)
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

// ExistsByName checks if a time zone exists by ID.
func (instance *PgTimeZoneRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	// Get from Redis.
	res, err := instance.store.GetByName(name)
	if err == nil {
		return instance.exists(res), nil
	}
	// Get from database.
	res, err = dao.NewPgTimeZoneDAO().GetByName(ctx, name)
	if err == nil {
		// Save to Redis.
		instance.store.Save(res)
		return instance.exists(res), nil
	} else if err == sql.ErrNoRows {
		// Save nil to Redis.
		instance.store.SaveNilByName(name)
		return false, nil
	}
	return false, instance.ErrDatabase
}

func (instance *PgTimeZoneRepo) exists(item model.PgTimeZone) bool {
	return !item.RedisNil && item.Name != ""
}
