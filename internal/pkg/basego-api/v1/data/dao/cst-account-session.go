package dao

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// CstAccountSessionDAO manages database operations for customer account sessions.
type CstAccountSessionDAO struct {
	dao
}

// NewCstAccountSessionDAO returns new instance of CstAccountSessionDAO.
func NewCstAccountSessionDAO() *CstAccountSessionDAO {
	return &CstAccountSessionDAO{
		dao: dao{db.Get(), false},
	}
}

// GetSessionByID returns a customer account session's details and tokens by session ID.
func (instance *CstAccountSessionDAO) GetSessionByID(sessionID int64) (model.CstAccountSession, model.CstAccountSessionToken, error) {
	var s model.CstAccountSession
	var t model.CstAccountSessionToken
	var (
		tokenID       sql.NullInt64
		accessToken   sql.NullString
		accessExpiry  int64
		refreshToken  sql.NullString
		refreshExpiry int64
		tokenCreated  int64
		tokenDeleted  int64
	)
	var sqlWhereDeleted string
	if !instance.withDeleted {
		sqlWhereDeleted = `AND s.deleted_at IS NULL`
	}
	err := instance.db.QueryRow(`SELECT
				s.id, s.account_id, s.platform, s.device_model, s.device_id, s.user_agent, s.ip_address,
				`+sqlTimestampToUnixMilliseconds("s.logout_time")+` AS logout_time,
				`+sqlTimestampToUnixMilliseconds("s.created_at")+` AS created_time,
				`+sqlTimestampToUnixMilliseconds("s.updated_at")+` AS updated_time,
				`+sqlTimestampToUnixMilliseconds("s.deleted_at")+` AS deleted_time,
				t.id, t.access_token, `+sqlTimestampToUnixMilliseconds("t.access_token_expiry_time")+` AS access_token_expiry,
				t.refresh_token, `+sqlTimestampToUnixMilliseconds("t.refresh_token_expiry_time")+` AS refresh_token_expiry,
				`+sqlTimestampToUnixMilliseconds("t.created_at")+` AS token_created_time,
				`+sqlTimestampToUnixMilliseconds("t.deleted_at")+` AS token_deleted_time
			FROM tb_t_cst_account_session s
			LEFT JOIN tb_t_cst_account_session_token t
				ON t.session_id = s.id
				AND t.deleted_at IS NULL
			WHERE s.id = $1 `+sqlWhereDeleted+`
			ORDER BY t.id DESC
			LIMIT 1
		`, sessionID).
		Scan(&s.ID, &s.AccountID, &s.Platform, &s.DeviceModel, &s.DeviceID, &s.UserAgent, &s.IPAddress,
			&s.LogoutTime, &s.CreatedTime, &s.UpdatedTime, &s.DeletedTime,
			&tokenID, &accessToken, &accessExpiry, &refreshToken, &refreshExpiry, &tokenCreated, &tokenDeleted)
	if err != nil && err != sql.ErrNoRows {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return s, t, err
	}
	if tokenID.Valid {
		t = model.CstAccountSessionToken{
			ID:                 tokenID.Int64,
			SessionID:          s.ID,
			AccessToken:        accessToken.String,
			AccessTokenExpiry:  accessExpiry,
			RefreshToken:       refreshToken.String,
			RefreshTokenExpiry: refreshExpiry,
			CreatedTime:        tokenCreated,
			DeletedTime:        tokenDeleted,
		}
	}
	return s, t, err
}

// InsertSession inserts new record of customer account session to database. This method requires database transaction to be passed.
func (instance *CstAccountSessionDAO) InsertSession(tx *sql.Tx, accountID int64, platform, deviceModel, deviceID, userAgent, ipAddress string) (int64, error) {
	var id int64
	err := tx.QueryRow(`INSERT INTO tb_t_cst_account_session
			(account_id, platform, device_model, device_id, user_agent, ip_address)
			VALUES($1, $2, $3, $4, $5, $6) RETURNING id
		`, accountID, platform, deviceModel, deviceID, userAgent, ipAddress,
	).Scan(&id)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
	}
	return id, err
}

// InsertSessionToken inserts new access & refresh token for a customer account session.
// The token expiry time is in milliseconds.
// This method requires database transaction to be passed.
func (instance *CstAccountSessionDAO) InsertSessionToken(tx *sql.Tx, sessionID int64, accessToken string, accessTokenExpiry int64, refreshToken string, refreshTokenExpiry int64) (int64, error) {
	var id int64
	err := tx.QueryRow(`INSERT INTO tb_t_cst_account_session_token
			(session_id, access_token, access_token_expiry_time, refresh_token, refresh_token_expiry_time)
			VALUES ($1, $2, TO_TIMESTAMP($3), $4, TO_TIMESTAMP($5)) RETURNING id
		`, sessionID, accessToken, accessTokenExpiry/1000, refreshToken, refreshTokenExpiry/1000,
	).Scan(&id)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
	}
	return id, err
}

// DeleteSessionByID deletes a customer account session by session ID.
func (instance *CstAccountSessionDAO) DeleteSessionByID(tx *sql.Tx, sessionID int64, isLogout bool) (bool, error) {
	sqlUpdate := `UPDATE tb_t_cst_account_session `
	if isLogout {
		sqlUpdate += `
			SET logout_time = CURRENT_TIMESTAMP,
				deleted_at = CURRENT_TIMESTAMP `
	} else {
		sqlUpdate += `
			SET deleted_at = CURRENT_TIMESTAMP `
	}
	result, err := tx.Exec(sqlUpdate+`
			WHERE id = $1 
				AND deleted_at IS NULL 
		`, sessionID)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	return rowCount > 0, nil
}

// DeleteSessionsByDevice deletes active customer account sessions for a device,
// returning an array of the deleted customer account session IDs.
func (instance *CstAccountSessionDAO) DeleteSessionsByDevice(tx *sql.Tx, deviceID string) ([]int64, error) {
	if deviceID == "" {
		return nil, nil
	}
	rows, err := tx.Query(`UPDATE tb_t_cst_account_session
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE device_id = $1
				AND deleted_at IS NULL
			RETURNING id
		`, deviceID)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	sessionIDs := make([]int64, 0)
	for rows.Next() {
		var sid int64
		err = rows.Scan(&sid)
		if err != nil {
			logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
			return sessionIDs, nil
		}
		sessionIDs = append(sessionIDs, sid)
	}
	return sessionIDs, nil
}

// DeleteSessionsByAccountAndDevice deletes a customer account's active customer account sessions for a device,
// returning an array of the deleted customer account session IDs.
func (instance *CstAccountSessionDAO) DeleteSessionsByAccountAndDevice(tx *sql.Tx, accountID int64, deviceID string) ([]int64, error) {
	rows, err := tx.Query(`UPDATE tb_t_cst_account_session
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE account_id = $1
				AND device_id = $2
				AND deleted_at IS NULL
			RETURNING id
		`, accountID, deviceID)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return nil, err
	}
	defer rows.Close()
	sessionIDs := make([]int64, 0)
	for rows.Next() {
		var sid int64
		err = rows.Scan(&sid)
		if err != nil {
			logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
			return sessionIDs, nil
		}
		sessionIDs = append(sessionIDs, sid)
	}
	return sessionIDs, nil
}

// DeleteSessionTokenByID deletes a customer account session's tokens by token ID and session ID.
func (instance *CstAccountSessionDAO) DeleteSessionTokenByID(tx *sql.Tx, tokenID, sessionID int64) (bool, error) {
	result, err := tx.Exec(`UPDATE tb_t_cst_account_session_token 
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE id = $1 
				AND session_id = $2 
				AND deleted_at IS NULL
		`, tokenID, sessionID)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	return rowCount > 0, nil
}

// DeleteSessionTokenBySessionID deletes a customer account session's tokens by session ID.
func (instance *CstAccountSessionDAO) DeleteSessionTokenBySessionID(tx *sql.Tx, sessionID int64) (bool, error) {
	result, err := tx.Exec(`UPDATE tb_t_cst_account_session_token 
			SET deleted_at = CURRENT_TIMESTAMP
			WHERE session_id = $1 
				AND deleted_at IS NULL
		`, sessionID)
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		logger.Fatal("CstAccountSessionDAO", logger.FromError(err))
		return false, err
	}
	return rowCount > 0, nil
}
