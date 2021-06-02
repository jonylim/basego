package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/constant"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// CstAccountDAO manages database operations for customer accounts.
type CstAccountDAO struct {
	dao
	sqlSelectFrom string
}

// NewCstAccountDAO returns new instance of CstAccountDAO.
func NewCstAccountDAO() *CstAccountDAO {
	return &CstAccountDAO{
		dao: dao{db.Get(), false},
		sqlSelectFrom: `SELECT
				a.id, a.full_name, COALESCE(a.email, '') AS email, a.is_email_verified,
				a.country_id, a.country_calling_code,
				COALESCE(a.phone, '') AS phone, COALESCE(a.phone_with_code, '') AS phone_with_code, a.is_phone_verified,
				a.password, a.password_salt, a.use_2fa,
				f.filename AS photo_filename, f.storage AS photo_storage, f.is_encrypted AS photo_is_encrypted, 
				` + sqlTimestampToUnixMilliseconds("a.last_login_time") + ` AS last_login_time,
				` + sqlTimestampToUnixMilliseconds("a.last_activity_time") + ` AS last_activity_time,
				a.is_password_change_required,
				` + sqlTimestampToUnixMilliseconds("a.created_at") + ` AS created_time,
				` + sqlTimestampToUnixMilliseconds("a.updated_at") + ` AS updated_time,
				` + sqlTimestampToUnixMilliseconds("a.deleted_at") + ` AS deleted_time
			FROM tb_m_cst_account a
			LEFT JOIN tb_m_file f
				ON f.owner_type = '` + constant.FileOwnerTypeCstAccount + `'
				AND f.owner_id = a.id
				AND f.category = '` + constant.FileCategoryPhoto + `'
				AND f.filename = a.photo_filename
				AND f.deleted_at IS NULL 
			`,
	}
}

func (instance *CstAccountDAO) scanRow(r SQLRowOrRows) (res model.CstAccount, err error) {
	var photo cstAccountPhoto
	err = r.Scan(&res.ID, &res.FullName, &res.Email, &res.IsEmailVerified,
		&res.CountryID, &res.CountryCallingCode,
		&res.Phone, &res.PhoneWithCode, &res.IsPhoneVerified,
		&res.Password, &res.PasswordSalt, &res.Use2FA,
		&photo.Filename, &photo.Storage, &photo.IsEncrypted,
		&res.LastLoginTime, &res.LastActivityTime,
		&res.RequireChangePassword,
		&res.CreatedTime, &res.UpdatedTime, &res.DeletedTime)
	if err == nil && photo.Filename.String != "" {
		res.ImageURL = photo.ImageURL()
	}
	return
}

func (instance *CstAccountDAO) scanRows(rows *sql.Rows) ([]model.CstAccount, error) {
	items := make([]model.CstAccount, 0)
	for rows.Next() {
		res, err := instance.scanRow(rows)
		if err != nil {
			return items, err
		}
		items = append(items, res)
	}
	return items, nil
}

func (instance *CstAccountDAO) getOneWhere(sqlWhere string, params ...interface{}) (res model.CstAccount, err error) {
	row := instance.db.QueryRow(
		instance.sqlSelectFrom+
			sqlWhere,
		params...)
	res, err = instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
	}
	return
}

func (instance *CstAccountDAO) getListWhere(sqlWhere, sqlOrderBy, sqlLimit string, params ...interface{}) ([]model.CstAccount, error) {
	rows, err := instance.db.Query(
		instance.sqlSelectFrom+
			sqlWhere+
			sqlOrderBy+
			sqlLimit,
		params...)
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	items, err := instance.scanRows(rows)
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
	}
	return items, err
}

func (instance *CstAccountDAO) getByFieldAndValue(field string, value interface{}, ignoreCase bool) (res model.CstAccount, err error) {
	param := "$1"
	if ignoreCase {
		field = fmt.Sprintf("LOWER(a.%s)", field)
		param = "LOWER($1)"
	} else {
		field = "a." + field
	}
	sqlWhere := `WHERE ` + field + ` = ` + param
	if !instance.withDeleted {
		sqlWhere += `
				AND a.deleted_at IS NULL `
	}
	return instance.getOneWhere(sqlWhere, value)
}

// GetByID returns a customer account's details by ID.
func (instance *CstAccountDAO) GetByID(id int64) (model.CstAccount, error) {
	return instance.getByFieldAndValue("id", id, false)
}

// GetByEmail returns a customer account's details by email address.
func (instance *CstAccountDAO) GetByEmail(email string) (model.CstAccount, error) {
	return instance.getByFieldAndValue("email", email, true)
}

func (instance *CstAccountDAO) getByFieldAndValues(field string, values []interface{}, ignoreCase bool) ([]model.CstAccount, error) {
	if len(values) == 0 {
		return nil, errors.New("CstAccountDAO: values can't be empty")
	}
	argp := argPlaceholder{}
	valuePlaceholders := make([]string, len(values))
	for i := range values {
		valuePlaceholders[i] = argp.NextPlaceholder()
	}
	var sqlWhere string
	if ignoreCase {
		sqlWhere = fmt.Sprintf(`WHERE LOWER(a.%s) IN (%s)`,
			field,
			`LOWER(`+strings.Join(valuePlaceholders, "), LOWER(")+`)`)
	} else {
		sqlWhere = fmt.Sprintf(`WHERE a.%s IN (%s)`,
			field,
			strings.Join(valuePlaceholders, ", "))
	}
	if !instance.withDeleted {
		sqlWhere += `
				AND a.deleted_at IS NULL `
	}
	return instance.getListWhere(sqlWhere,
		"ORDER BY a.full_name, a.id",
		"",
		values...)
}

// GetByIDs returns a list of customer accounts' details by IDs.
func (instance *CstAccountDAO) GetByIDs(ids []int64, sortByParams bool) (items []model.CstAccount, err error) {
	values := make([]interface{}, len(ids))
	for i, id := range ids {
		values[i] = id
	}
	items, err = instance.getByFieldAndValues("id", values, false)
	if err == nil && sortByParams {
		x := 0
		for _, id := range ids {
			for i := x; i < len(items); i++ {
				if items[i].ID == id {
					if i != x {
						items[i], items[x] = items[x], items[i]
					}
					x++
					break
				}
			}
		}
	}
	return
}

// Insert inserts a new customer account to database.
func (instance *CstAccountDAO) Insert(tx *sql.Tx, account model.CstAccount) (insertedID, createdMillis int64, err error) {
	account.PhoneWithCode = account.CountryCallingCode + account.Phone
	err = tx.QueryRow(`INSERT INTO tb_m_cst_account (
				full_name, email, is_email_verified,
				country_id, country_calling_code, phone, phone_with_code, is_phone_verified,
				password, password_salt, use_2fa, is_password_change_required
			) VALUES (
				$1, $2, $3,
				$4, $5, $6, $7, $8,
				$9, $10, $11, $12
			) RETURNING id, `+sqlTimestampToUnixMilliseconds("created_at"),
		account.FullName, account.Email, account.IsEmailVerified,
		account.CountryID, account.CountryCallingCode, account.Phone, account.PhoneWithCode, account.IsPhoneVerified,
		account.Password, account.PasswordSalt, account.Use2FA, account.RequireChangePassword,
	).Scan(&insertedID, &createdMillis)
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
	}
	return
}

// ChangePassword updates a customer account's password hash & salt.
func (instance *CstAccountDAO) ChangePassword(tx *sql.Tx, accountID int64, passwordHash, passwordSalt string) (updated bool, err error) {
	var chkHash, chkSalt string
	err = tx.QueryRow(`UPDATE tb_m_cst_account
			SET password = $1,
				password_salt = $2,
				is_password_change_required = FALSE,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $3
				AND deleted_at IS NULL
			RETURNING password, password_salt
		`, passwordHash, passwordSalt, accountID).
		Scan(&chkHash, &chkSalt)
	if err == nil {
		if chkHash == passwordHash && chkSalt == passwordSalt {
			updated = true
		} else {
			logger.Error("CstAccountDAO", logger.FromError(fmt.Errorf("CstAccountDAO: no change (accountID = %v)", accountID)))
		}
	} else if err == sql.ErrNoRows {
		logger.Error("CstAccountDAO", logger.FromError(fmt.Errorf("CstAccountDAO: no rows affected (accountID = %v)", accountID)))
		err = nil
	} else {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
	}
	return
}

// UpdateLastLogin updates a customer account's last login and last activity time.
func (instance *CstAccountDAO) UpdateLastLogin(tx *sql.Tx, accountID int64, loginTime time.Time) (bool, error) {
	result, err := tx.Exec(`UPDATE tb_m_cst_account
			SET last_login_time = TO_TIMESTAMP($1),
				last_activity_time = TO_TIMESTAMP($1)
			WHERE id = $2
				AND deleted_at IS NULL
		`, loginTime.Unix(), accountID)
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
		return false, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
		return false, err
	}
	return rowCount > 0, nil
}

// UpdateLastActivity updates a customer account's last activity time.
func (instance *CstAccountDAO) UpdateLastActivity(tx *sql.Tx, accountID int64, lastActivityTime time.Time) (updated bool, err error) {
	var result sql.Result
	if tx != nil {
		result, err = tx.Exec(`UPDATE tb_m_cst_account
				SET last_activity_time = TO_TIMESTAMP($1)
				WHERE id = $2
					AND deleted_at IS NULL
			`, lastActivityTime.Unix(), accountID)
	} else {
		result, err = instance.db.Exec(`UPDATE tb_m_cst_account
				SET last_activity_time = TO_TIMESTAMP($1)
				WHERE id = $2
					AND deleted_at IS NULL
			`, lastActivityTime.Unix(), accountID)
	}
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
		return
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountDAO", logger.FromError(err))
		return
	}
	updated = rowCount > 0
	return
}

// SetVerifiedEmail marks a customer account's email address as verified.
func (instance *CstAccountDAO) SetVerifiedEmail(tx *sql.Tx, id int64) (isVerified bool, err error) {
	err = tx.QueryRow(`UPDATE tb_m_cst_account
			SET is_email_verified = TRUE,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $1 
				AND is_email_verified = FALSE
				AND deleted_at IS NULL
			RETURNING is_email_verified
		`, id).
		Scan(&isVerified)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Fatal("CstAccountDAO", logger.FromError(err))
		} else {
			isVerified = false
			err = nil
		}
	}
	return
}
