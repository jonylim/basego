package requestvalidator

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/app/basego-api/v1/token/accesstoken"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

const tag = "requestvalidator"
const apiDomain = "customer"

var (
	errDatabase = errors.New("An error occurred while processing your request")
	errInternal = errors.New("An error occurred while processing your request")
)

// Validator manages validation for incoming requests.
type Validator struct {
	w     http.ResponseWriter
	r     *http.Request
	ctx   context.Context
	reqID string
}

// New returns Validator for an incoming requests.
func New(w http.ResponseWriter, r *http.Request, reqID string) Validator {
	return Validator{w, r, r.Context(), reqID}
}

func (v Validator) sendAPIResponseWithError(statusCode int, errCode, errMsg string) {
	response := api.NewAPIResponseWithError(v.reqID, errCode, errMsg)
	api.SendResponseJSONWithStatusCode(v.w, response, statusCode)
	return
}

// ValidateAPIKey checks if an API key string is valid and returns the API key's details.
// The boolean is false if the API key validation fails and the request should not be processed any further.
func (v Validator) ValidateAPIKey(apiKeyStr, appIdentifier, appPlatform string) (apiKey model.XAPIKey, ok bool) {
	if apiKeyStr == "" {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyEmpty, "API-Key is required")
		return
	}

	// Parse the string.
	bytes, err := base64.StdEncoding.DecodeString(apiKeyStr)
	if err != nil {
		logger.Error(tag, "ValidateAPIKey: "+logger.FromError(err))
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyInvalid, "API-Key failed to parse")
		return
	}
	parts := strings.SplitN(string(bytes), ":", 2)
	if len(parts) < 2 {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyInvalid, "API-Key format is invalid")
		return
	}

	// Get the API key ID & secret.
	apiKeyID := parts[0]
	apiKeySecret := parts[1]

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Get the API key.
	apiKeyRepo := repository.NewAPIKeyRepo(redisConn, apiDomain)
	apiKey, err = apiKeyRepo.GetByAPIKeyID(apiKeyID)
	if err != nil {
		if err == apiKeyRepo.ErrNotFound {
			v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyNotFound, "API-Key is not found")
			return
		}
		v.sendAPIResponseWithError(httpstatus.InternalServerError, errcode.InternalAPIKeyValidationFailed, "An error occurred while validating API-Key")
		return
	}

	// Validate the API key secret.
	if apiKey.APIKeyID != apiKeyID || apiKey.APIKeySecret != apiKeySecret {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyInvalid, "API-Key is invalid")
		return
	}
	// Check if the device platform matches the API key's platform.
	if apiKey.AppPlatform != appPlatform {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyAppPlatformInvalid, "API-Key is invalid for the platform")
		return
	}
	// Check if the app identifier matches the API key's.
	if apiKey.AppIdentifier != "" && apiKey.AppIdentifier != appIdentifier {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyAppIdentifierInvalid, "API-Key is invalid for the app identifier")
		return
	}
	// Check the API key's expiry.
	now := helper.UnixMillisecond(time.Now())
	if now >= apiKey.ExpiryTime {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyExpired, "API-Key has expired")
		return
	}
	// Check if the API key is enabled.
	if !apiKey.IsEnabled {
		v.sendAPIResponseWithError(httpstatus.APIKeyInvalid, errcode.APIKeyDisabled, "API-Key is disabled")
		return
	}

	// Validation is successful.
	ok = true
	return
}

// ValidateAccessToken checks if an access token is valid and returns the account session and account's details.
// The boolean is false if the access token validation fails and the request should not be processed any further.
func (v Validator) ValidateAccessToken(authorization, deviceID string, apiKey model.XAPIKey) (model.CstAccountSession, model.CstAccount, bool) {
	var emptySession model.CstAccountSession
	var emptyAccount model.CstAccount

	// Get access token from header.
	if authorization == "" {
		v.sendAPIResponseWithError(httpstatus.Unauthorized, errcode.AuthorizationEmpty, "Authorization is required")
		return emptySession, emptyAccount, false
	}
	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) < 2 {
		v.sendAPIResponseWithError(httpstatus.Unauthorized, errcode.AuthorizationFormatInvalid, "Authorization format is invalid")
		return emptySession, emptyAccount, false
	} else if parts[0] != "Bearer" {
		v.sendAPIResponseWithError(httpstatus.Unauthorized, errcode.AuthorizationFormatInvalid, "Authorization type is invalid or not supported")
		return emptySession, emptyAccount, false
	}

	// Parse the access token, get the token details.
	claims, code, err := accesstoken.ParseJWT(parts[1])
	if err != nil {
		var errCode string
		switch code {
		case accesstoken.ErrParseFailed:
			errCode = errcode.AuthorizationTokenInvalid
		case accesstoken.ErrTokenExpired:
			errCode = errcode.AuthorizationTokenExpired
		}
		v.sendAPIResponseWithError(httpstatus.Unauthorized, errCode, err.Error())
		return emptySession, emptyAccount, false
	}

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Get account session's details by session ID.
	sessionRepo := repository.NewCstAccountSessionRepo(redisConn)
	session, sessionToken, err := sessionRepo.GetSessionDetailsBySessionID(claims.SessionID)
	if err != nil {
		if err == sessionRepo.ErrNotFound {
			msg := "Account session is not found"
			v.sendAPIResponseWithError(httpstatus.Unauthorized, errcode.AuthorizationTokenInvalid, msg)
		} else {
			if err == sessionRepo.ErrDatabase {
				err = errDatabase
			} else {
				err = errInternal
			}
			v.sendAPIResponseWithError(httpstatus.InternalServerError, errcode.Other, err.Error())
		}
		return emptySession, emptyAccount, false
	}

	// Validate the access token.
	if code, err := claims.ValidateState(session, sessionToken, apiKey, deviceID); err != nil {
		var errCode string
		switch code {
		case accesstoken.ErrTokenInvalid:
			errCode = errcode.AuthorizationTokenInvalid
		case accesstoken.ErrDeviceInvalid:
			errCode = errcode.AuthorizationNotTokenOwner
		case accesstoken.ErrNotOwner:
			errCode = errcode.AuthorizationNotTokenOwner
		case accesstoken.ErrTokenExpired:
			errCode = errcode.AuthorizationTokenExpired
		}
		v.sendAPIResponseWithError(httpstatus.Unauthorized, errCode, err.Error())
		return emptySession, emptyAccount, false
	}

	// Get the account's details.
	accRepo := repository.NewCstAccountRepo(redisConn)
	account, err := accRepo.GetByID(session.AccountID)
	if err != nil {
		if err == accRepo.ErrNotFound {
			msg := "Account is not found"
			v.sendAPIResponseWithError(httpstatus.Unauthorized, errcode.AuthorizationUserNotFound, msg)
		} else {
			if err == accRepo.ErrDatabase {
				err = errDatabase
			} else {
				err = errInternal
			}
			v.sendAPIResponseWithError(httpstatus.InternalServerError, errcode.Other, err.Error())
		}
		return emptySession, emptyAccount, false
	}

	// Update the account's last activity time.
	now := time.Now()
	go dao.NewCstAccountDAO().UpdateLastActivity(nil, account.ID, now)
	account.LastActivityTime = helper.UnixMillisecond(time.Now())
	accRepo.RedisStore().Save(account)

	// Return the account session and account's details.
	return session, account, true
}
