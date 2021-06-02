package repository

import (
	"database/sql"
	"fmt"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

type xCountryRepoFieldMethods struct {
	getFromDAO     func(string) (model.XCountry, error)
	getFromRedis   func(string) (model.XCountry, error)
	saveNilToRedis func(string) error
}

// XCountryRepo manages data operations for country data, especially cache operations.
type XCountryRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn
	orgID     int64
	store     *redisstore.XCountryStore
	da        *dao.XCountryDAO

	methods map[string]xCountryRepoFieldMethods
}

// NewXCountryRepo returns new instance of XCountryRepo.
func NewXCountryRepo(redisConn redis.Conn) *XCountryRepo {
	repo := &XCountryRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn: redisConn,
		store:     redisstore.NewXCountryStore(redisConn),
	}

	repo.da = dao.NewXCountryDAO()
	repo.da.WithDeleted()

	repo.methods = map[string]xCountryRepoFieldMethods{
		"iso2Code": xCountryRepoFieldMethods{
			getFromDAO:     repo.da.GetByISO2Code,
			getFromRedis:   repo.store.GetByISO2Code,
			saveNilToRedis: repo.store.SaveNilByISO2Code,
		},
		"iso3Code": xCountryRepoFieldMethods{
			getFromDAO:     repo.da.GetByISO3Code,
			getFromRedis:   repo.store.GetByISO3Code,
			saveNilToRedis: repo.store.SaveNilByISO3Code,
		},
	}

	return repo
}

// RedisStore returns Redis store used by the repository.
func (instance *XCountryRepo) RedisStore() *redisstore.XCountryStore {
	return instance.store
}

// GetByID returns a country's details by country ID.
func (instance *XCountryRepo) GetByID(id int32) (model.XCountry, error) {
	// Get from Redis.
	res, err := instance.store.GetByID(id)
	if err != nil {
		// Get from database.
		res, err = instance.da.GetByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.store.SaveNilByID(id)
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

// GetByISO2Code returns a country's details by 2-letters country code.
func (instance *XCountryRepo) GetByISO2Code(iso2Code string) (model.XCountry, error) {
	return instance.getByField("iso2Code", iso2Code)
}

// GetByISO3Code returns a country's details by 3-letters country code.
func (instance *XCountryRepo) GetByISO3Code(iso3Code string) (model.XCountry, error) {
	return instance.getByField("iso3Code", iso3Code)
}

func (instance *XCountryRepo) getByField(field, param string) (model.XCountry, error) {
	if _, ok := instance.methods[field]; !ok {
		return model.XCountry{}, fmt.Errorf("Method not found, cant't get country by field '%s'", field)
	}
	// Get from Redis.
	country, err := instance.methods[field].getFromRedis(param)
	if err != nil {
		// Get from database.
		country, err = instance.methods[field].getFromDAO(param)
		if err != nil {
			if err == sql.ErrNoRows {
				// Save nil to Redis.
				instance.methods[field].saveNilToRedis(param)
				return country, instance.ErrNotFound
			}
			return country, instance.ErrDatabase
		}
		// Save to Redis.
		instance.store.Save(country)
	}
	if instance.exists(country) {
		return country, nil
	}
	return country, instance.ErrNotFound
}

func (instance *XCountryRepo) exists(country model.XCountry) bool {
	return !country.RedisNil && country.ID != 0
}
