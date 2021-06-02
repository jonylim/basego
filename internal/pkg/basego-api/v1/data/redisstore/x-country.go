package redisstore

import (
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/gomodule/redigo/redis"
)

// XCountryStore manages Redis operations for countries.
type XCountryStore struct {
	redisStore
	ttl        int
	byID       string
	byISO2Code string
	byISO3Code string

	orgID int32
}

// NewXCountryStore returns new instance of XCountryStore.
func NewXCountryStore(conn redis.Conn) *XCountryStore {
	return &XCountryStore{
		redisStore: redisStore{
			conn:    conn,
			baseKey: "xCountry",
		},
		ttl:        30 * 86400, // 30 days
		byID:       "id",
		byISO2Code: "iso2",
		byISO3Code: "iso3",
	}
}

// GetByID returns a country's details by ID.
func (store *XCountryStore) GetByID(id int32) (model.XCountry, error) {
	return store.getByKey(store.generateStoreKeyByID(id))
}

// GetByISO2Code returns a country's details by 2-letters country code.
func (store *XCountryStore) GetByISO2Code(iso2Code string) (model.XCountry, error) {
	return store.getByKey(store.generateStoreKeyByISO2Code(iso2Code))
}

// GetByISO3Code returns a country's details by 3-letters country code.
func (store *XCountryStore) GetByISO3Code(iso3Code string) (model.XCountry, error) {
	return store.getByKey(store.generateStoreKeyByISO3Code(iso3Code))
}

func (store *XCountryStore) getByKey(key string) (res model.XCountry, err error) {
	var tmp model.XCountry
	err = store.DoHGETALL(key, &tmp)
	if err == nil {
		if !tmp.RedisNil {
			res = tmp
		}
	} else if err != redis.ErrNil {
		logger.Error("XCountryStore", logger.FromError(err))
	}
	return
}

// Save saves a country's details.
func (store *XCountryStore) Save(item model.XCountry) (err error) {
	// Delete old data.
	store.Delete(item)

	// Save new data.
	keys := store.generateStoreKeys(item)
	for _, key := range keys {
		if err = store.DoHMSET(key, &item, store.ttl); err != nil {
			logger.Fatal("XCountryStore", logger.FromError(err))
			return
		}
	}
	return
}

// SaveNilByID saves an empty country's details for a country ID.
func (store *XCountryStore) SaveNilByID(id int32) error {
	return store.saveNilByKey(store.generateStoreKeyByID(id))
}

// SaveNilByISO2Code saves an empty country's details for a 2-letters country code.
func (store *XCountryStore) SaveNilByISO2Code(iso2Code string) error {
	return store.saveNilByKey(store.generateStoreKeyByISO2Code(iso2Code))
}

// SaveNilByISO3Code saves an empty country's details for a 3-letters country code.
func (store *XCountryStore) SaveNilByISO3Code(iso3Code string) error {
	return store.saveNilByKey(store.generateStoreKeyByISO3Code(iso3Code))
}

// SaveNilByIDs saves empty country's details for a list of country IDs.
func (store *XCountryStore) SaveNilByIDs(ids []int32) (countSuccess int, errs []error) {
	errs = make([]error, len(ids))
	for i, id := range ids {
		if errs[i] = store.DoHMSET(store.generateStoreKeyByID(id), emptyItem, store.ttl); errs[i] != nil {
			logger.Fatal("XCountryStore", logger.FromError(errs[i]))
		} else {
			countSuccess++
		}
	}
	return
}

func (store *XCountryStore) saveNilByKey(key string) (err error) {
	if err = store.DoHMSET(key, emptyItem, store.ttl); err != nil {
		logger.Fatal("XCountryStore", logger.FromError(err))
	}
	return
}

// Delete deletes a country's details by all possible keys.
func (store *XCountryStore) Delete(item model.XCountry) (bool, error) {
	checkKeys := store.generateStoreKeys(item)
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
		logger.Fatal("XCountryStore", logger.FromError(err))
		return false, err
	}
	return count != 0, nil
}

// DeleteByID deletes a country's details by ID.
func (store *XCountryStore) DeleteByID(id int32) (deleted bool, err error) {
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
		logger.Fatal("XCountryStore", logger.FromError(err))
	}
	deleted = count != 0
	return
}

func (store *XCountryStore) generateStoreKeys(item model.XCountry) []string {
	return []string{
		store.generateStoreKeyByID(item.ID),
		store.generateStoreKeyByISO2Code(item.CountryCodeISO2),
		store.generateStoreKeyByISO3Code(item.CountryCodeISO3),
	}
}

func (store *XCountryStore) generateStoreKeyByID(id int32) string {
	return fmt.Sprintf("%s:%s:%v", store.baseKey, store.byID, id)
}

func (store *XCountryStore) generateStoreKeyByISO2Code(iso2Code string) string {
	return fmt.Sprintf("%s:%s:%s", store.baseKey, store.byISO2Code, iso2Code)
}

func (store *XCountryStore) generateStoreKeyByISO3Code(iso3Code string) string {
	return fmt.Sprintf("%s:%s:%s", store.baseKey, store.byISO3Code, iso3Code)
}
