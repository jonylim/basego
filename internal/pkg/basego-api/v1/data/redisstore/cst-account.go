package redisstore

import (
	"fmt"
	"strings"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// CstAccountStore manages Redis operations for customer account data.
type CstAccountStore struct {
	redisStore
	ttl     int
	byID    string
	byEmail string
}

type cstAccountStoreModel struct {
	RedisNil              bool   `redis:"redisNil"`
	ID                    int64  `redis:"id"`
	FullName              string `redis:"fullName"`
	Email                 string `redis:"email"`
	IsEmailVerified       bool   `redis:"isEmailVerified"`
	CountryID             int32  `redis:"countryID"`
	CountryCallingCode    string `redis:"countryCC"`
	Phone                 string `redis:"phone"`
	PhoneWithCode         string `redis:"phoneWithCode"`
	IsPhoneVerified       bool   `redis:"isPhoneVerified"`
	Password              string `redis:"password"`
	PasswordSalt          string `redis:"passwordSalt"`
	Use2FA                bool   `redis:"use2FA"`
	ImgThumbnailURL       string `redis:"imgThumbURL"`
	ImgFullsizeURL        string `redis:"imgFullURL"`
	LastLoginTime         int64  `redis:"lastLoginTime"`
	LastActivityTime      int64  `redis:"lastActivityTime"`
	RequireChangePassword bool   `redis:"requireChangePassword"`
	CreatedTime           int64  `redis:"createdTime"`
	UpdatedTime           int64  `redis:"updatedTime"`
	DeletedTime           int64  `redis:"deletedTime"`
}

func (src cstAccountStoreModel) Inflate() (res model.CstAccount) {
	if !src.RedisNil {
		imgURL := model.ImageURL{
			Thumbnail: src.ImgThumbnailURL,
			Fullsize:  src.ImgFullsizeURL,
		}
		res = model.CstAccount{
			ID:                    src.ID,
			FullName:              src.FullName,
			Email:                 src.Email,
			IsEmailVerified:       src.IsEmailVerified,
			CountryID:             src.CountryID,
			CountryCallingCode:    src.CountryCallingCode,
			Phone:                 src.Phone,
			PhoneWithCode:         src.PhoneWithCode,
			IsPhoneVerified:       src.IsPhoneVerified,
			Password:              src.Password,
			PasswordSalt:          src.PasswordSalt,
			Use2FA:                src.Use2FA,
			ImageURL:              imgURL,
			LastLoginTime:         src.LastLoginTime,
			LastActivityTime:      src.LastActivityTime,
			RequireChangePassword: src.RequireChangePassword,
			CreatedTime:           src.CreatedTime,
			UpdatedTime:           src.UpdatedTime,
			DeletedTime:           src.DeletedTime,
		}
	}
	return
}

func toCstAccountStoreModel(src model.CstAccount) cstAccountStoreModel {
	return cstAccountStoreModel{
		ID:                    src.ID,
		FullName:              src.FullName,
		Email:                 src.Email,
		IsEmailVerified:       src.IsEmailVerified,
		CountryID:             src.CountryID,
		CountryCallingCode:    src.CountryCallingCode,
		Phone:                 src.Phone,
		PhoneWithCode:         src.PhoneWithCode,
		IsPhoneVerified:       src.IsPhoneVerified,
		Password:              src.Password,
		PasswordSalt:          src.PasswordSalt,
		Use2FA:                src.Use2FA,
		ImgThumbnailURL:       src.ImageURL.Thumbnail,
		ImgFullsizeURL:        src.ImageURL.Fullsize,
		LastLoginTime:         src.LastLoginTime,
		LastActivityTime:      src.LastActivityTime,
		RequireChangePassword: src.RequireChangePassword,
		CreatedTime:           src.CreatedTime,
		UpdatedTime:           src.UpdatedTime,
		DeletedTime:           src.DeletedTime,
	}
}

// NewCstAccountStore returns new instance to manage customer account data.
func NewCstAccountStore(conn redis.Conn) *CstAccountStore {
	return &CstAccountStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "cstAcc",
		},
		ttl:     3600, // 1 hour
		byID:    "id",
		byEmail: "email",
	}
}

// GetByID returns a customer account's details by ID.
func (store *CstAccountStore) GetByID(id int64) (model.CstAccount, error) {
	return store.getByKey(store.generateStoreKeyByID(id))
}

// GetByEmail returns a customer account's details by email address.
func (store *CstAccountStore) GetByEmail(email string) (model.CstAccount, error) {
	return store.getByKey(store.generateStoreKeyByEmail(email))
}

// GetByIDs returns customer accounts' details by IDs.
func (store *CstAccountStore) GetByIDs(ids []int64) (items []model.CstAccount, errs map[int64]error) {
	items = make([]model.CstAccount, 0, len(ids))
	errs = make(map[int64]error)
	for _, id := range ids {
		res, err := store.getByKey(store.generateStoreKeyByID(id))
		if err != nil {
			errs[id] = err
		} else if res.ID != 0 {
			items = append(items, res)
		}
	}
	return
}

// ExistsByID checks if a customer account exists by ID.
/* func (store *CstAccountStore) ExistsByID(id int64) (bool, error) {
	return store.existsByKey(store.generateStoreKeyByID(id))
} */

// ExistsByEmail checks if a customer account exists by email address.
/* func (store *CstAccountStore) ExistsByEmail(email string) (bool, error) {
	return store.existsByKey(store.generateStoreKeyByEmail(email))
} */

// Save saves a customer account's details.
func (store *CstAccountStore) Save(account model.CstAccount) error {
	// Delete old data.
	store.Delete(account)

	// Save new data.
	keys := store.generateStoreKeys(account)
	data := toCstAccountStoreModel(account)
	for _, key := range keys {
		if err := store.DoHMSET(key, &data, store.ttl); err != nil {
			logger.Fatal("CstAccountStore", logger.FromError(err))
			return err
		}
	}
	return nil
}

// SaveNilByID saves an empty customer account's details for the ID.
func (store *CstAccountStore) SaveNilByID(id int64) error {
	if err := store.DoHMSET(store.generateStoreKeyByID(id), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountStore", logger.FromError(err))
		return err
	}
	return nil
}

// SaveNilByEmail saves an empty customer account's details for the email address.
func (store *CstAccountStore) SaveNilByEmail(email string) error {
	if err := store.DoHMSET(store.generateStoreKeyByEmail(email), emptyItem, store.ttl); err != nil {
		logger.Fatal("CstAccountStore", logger.FromError(err))
		return err
	}
	return nil
}

// SaveNilByIDs saves empty customer account's details for the IDs.
func (store *CstAccountStore) SaveNilByIDs(ids []int64) (countSuccess int, errs []error) {
	errs = make([]error, len(ids))
	for i, sid := range ids {
		if errs[i] = store.DoHMSET(store.generateStoreKeyByID(sid), emptyItem, store.ttl); errs[i] != nil {
			logger.Fatal("CstAccountStore", logger.FromError(errs[i]))
		} else {
			countSuccess++
		}
	}
	return
}

// UpdateLastActivityTime updates a customer account's last activity time.
func (store *CstAccountStore) UpdateLastActivityTime(account model.CstAccount) (bool, error) {
	checkKeys := store.generateStoreKeys(account)
	updated := false
	for _, key := range checkKeys {
		if existing, err := store.getByKey(key); err == nil {
			if existing.ID == account.ID {
				if _, err := store.conn.Do("HMSET", key, "lastActivityTime", account.LastActivityTime); err != nil {
					logger.Fatal("CstAccountStore", logger.FromError(err))
					return false, err
				}
			} else {
				store.Delete(existing)
			}
			updated = true
		}
	}
	return updated, nil
}

// Delete deletes a customer account's details by all possible keys.
func (store *CstAccountStore) Delete(account model.CstAccount) (bool, error) {
	checkKeys := store.generateStoreKeys(account)
	deleteKeys := make([]string, 0, len(checkKeys))
	deleteKeys = append(deleteKeys, checkKeys...)
	for _, key := range checkKeys {
		if existing, err := store.getByKey(key); err == nil {
			existingKeys := store.generateStoreKeys(existing)
			deleteKeys = append(deleteKeys, existingKeys...)
		}
	}
	if len(deleteKeys) == 0 {
		return false, nil
	}
	count, err := store.DoDEL(deleteKeys...)
	if err != nil && err != redis.ErrNil {
		logger.Fatal("CstAccountStore", logger.FromError(err))
		return false, err
	}
	return count != 0, nil
}

// DeleteByID deletes a customer account's details by ID.
func (store *CstAccountStore) DeleteByID(id int64) (deleted bool, err error) {
	var count int
	key := store.generateStoreKeyByID(id)
	if existing, err1 := store.getByKey(key); err1 == nil {
		deleteKeys := store.generateStoreKeys(existing)
		deleteKeys = append(deleteKeys, key)
		count, err = store.DoDEL(deleteKeys...)
	} else {
		count, err = store.DoDEL(key)
	}
	if err != nil && err != redis.ErrNil {
		logger.Fatal("CstAccountStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

func (store *CstAccountStore) getByKey(key string) (res model.CstAccount, err error) {
	var tmp cstAccountStoreModel
	err = store.DoHGETALL(key, &tmp)
	if err == nil {
		res = tmp.Inflate()
	} else if err != redis.ErrNil {
		logger.Error("CstAccountStore", logger.FromError(err))
	}
	return
}

/* func (store *CstAccountStore) existsByKey(key string) (exists bool, err error) {
	exists, err = store.DoEXISTS(key)
	if err != nil {
		logger.Error("CstAccountStore", logger.FromError(err))
	}
	return
} */

func (store *CstAccountStore) generateStoreKeys(account model.CstAccount) []string {
	return []string{
		store.generateStoreKeyByID(account.ID),
		store.generateStoreKeyByEmail(account.Email),
	}
}

func (store *CstAccountStore) generateStoreKeyByID(id int64) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, id)
}

func (store *CstAccountStore) generateStoreKeyByEmail(email string) string {
	// IMPORTANT! Always store the email key in lowercase.
	return fmt.Sprintf("%s:%s:%s", store.baseKey, store.byEmail, strings.ToLower(email))
}
