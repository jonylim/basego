package redisstore

import (
	"fmt"
	"strings"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// PgTimeZoneStore manages Redis operations for time zones.
type PgTimeZoneStore struct {
	redisStore
	byName string
	ttl    int
}

// NewPgTimeZoneStore returns new instance of PgTimeZoneStore.
func NewPgTimeZoneStore(conn redis.Conn) *PgTimeZoneStore {
	return &PgTimeZoneStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "pgTZ",
		},
		byName: "name",
		ttl:    86400, // 1 day
	}
}

// GetByName returns a time zone's details by name.
func (store *PgTimeZoneStore) GetByName(name string) (model.PgTimeZone, error) {
	var tmp, res model.PgTimeZone
	err := store.DoHGETALL(store.generateStoreKeyByName(name), &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("PgTimeZoneStore", logger.FromError(err))
	}
	return res, err
}

// Save saves a time zone's details.
func (store *PgTimeZoneStore) Save(item model.PgTimeZone) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyByName(item.Name), item, store.ttl); err != nil {
		logger.Fatal("PgTimeZoneStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// SaveNilByName saves an empty time zone's details by name.
func (store *PgTimeZoneStore) SaveNilByName(name string) (bool, error) {
	if err := store.DoHMSET(store.generateStoreKeyByName(name), emptyItem, store.ttl); err != nil {
		logger.Fatal("PgTimeZoneStore", logger.FromError(err))
		return false, err
	}
	return true, nil
}

// DeleteByName deletes a time zone's details by name.
func (store *PgTimeZoneStore) DeleteByName(name string) (bool, error) {
	count, err := store.DoDEL(store.generateStoreKeyByName(name))
	if err != nil {
		logger.Fatal("PgTimeZoneStore", logger.FromError(err))
		return false, err
	}
	return (count != 0), nil
}

func (store *PgTimeZoneStore) generateStoreKeyByName(name string) string {
	return fmt.Sprintf("%s:%s:%s", store.baseKey, store.byName, strings.ToLower(name))
}
