package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// CstAccountRepo manages data operations for customer account data, especially cache operations.
type CstAccountRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	store     *redisstore.CstAccountStore
}

// NewCstAccountRepo returns new instance of CstAccountRepo.
func NewCstAccountRepo(redisConn redis.Conn) *CstAccountRepo {
	return &CstAccountRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewCstAccountStore(redisConn),
	}
}

// RedisConn returns Redis connection used by the repository.
func (instance *CstAccountRepo) RedisConn() redis.Conn {
	return instance.redisConn
}

// RedisStore returns Redis store used by the repository.
func (instance *CstAccountRepo) RedisStore() *redisstore.CstAccountStore {
	return instance.store
}

// GetByID returns a customer account's details by ID.
func (instance *CstAccountRepo) GetByID(id int64) (model.CstAccount, error) {
	// Get from Redis.
	account, err := instance.store.GetByID(id)
	if err != nil {
		// Get from database.
		da := dao.NewCstAccountDAO()
		da.WithDeleted()
		account, err = da.GetByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByID(id)
				return account, instance.ErrNotFound
			}
			return account, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(account)
	}
	if instance.exists(account) {
		return account, nil
	}
	return account, instance.ErrNotFound
}

// GetByEmail returns a customer account's details by email address.
func (instance *CstAccountRepo) GetByEmail(email string) (model.CstAccount, error) {
	// Get from Redis.
	account, err := instance.store.GetByEmail(email)
	if err != nil {
		// Get from database.
		da := dao.NewCstAccountDAO()
		da.WithDeleted()
		account, err = da.GetByEmail(email)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByEmail(email)
				return account, instance.ErrNotFound
			}
			return account, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(account)
	}
	if instance.exists(account) {
		return account, nil
	}
	return account, instance.ErrNotFound
}

// GetByIDs returns a list of customer accounts' details by IDs.
func (instance *CstAccountRepo) GetByIDs(ids []int64) (res []model.CstAccount, err error) {
	// Get from Redis.
	accounts, errs := instance.store.GetByIDs(ids)
	if len(errs) != 0 {
		// Get from database.
		da := dao.NewCstAccountDAO()
		da.WithDeleted()
		accounts, err = da.GetByIDs(ids, true)
		if err != nil {
			err = instance.ErrDatabase
			return
		}
		switch len(accounts) {
		case 0:
			// Save all nil to Redis.
			instance.store.SaveNilByIDs(ids)
		case len(ids):
			// Save all to Redis.
			for _, acc := range accounts {
				instance.store.Save(acc)
			}
		default:
			nilIDs := make([]int64, 0)
			for _, id := range ids {
				found := false
				for _, acc := range accounts {
					if acc.ID == id {
						found = true
						break
					}
				}
				if !found {
					nilIDs = append(nilIDs, id)
				}
			}
			// Save to Redis.
			instance.store.SaveNilByIDs(nilIDs)
			for _, acc := range accounts {
				instance.store.Save(acc)
			}
		}
	}
	res = make([]model.CstAccount, 0, len(accounts))
	for _, acc := range accounts {
		if instance.exists(acc) {
			res = append(res, acc)
		}
	}
	return
}

// ExistsByID checks if a customer account exists by ID.
func (instance *CstAccountRepo) ExistsByID(id int64) (bool, error) {
	// Get from Redis.
	account, err := instance.store.GetByID(id)
	if err == nil {
		return instance.exists(account), nil
	}
	// Get from database.
	da := dao.NewCstAccountDAO()
	da.WithDeleted()
	account, err = da.GetByID(id)
	if err == nil {
		// Save to Redis.
		instance.store.Save(account)
		return instance.exists(account), nil
	} else if err == sql.ErrNoRows {
		// Save nil to Redis.
		instance.store.SaveNilByID(id)
		return false, nil
	}
	return false, instance.ErrDatabase
}

// ExistsByEmail checks if a customer account exists by email address.
func (instance *CstAccountRepo) ExistsByEmail(email string) (bool, error) {
	// Get from Redis.
	account, err := instance.store.GetByEmail(email)
	if err == nil {
		return instance.exists(account), nil
	}
	// Get from database.
	da := dao.NewCstAccountDAO()
	da.WithDeleted()
	account, err = da.GetByEmail(email)
	if err == nil {
		// Save to Redis.
		instance.store.Save(account)
		return instance.exists(account), nil
	} else if err == sql.ErrNoRows {
		// Save nil to Redis.
		instance.store.SaveNilByEmail(email)
		return false, nil
	}
	return false, instance.ErrDatabase
}

// SyncByID syncs a customer account's details to Redis by ID.
func (instance *CstAccountRepo) SyncByID(id int64) (model.CstAccount, error) {
	// Get from database.
	da := dao.NewCstAccountDAO()
	da.WithDeleted()
	account, err := da.GetByID(id)
	if err == nil {
		// Save to Redis.
		instance.store.Save(account)
		if instance.exists(account) {
			return account, nil
		}
		return account, instance.ErrNotFound
	} else if err == sql.ErrNoRows {
		// Save nil to Redis.
		instance.store.SaveNilByID(id)
		return account, instance.ErrNotFound
	}
	// Delete from Redis.
	instance.store.DeleteByID(id)
	return account, instance.ErrDatabase
}

func (instance *CstAccountRepo) exists(account model.CstAccount) bool {
	return account.ID != 0 && account.Email != "" && account.DeletedTime == 0
}
