package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/token/otp"

	"github.com/gomodule/redigo/redis"
)

// CstAccountOTPRepo manages data operations for customer account OTP data, especially cache operations.
type CstAccountOTPRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	store     *redisstore.CstAccountOTPStore
	cacheTTL  int
}

// NewCstAccountOTPRepo returns new instance of CstAccountOTPRepo.
func NewCstAccountOTPRepo(redisConn redis.Conn) *CstAccountOTPRepo {
	return &CstAccountOTPRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewCstAccountOTPStore(redisConn),
		cacheTTL:  otp.TTL + 300,
	}
}

// RedisStore returns Redis store used by the repository.
func (instance *CstAccountOTPRepo) RedisStore() *redisstore.CstAccountOTPStore {
	return instance.store
}

// GetOTPByID returns a customer account OTP's details by ID.
func (instance *CstAccountOTPRepo) GetOTPByID(id int64) (model.CstAccountOTP, error) {
	// Get from Redis.
	otp, err := instance.store.GetOTPByID(id)
	if err != nil {
		// Get from database.
		da := dao.NewCstAccountOTPDAO()
		da.WithDeleted()
		otp, err = da.GetOTPByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByID(id, instance.cacheTTL)
				return otp, instance.ErrNotFound
			}
			return otp, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.SaveOTPByID(otp, instance.cacheTTL)
	}
	if instance.exists(otp) {
		return otp, nil
	}
	return otp, instance.ErrNotFound
}

// GetActiveOTPByAccountAndAction returns an active OTP's details by ID.
func (instance *CstAccountOTPRepo) GetActiveOTPByAccountAndAction(accountID int64, action string) (model.CstAccountOTP, error) {
	// Get from Redis.
	otp, err := instance.store.GetActiveOTPByAccountAndAction(accountID, action)
	if err != nil {
		// Get from database.
		da := dao.NewCstAccountOTPDAO()
		da.WithDeleted()
		otp, err = da.GetActiveOTPByAccountAndAction(accountID, action)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByAccountAndAction(accountID, action, instance.cacheTTL)
				return otp, instance.ErrNotFound
			}
			return otp, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.SaveOTPByAccountAndAction(otp, instance.cacheTTL)
	}
	if instance.exists(otp) {
		return otp, nil
	}
	return otp, instance.ErrNotFound
}

func (instance *CstAccountOTPRepo) exists(otp model.CstAccountOTP) bool {
	return !otp.RedisNil && otp.ID != 0 && otp.AccountID != 0 && otp.Action != "" && otp.DeletedTime == 0
}
