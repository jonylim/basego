package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// CstAccountTOSDAO manages database operations for customer account's Terms of Service status.
type CstAccountTOSDAO struct {
	dao
	selectColumns string
}

// NewCstAccountTOSDAO returns new instance of CstAccountTOSDAO.
func NewCstAccountTOSDAO() *CstAccountTOSDAO {
	return &CstAccountTOSDAO{
		dao: dao{db.Get(), false},
		selectColumns: `
				id, account_id,
				` + sqlTimestampToUnixMilliseconds("created_at") + ` AS created_time,
				` + sqlTimestampToUnixMilliseconds("updated_at") + ` AS updated_time,
				` + sqlTimestampToUnixMilliseconds("deleted_at") + ` AS deleted_time`,
	}
}

func (instance *CstAccountTOSDAO) scanRow(r SQLRowOrRows) (res model.CstAccountTOS, err error) {
	err = r.Scan(
		&res.ID, &res.AccountID,
		&res.CreatedTime, &res.UpdatedTime, &res.DeletedTime)
	return
}

func (instance *CstAccountTOSDAO) scanRows(rows *sql.Rows) ([]model.CstAccountTOS, error) {
	items := make([]model.CstAccountTOS, 0)
	for rows.Next() {
		res, err := instance.scanRow(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, res)
	}
	return items, nil
}

func (instance *CstAccountTOSDAO) getOneWhere(sqlWhere string, params ...interface{}) (res model.CstAccountTOS, err error) {
	row := instance.db.QueryRow(`
			SELECT `+instance.selectColumns+`
			FROM tb_m_cst_account_tos
			`+sqlWhere, params...)
	res, err = instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountTOSDAO", logger.FromError(err))
	}
	return
}

func (instance *CstAccountTOSDAO) getListWhere(sqlWhere, sqlOrderBy string, params ...interface{}) ([]model.CstAccountTOS, error) {
	rows, err := instance.db.Query(`SELECT `+instance.selectColumns+`
			FROM tb_m_cst_account_tos
			`+sqlWhere+`
			`+sqlOrderBy,
		params...)
	if err != nil {
		logger.Fatal("CstAccountTOSDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	items, err := instance.scanRows(rows)
	if err != nil {
		logger.Fatal("CstAccountTOSDAO", logger.FromError(err))
	}
	return items, err
}

// GetByID returns an account's TOS status by ID.
func (instance *CstAccountTOSDAO) GetByID(id int64) (model.CstAccountTOS, error) {
	sqlWhere := `WHERE id = $1
				AND deleted_at IS NULL`
	return instance.getOneWhere(sqlWhere, id)
}

// GetByAccountID returns an account's TOS status by account ID.
func (instance *CstAccountTOSDAO) GetByAccountID(accountID int64) (model.CstAccountTOS, error) {
	sqlWhere := `WHERE account_id = $1
				AND deleted_at IS NULL
			ORDER BY id DESC`
	return instance.getOneWhere(sqlWhere, accountID)
}

// Insert inserts an account's TOS status to database.
func (instance *CstAccountTOSDAO) Insert(tx *sql.Tx, accountID int64) (inserted model.CstAccountTOS, err error) {
	row := tx.QueryRow(`INSERT INTO tb_m_cst_account_tos (account_id)
			VALUES ($1)
			RETURNING `+instance.selectColumns,
		accountID)
	inserted, err = instance.scanRow(row)
	if err != nil {
		logger.Fatal("CstAccountTOSDAO", logger.FromError(err))
	}
	return
}
