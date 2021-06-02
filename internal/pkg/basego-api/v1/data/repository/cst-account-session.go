package repository

import (
	"database/sql"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"

	"github.com/gomodule/redigo/redis"
)

// CstAccountSessionRepo manages data operations for customer account sessions, especially cache operations.
type CstAccountSessionRepo struct {
	ErrNotFound error
	ErrDatabase error

	redisConn redis.Conn

	sessionStore *redisstore.CstAccountSessionStore
	tokenStore   *redisstore.CstAccountSessionTokenStore
}

// NewCstAccountSessionRepo returns new instance of CstAccountSessionRepo.
func NewCstAccountSessionRepo(redisConn redis.Conn) *CstAccountSessionRepo {
	return &CstAccountSessionRepo{
		ErrNotFound: errNotFound,
		ErrDatabase: errDatabase,

		redisConn:    redisConn,
		sessionStore: redisstore.NewCstAccountSessionStore(redisConn),
		tokenStore:   redisstore.NewCstAccountSessionTokenStore(redisConn),
	}
}

// GetSessionDetailsBySessionID returns customer account session's details and tokens by session ID.
func (instance *CstAccountSessionRepo) GetSessionDetailsBySessionID(sessionID int64) (model.CstAccountSession, model.CstAccountSessionToken, error) {
	// Get customer account session's details from Redis.
	session, err := instance.sessionStore.GetSessionByID(sessionID)
	if err != nil {
		// Get from database.
		return instance.getFromDB(sessionID)
	}

	// Get token details from Redis.
	token, err := instance.tokenStore.GetTokenBySessionID(sessionID)
	if err != nil {
		// Get from database.
		return instance.getFromDB(sessionID)
	}

	// Return the customer account session's details & tokens.
	if instance.sessionExists(session) {
		return session, token, nil
	}
	return session, token, instance.ErrNotFound
}

func (instance *CstAccountSessionRepo) getFromDB(sessionID int64) (model.CstAccountSession, model.CstAccountSessionToken, error) {
	da := dao.NewCstAccountSessionDAO()
	da.WithDeleted()
	session, token, err := da.GetSessionByID(sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Save nil to Redis.
			instance.sessionStore.SaveNilByID(sessionID)
			instance.tokenStore.SaveNilBySessionID(sessionID)
			return session, token, instance.ErrNotFound
		}
		return session, token, instance.ErrDatabase
	}
	// Save to Redis.
	instance.sessionStore.SaveSession(session)
	if token.ID != 0 {
		instance.tokenStore.SaveToken(token)
	} else {
		instance.tokenStore.SaveNilBySessionID(sessionID)
	}
	return session, token, err
}

func (instance *CstAccountSessionRepo) sessionExists(session model.CstAccountSession) bool {
	return !session.RedisNil && session.ID != 0 && session.AccountID != 0 && session.DeletedTime == 0
}

func (instance *CstAccountSessionRepo) tokenExists(token model.CstAccountSessionToken) bool {
	return !token.RedisNil && token.ID != 0 && token.SessionID != 0 && token.DeletedTime == 0
}
