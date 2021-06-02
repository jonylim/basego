package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// XCountryDAO manages database operations for countries.
type XCountryDAO struct {
	dao
	selectColumns string
}

// NewXCountryDAO returns new instance of XCountryDAO.
func NewXCountryDAO() *XCountryDAO {
	return &XCountryDAO{
		dao: dao{db.Get(), false},
		selectColumns: `
			id, common_name, official_name, 
			iso2_code, iso3_code, calling_code, currency_code, 
			is_enabled, is_hidden `,
	}
}

func (instance *XCountryDAO) scanRows(rows *sql.Rows) ([]model.XCountry, error) {
	countries := make([]model.XCountry, 0)
	for rows.Next() {
		var data model.XCountry
		err := rows.Scan(&data.ID, &data.CommonName, &data.OfficialName,
			&data.CountryCodeISO2, &data.CountryCodeISO3, &data.CallingCode, &data.CurrencyCode,
			&data.IsEnabled, &data.IsHidden)
		if err != nil {
			return countries, err
		}
		countries = append(countries, data)
	}
	return countries, nil
}

func (instance *XCountryDAO) scanRow(row *sql.Row) (data model.XCountry, err error) {
	err = row.Scan(
		&data.ID, &data.CommonName, &data.OfficialName,
		&data.CountryCodeISO2, &data.CountryCodeISO3, &data.CallingCode, &data.CurrencyCode,
		&data.IsEnabled, &data.IsHidden)
	return
}

func (instance *XCountryDAO) getListWhere(sqlWhere, sqlOrderBy string, params ...interface{}) ([]model.XCountry, error) {
	rows, err := instance.db.Query(`
			SELECT `+instance.selectColumns+`
			FROM tb_x_country
			`+sqlWhere+`
			`+sqlOrderBy,
		params...)
	if err != nil {
		logger.Fatal("XCountryDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	countries, err := instance.scanRows(rows)
	if err != nil {
		logger.Fatal("XCountryDAO", logger.FromError(err))
	}
	return countries, err
}

func (instance *XCountryDAO) getOneWhere(sqlWhere string, params ...interface{}) (res model.XCountry, err error) {
	row := instance.db.QueryRow(`
			SELECT `+instance.selectColumns+`
			FROM tb_x_country
			`+sqlWhere,
		params...)
	res, err = instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("XCountryDAO", logger.FromError(err))
	}
	return
}

// GetAllCountryList returns the list of all countries.
func (instance *XCountryDAO) GetAllCountryList() ([]model.XCountry, error) {
	return instance.getListWhere(
		`WHERE deleted_at IS NULL`,
		`ORDER BY common_name, calling_code, id`)
}

// GetActiveCountryList returns the list of enabled non-hidden countries.
func (instance *XCountryDAO) GetActiveCountryList() ([]model.XCountry, error) {
	return instance.getListWhere(
		`WHERE is_enabled = TRUE
				AND is_hidden = FALSE
				AND deleted_at IS NULL`,
		`ORDER BY common_name, calling_code, id`)
}

// GetByID returns a country's details by ID.
func (instance *XCountryDAO) GetByID(countryID int32) (model.XCountry, error) {
	return instance.getOneWhere(
		`WHERE id = $1 AND deleted_at IS NULL`,
		countryID)
}

// GetByISO2Code returns a country's details by 2-letters country code.
func (instance *XCountryDAO) GetByISO2Code(iso2Code string) (model.XCountry, error) {
	return instance.getOneWhere(
		`WHERE iso2_code = $1 AND deleted_at IS NULL`,
		iso2Code)
}

// GetByISO3Code returns a country's details by 3-letters country code.
func (instance *XCountryDAO) GetByISO3Code(iso3Code string) (model.XCountry, error) {
	return instance.getOneWhere(
		`WHERE iso3_code = $1 AND deleted_at IS NULL`,
		iso3Code)
}
