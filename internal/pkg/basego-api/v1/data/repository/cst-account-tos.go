package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// CstAccountTOSRepo manages data operations for customer account's Terms of Service status, especially cache operations.
type CstAccountTOSRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	store     *redisstore.CstAccountTOSStore
}

// NewCstAccountTOSRepo returns new instance of CstAccountTOSRepo.
func NewCstAccountTOSRepo(redisConn redis.Conn) *CstAccountTOSRepo {
	return &CstAccountTOSRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewCstAccountTOSStore(redisConn),
	}
}

// RedisConn returns Redis connection used by the repository.
func (instance *CstAccountTOSRepo) RedisConn() redis.Conn {
	return instance.redisConn
}

// RedisStore returns Redis store used by the repository.
func (instance *CstAccountTOSRepo) RedisStore() *redisstore.CstAccountTOSStore {
	return instance.store
}

// GetByID returns a customer account's TOS status by ID.
func (instance *CstAccountTOSRepo) GetByID(id int64) (model.CstAccountTOS, error) {
	// Get from Redis.
	tos, err := instance.store.GetByID(id)
	if err != nil {
		// Get from database.
		tos, err = dao.NewCstAccountTOSDAO().GetByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByID(id)
				return tos, instance.ErrNotFound
			}
			return tos, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(tos)
	}
	if instance.exists(tos) {
		return tos, nil
	}
	return tos, instance.ErrNotFound
}

// GetByAccountID returns a customer account's TOS status by ID.
func (instance *CstAccountTOSRepo) GetByAccountID(accountID int64) (model.CstAccountTOS, error) {
	// Get from Redis.
	tos, err := instance.store.GetByAccountID(accountID)
	if err != nil {
		// Get from database.
		tos, err = dao.NewCstAccountTOSDAO().GetByAccountID(accountID)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByAccountID(accountID)
				return tos, instance.ErrNotFound
			}
			return tos, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(tos)
	}
	if instance.exists(tos) {
		return tos, nil
	}
	return tos, instance.ErrNotFound
}

// ExistsByAccountID checks if a customer account's TOS status exists by account ID.
func (instance *CstAccountTOSRepo) ExistsByAccountID(accountID int64) (bool, error) {
	// Get from Redis.
	tos, err := instance.store.GetByAccountID(accountID)
	if err == nil {
		return instance.exists(tos), nil
	}
	// Get from database.
	tos, err = dao.NewCstAccountTOSDAO().GetByAccountID(accountID)
	if err == nil {
		// Save to Redis.
		instance.store.Save(tos)
		return instance.exists(tos), nil
	} else if err == sql.ErrNoRows {
		// Save nil to Redis.
		instance.store.SaveNilByAccountID(accountID)
		return false, nil
	}
	return false, instance.ErrDatabase
}

func (instance *CstAccountTOSRepo) exists(tos model.CstAccountTOS) bool {
	return !tos.RedisNil && tos.ID != 0 && tos.AccountID != 0 && tos.DeletedTime == 0
}
