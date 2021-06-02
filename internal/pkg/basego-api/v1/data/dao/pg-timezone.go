package dao

import (
	"context"
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// PgTimeZoneDAO manages database operations for time zones.
type PgTimeZoneDAO struct {
	dao
}

// NewPgTimeZoneDAO returns new instance of PgTimeZoneDAO.
func NewPgTimeZoneDAO() *PgTimeZoneDAO {
	return &PgTimeZoneDAO{
		dao: dao{db.Get(), false},
	}
}

func (instance *PgTimeZoneDAO) scanRow(row *sql.Row) (res model.PgTimeZone, err error) {
	err = row.Scan(&res.Name, &res.Abbrev, &res.UTCOffset, &res.IsDST)
	return
}

func (instance *PgTimeZoneDAO) scanRows(rows *sql.Rows) (items []model.PgTimeZone, err error) {
	items = make([]model.PgTimeZone, 0)
	for rows.Next() {
		var res model.PgTimeZone
		err = rows.Scan(&res.Name, &res.Abbrev, &res.UTCOffset, &res.IsDST)
		if err != nil {
			return
		}
		items = append(items, res)
	}
	return
}

func (instance *PgTimeZoneDAO) buildQuery(sqlWhere, sqlOrderBy, sqlLimit string) string {
	sqlQuery := `
			SELECT name, abbrev, utc_offset, is_dst
			FROM pg_timezone_names `
	if sqlWhere != "" {
		sqlQuery += sqlWhere
	}
	if sqlOrderBy != "" {
		sqlQuery += sqlOrderBy
	}
	if sqlLimit != "" {
		sqlQuery += sqlLimit
	}
	return sqlQuery
}

func (instance *PgTimeZoneDAO) getOneWhere(ctx context.Context, sqlWhere string, params ...interface{}) (model.PgTimeZone, error) {
	row := instance.db.QueryRowContext(ctx,
		instance.buildQuery(sqlWhere, "", ""),
		params...)
	res, err := instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("PgTimeZoneDAO", logger.FromError(err))
	}
	return res, err
}

func (instance *PgTimeZoneDAO) getListWhere(ctx context.Context, sqlWhere, sqlOrderBy, sqlLimit string, params ...interface{}) ([]model.PgTimeZone, error) {
	rows, err := instance.db.QueryContext(ctx,
		instance.buildQuery(sqlWhere, sqlOrderBy, sqlLimit),
		params...)
	if err != nil {
		logger.Fatal("PgTimeZoneDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	items, err := instance.scanRows(rows)
	if err != nil {
		logger.Fatal("PgTimeZoneDAO", logger.FromError(err))
	}
	return items, err
}

// GetAll returns the list of time zones.
func (instance *PgTimeZoneDAO) GetAll(ctx context.Context, search string) ([]model.PgTimeZone, error) {
	sqlOrderBy := `
			ORDER BY name ASC `
	sqlWhere := `
			WHERE name NOT LIKE 'Etc/%' `
	if search != "" {
		sqlWhere += `
				LOWER(name) LIKE LOWER($1) `
		return instance.getListWhere(ctx, sqlWhere, sqlOrderBy, "", "%"+search+"%")
	}
	return instance.getListWhere(ctx, sqlWhere, sqlOrderBy, "")
}

// GetByName returns a time zone by name.
func (instance *PgTimeZoneDAO) GetByName(ctx context.Context, name string) (model.PgTimeZone, error) {
	return instance.getOneWhere(ctx, `
			WHERE name NOT LIKE 'Etc/%'
				AND LOWER(name) = LOWER($1) `, name)
}
