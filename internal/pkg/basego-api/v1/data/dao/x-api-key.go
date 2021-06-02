package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// XAPIKeyDAO manages database operations for API key data.
type XAPIKeyDAO struct {
	dao
	domain string
}

// NewXAPIKeyDAO returns new instance of XAPIKeyDAO.
func NewXAPIKeyDAO(domain string) *XAPIKeyDAO {
	return &XAPIKeyDAO{
		dao:    dao{db.Get(), false},
		domain: domain,
	}
}

func (instance *XAPIKeyDAO) getByFieldAndValue(field string, value interface{}) (res model.XAPIKey, err error) {
	var sqlWhereDeleted string
	if !instance.withDeleted {
		sqlWhereDeleted = `AND deleted_at IS NULL`
	}
	err = instance.db.QueryRow(`SELECT
				id, api_key_id, api_key_secret,
				domain, app_platform, app_identifier,
				`+sqlTimestampToUnixMilliseconds("expiry_time")+` AS expiry_time, is_enabled,
				`+sqlTimestampToUnixMilliseconds("created_at")+` AS created_time,
				`+sqlTimestampToUnixMilliseconds("updated_at")+` AS updated_time,
				`+sqlTimestampToUnixMilliseconds("deleted_at")+` AS deleted_time
			FROM tb_x_api_key 
			WHERE domain = $1
				AND `+field+` = $2
				`+sqlWhereDeleted,
		instance.domain, value).
		Scan(&res.ID, &res.APIKeyID, &res.APIKeySecret,
			&res.Domain, &res.AppPlatform, &res.AppIdentifier,
			&res.ExpiryTime, &res.IsEnabled,
			&res.CreatedTime, &res.UpdatedTime, &res.DeletedTime)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("XAPIKeyDAO", logger.FromError(err))
	}
	return
}

// GetByID returns an API key's details by record ID.
func (instance *XAPIKeyDAO) GetByID(id int64) (model.XAPIKey, error) {
	return instance.getByFieldAndValue("id", id)
}

// GetByAPIKeyID returns an API key's details by API key ID.
func (instance *XAPIKeyDAO) GetByAPIKeyID(appKeyID string) (model.XAPIKey, error) {
	return instance.getByFieldAndValue("api_key_id", appKeyID)
}
