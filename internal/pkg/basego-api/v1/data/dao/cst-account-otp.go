package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// CstAccountOTPDAO manages database operations for customer account OTP data.
type CstAccountOTPDAO struct {
	dao
	selectColumns string
}

// NewCstAccountOTPDAO returns new instance of CstAccountOTPDAO.
func NewCstAccountOTPDAO() *CstAccountOTPDAO {
	return &CstAccountOTPDAO{
		dao: dao{db.Get(), false},
		selectColumns: `
				id, account_id, key, code, action, method,
				COALESCE(email, '') AS email, country_id, country_calling_code,
				COALESCE(phone, '') AS phone, COALESCE(phone_with_code, '') AS phone_with_code,
				` + sqlTimestampToUnixMilliseconds("expiry_time") + ` AS expiry_time,
				send_count, attempt_count, is_verified,
				` + sqlTimestampToUnixMilliseconds("created_at") + ` AS created_time,
				` + sqlTimestampToUnixMilliseconds("updated_at") + ` AS updated_time,
				` + sqlTimestampToUnixMilliseconds("deleted_at") + ` AS deleted_time`,
	}
}

func (instance *CstAccountOTPDAO) scanRow(row *sql.Row) (model.CstAccountOTP, error) {
	var data model.CstAccountOTP
	err := row.Scan(
		&data.ID, &data.AccountID, &data.Key, &data.Code, &data.Action, &data.Method,
		&data.Email, &data.CountryID, &data.CountryCallingCode,
		&data.Phone, &data.PhoneWithCode,
		&data.ExpiryTime, &data.SendCount, &data.AttemptCount, &data.IsVerified,
		&data.CreatedTime, &data.UpdatedTime, &data.DeletedTime)
	return data, err
}

func (instance *CstAccountOTPDAO) getWhere(sqlWhere string, params ...interface{}) (model.CstAccountOTP, error) {
	row := instance.db.QueryRow(`SELECT `+instance.selectColumns+`
			FROM tb_t_cst_account_otp
			`+sqlWhere, params...)
	data, err := instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
	}
	return data, err
}

// GetOTPByID returns an OTP's details by OTP ID.
func (instance *CstAccountOTPDAO) GetOTPByID(id int64) (model.CstAccountOTP, error) {
	var sqlWhereDeleted string
	if !instance.withDeleted {
		sqlWhereDeleted = `AND deleted_at IS NULL`
	}
	var data model.CstAccountOTP
	row := instance.db.QueryRow(`
			SELECT `+instance.selectColumns+`
			FROM tb_t_cst_account_otp
			WHERE id = $1 `+sqlWhereDeleted,
		id)
	data, err := instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
	}
	return data, err
}

// GetActiveOTPByAccountAndAction returns an active OTP's details by account ID and action.
func (instance *CstAccountOTPDAO) GetActiveOTPByAccountAndAction(accountID int64, action string) (model.CstAccountOTP, error) {
	var data model.CstAccountOTP
	row := instance.db.QueryRow(`
			SELECT `+instance.selectColumns+`
			FROM tb_t_cst_account_otp
			WHERE account_id = $1
				AND action = $2
				AND expiry_time > CURRENT_TIMESTAMP
				AND is_verified = FALSE
				AND deleted_at IS NULL
			ORDER BY id DESC
			LIMIT 1
		`, accountID, action)
	data, err := instance.scanRow(row)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
	}
	return data, err
}

// IsEmailVerified checks if an email address is already verified.
func (instance *CstAccountOTPDAO) IsEmailVerified(email string) (bool, error) {
	var id int64
	err := instance.db.QueryRow(`SELECT id
			FROM tb_t_cst_account_otp
			WHERE email = $1
				AND is_verified = TRUE
			ORDER BY id DESC
			LIMIT 1
		`, email).
		Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return (id != 0 && err == nil), err
}

// IsPhoneVerified checks if a phone number is already verified.
func (instance *CstAccountOTPDAO) IsPhoneVerified(countryCallingCode, phone string) (bool, error) {
	var id int64
	err := instance.db.QueryRow(`SELECT id
			FROM tb_t_cst_account_otp
			WHERE phone_with_code = $1
				AND is_verified = TRUE
			ORDER BY id DESC
			LIMIT 1
		`, countryCallingCode+phone).
		Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return (id != 0 && err == nil), err
}

// InsertOTP inserts new record of an OTP to database. This method requires database transaction to be passed.
func (instance *CstAccountOTPDAO) InsertOTP(tx *sql.Tx, item model.CstAccountOTP) (id, createdMillis int64, err error) {
	err = tx.QueryRow(`INSERT INTO tb_t_cst_account_otp (
				account_id, key, code, action, method, email,
				country_id, country_calling_code, phone, phone_with_code,
				expiry_time, send_count
			) VALUES (
				$1, $2, $3, $4, $5, $6,
				$7, $8, $9, $10,
				TO_TIMESTAMP($11), $12
			)
			RETURNING id, `+sqlTimestampToUnixMilliseconds("created_at"),
		item.AccountID, item.Key, item.Code, item.Action, item.Method, item.Email,
		item.CountryID, item.CountryCallingCode, item.Phone, item.CountryCallingCode+item.Phone,
		item.ExpiryTime/1000, item.SendCount).
		Scan(&id, &createdMillis)
	if err != nil {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
	}
	return
}

// IncrementAttemptCountByID increments an OTP's attempt count by OTP ID.
// If the OTP's state failed to change, attemptCount returns 0.
func (instance *CstAccountOTPDAO) IncrementAttemptCountByID(tx *sql.Tx, id int64) (attemptCount int, err error) {
	err = tx.QueryRow(`UPDATE tb_t_cst_account_otp
			SET attempt_count = attempt_count + 1,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $1
				AND expiry_time > CURRENT_TIMESTAMP
				AND is_verified = FALSE
				AND deleted_at IS NULL
			RETURNING attempt_count
		`, id).
		Scan(&attemptCount)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
		} else {
			attemptCount = 0
			err = nil
		}
	}
	return
}

// SetVerified marks an OTP's status as verified, also incrementing its attempt count.
// If the OTP's state failed to change, attemptCount returns 0.
func (instance *CstAccountOTPDAO) SetVerified(tx *sql.Tx, id int64) (attemptCount int, isVerified bool, err error) {
	err = tx.QueryRow(`UPDATE tb_t_cst_account_otp
			SET attempt_count = attempt_count + 1,
				is_verified = TRUE,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $1 
				AND expiry_time > CURRENT_TIMESTAMP
				AND is_verified = FALSE
				AND deleted_at IS NULL
			RETURNING attempt_count, is_verified
		`, id).
		Scan(&attemptCount, &isVerified)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
		} else {
			attemptCount = 0
			isVerified = false
			err = nil
		}
	}
	return
}

// DeleteOTPByID deletes an active OTP's details by OTP ID.
func (instance *CstAccountOTPDAO) DeleteOTPByID(tx *sql.Tx, id int64) (bool, error) {
	result, err := tx.Exec(`UPDATE tb_t_cst_account_otp
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE id = $1
				AND deleted_at IS NULL
		`, id)
	if err != nil {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
		return false, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
		return false, err
	}
	return rowCount > 0, nil
}

// DeleteActiveOTPByAccountAndAction deletes an active OTP's details by account ID and action.
// If there is no active OTP deleted, lastSendCount returns 0.
func (instance *CstAccountOTPDAO) DeleteActiveOTPByAccountAndAction(tx *sql.Tx, accountID int64, action string) (deletedID int64, lastSendCount int, err error) {
	err = tx.QueryRow(`UPDATE tb_t_cst_account_otp
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE account_id = $1
				AND action = $2
				AND expiry_time > CURRENT_TIMESTAMP
				AND is_verified = FALSE
				AND deleted_at IS NULL
			RETURNING id, send_count
		`, accountID, action).
		Scan(&deletedID, &lastSendCount)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Fatal("CstAccountOTPDAO", logger.FromError(err))
		} else {
			lastSendCount = 0
			err = nil
		}
	}
	return
}
