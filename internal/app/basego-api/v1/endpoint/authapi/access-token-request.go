/**
 * @apiDefine SuccessAccessToken
 *
 * @apiSuccess {string}  accessToken                   The access token.
 * @apiSuccess {long}    accessTokenExpiry             The access token's expiry time, in Unix milliseconds.
 * @apiSuccess {string}  refreshToken                  The refresh token.
 * @apiSuccess {long}    refreshTokenExpiry            The refresh token's expiry time, in Unix milliseconds.
 */

/**
 * @api           {post} /v1/auth/access_token/request Request Access Token
 * @apiVersion    1.0.0
 * @apiName       AccessToken_Request
 * @apiGroup      AuthAPI
 * @apiPermission client
 *
 * @apiDescription Request access token using the provided credentials.
 *
 * The credential is passed via request header `Authorization` using the following format:
 *
 * ```
 * Authorization: <type> <credentials>
 * ```
 *
 * The following authorization types are supported:
 *
 * | type  | credentials              |
 * |-------|--------------------------|
 * | Basic | `base64(email:password)` |
 *
 * The user session starts from the time the access token is generated.
 * An access token is valid for 24 hours, after which it must be refreshed using a refresh token.
 *
 * > __Note__
 * >
 * > Each user can only have 1 active session per device (defined by header `Device-Identifier`).<br>
 *
 * @apiParamExample {json} Request Header Example:
 * Content-Type: application/json
 * API-Key: YjQzYjQ4NzQ1ZGZhMGU0NGsxOk16STVYekV1TVYvOWhjSEJmYTJWNVgybGtYMkZ1WkQvb3hOVE0yT1RrM09EYzNNakkwTnpJNA==
 * Authorization: Basic am9obkBkb2UuY29tOnNlY3JldA==
 * Device-Identifier: aaaa-bbbb-cccc-dddd
 * Device-Model: Google Chrome
 * Device-Platform: web
 * User-Agent: Google Chrome/12.1.14
 *
 * @apiUse SuccessAccessToken
 * @apiUse SuccessAccountProfile
 *
 * @apiSuccessExample {json} Success Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "error": {
 *         "code": "",
 *         "message": "",
 *         "field": ""
 *       },
 *       "data": {
 *         "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiIxYWM2NzI3ZjVjMGY2YTUwNTBhZWYzOWE2NDM5ZWMzOTllMWI0M2I1YTMyZGUwM2FhNjU2MGE2NTczMWM1ZDgzZmZiMzEzMjM5MjkxY2FmZDRiZjAzZGJhNjY1NTUwNzk0MzE0MWMxYjNhOTk3OTRlMjgxYzA4NTZmZjNhYjc4OSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NDIwODM3MiwianRpIjoiMTU2NDEyMTk3MmEzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ENmon7QaarCoPP3cP74kpwWNSkyXBV486VSTLwrmNCo",
 *         "accessTokenExpiry": 1564208372000,
 *         "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiJkMjlhNTIzZWZjNzQwNzYzNDdmM2FlOGIwOTdhNTUwM2RiZWU1YjA2ZGYxNjAwYWU2NTY4N2Y5NmFkOGNiYWRlODYzYTNlOTk0MzFiMDkxNzc0YjE1ZTNhODk1ZDkzM2Y0NDQ2NzNkNmJlNWRhZDA5M2EwYjAyMWMyMWNiNTdhNSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NjcxMzk3MiwianRpIjoiMTU2NDEyMTk3MnIzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ip_Uf2rMfgrwUB6oLJ2TVI-NT-9IpxxibdUittg3CKU",
 *         "refreshTokenExpiry": 1566713972000,
 *         "account": {
 *           "id": 8,
 *           "fullName": "Jony",
 *           "email": "jony@example.com",
 *           "isEmailVerified": true,
 *           "countryID": 0,
 *           "countryCallingCode": "",
 *           "phone": "",
 *           "phoneWithCode": "",
 *           "isPhoneVerified": false,
 *           "imageURL": {
 *             "thumbnail": "",
 *             "fullsize": ""
 *           },
 *           "lastLoginTime": 1564121972641,
 *           "lastActivityTime": 1564121972641,
 *           "requireChangePassword": false,
 *           "createdTime": 1563868799147,
 *           "updatedTime": 1563880378559,
 *           "deletedTime": 0
 *         }
 *       }
 *     }
 *
 * @apiUse   ErrorAuthHeaderValidationFailed
 * @apiError AuthorizationFormatInvalid The header `Authorization` format is invalid.
 * @apiError CredentialsInvalid         The credentials is invalid.
 * @apiError AccountNotFound            The account is not found.
 * @apiError AccountNotVerified         The account is not verified yet.
 *
 * @apiErrorExample {json} AuthorizationFormatInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40102",
 *         "message": "Authorization type is invalid or not supported",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} CredentialsInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40103",
 *         "message": "The email and password does not match",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} AccountNotFound:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40106",
 *         "message": "Account is not found",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} AccountNotVerified:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40107",
 *         "message": "Your account has not been verified yet",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 */

package authapi

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/app/basego-api/v1/token/accesstoken"
	"github.com/jonylim/basego/internal/app/basego-api/v1/token/refreshtoken"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/crypto/password"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// AccessTokenRequestResponseData represents response data of Auth API "Request Access Token".
type AccessTokenRequestResponseData struct {
	api.ResponseData
	AccessToken        string           `json:"accessToken"`
	AccessTokenExpiry  int64            `json:"accessTokenExpiry"`
	RefreshToken       string           `json:"refreshToken"`
	RefreshTokenExpiry int64            `json:"refreshTokenExpiry"`
	Account            model.CstAccount `json:"account"`
}

// AccessTokenRequest requests access token using the given credentials.
func AccessTokenRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: authapi.AccessTokenRequest")

	// Get credentials from header.
	if ctx.ReqHeader.Authorization == "" {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationEmpty, "Authorization is required")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}
	parts := strings.SplitN(ctx.ReqHeader.Authorization, " ", 2)
	if len(parts) < 2 {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationFormatInvalid, "Authorization format is invalid")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}
	switch parts[0] {
	case "Basic":
	default:
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationFormatInvalid, "Authorization type is invalid or not supported")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}

	// Parse the credentials, get the email and password.
	bytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		logger.Error(ctx.ReqTag, err.Error())
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationTokenInvalid, "Failed to parse the credentials")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}
	parts = strings.SplitN(string(bytes), ":", 2)
	if len(parts) < 2 {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationTokenInvalid, "The credentials are invalid")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}
	email, pwd := parts[0], parts[1]
	if email == "" || pwd == "" {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationTokenInvalid, "The credentials are invalid")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}
	email = strings.ToLower(email)

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Get the account data.
	accRepo := repository.NewCstAccountRepo(redisConn)
	account, err := accRepo.GetByEmail(email)
	if err != nil {
		if err == accRepo.ErrNotFound {
			msg := "Account is not found"
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationUserNotFound, msg)
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		} else {
			if err == accRepo.ErrDatabase {
				err = errDatabase
			} else {
				err = errInternal
			}
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		}
		return
	}

	// Validate the password.
	if check := password.HashWithSalt(pwd, account.PasswordSalt); check != account.Password {
		msg := "The email and password does not match"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationTokenInvalid, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}

	// Check if the account has been verified.
	if !account.IsEmailVerified {
		msg := "Your account has not been verified yet"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationUserNotVerified, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	deviceID := ctx.ReqHeader.DeviceID
	ipAddr := api.GetClientIPAddress(r)

	var deletedSessionIDs []int64
	sessionDAO := dao.NewCstAccountSessionDAO()
	if deviceID != "" {
		// Delete existing customer account sessions from the device.
		// There can only be 1 customer account sessions per device.
		deletedSessionIDs, err = sessionDAO.DeleteSessionsByDevice(tx, deviceID)
		if err != nil {
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
			return
		}
	}

	// Save the new customer account session to database.
	sessionID, err := sessionDAO.InsertSession(tx, account.ID, ctx.ReqHeader.DevicePlatform, ctx.ReqHeader.DeviceModel, deviceID, ctx.ReqHeader.UserAgent, ipAddr)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Update the account's last login time.
	now := time.Now()
	nowSeconds := now.Unix()
	nowMillis := helper.UnixMillisecond(now)
	dao.NewCstAccountDAO().UpdateLastLogin(tx, account.ID, now)

	// Generate new access token & refresh tokens, calculate expiry times.
	accessTokenStr := accesstoken.GenerateAccessToken(sessionID)
	refreshTokenStr := refreshtoken.GenerateRefreshToken(sessionID)
	accessExpiryTime := now.Add(accesstoken.TokenTTL * time.Second)
	accessExpirySeconds := accessExpiryTime.Unix()
	accessExpiryMillis := accessExpirySeconds * 1000
	refreshExpiryTime := now.Add(refreshtoken.TokenTTL * time.Second)
	refreshExpirySeconds := refreshExpiryTime.Unix()
	refreshExpiryMillis := refreshExpirySeconds * 1000

	// Save the tokens to database.
	tokenID, err := sessionDAO.InsertSessionToken(tx, sessionID, accessTokenStr, accessExpiryMillis, refreshTokenStr, refreshExpiryMillis)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Generate JWT for the access token and refresh token.
	accessTokenJWT, err := accesstoken.GenerateJWT(tokenID, accessTokenStr, nowSeconds, accessExpirySeconds, sessionID, account.ID)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errInternal.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	refreshTokenJWT, err := refreshtoken.GenerateJWT(tokenID, refreshTokenStr, nowSeconds, refreshExpirySeconds, sessionID, account.ID)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errInternal.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Commit database transaction.
	err = tx.Commit()
	if err != nil {
		logger.Fatal("tx.Commit", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	session := model.CstAccountSession{
		ID:          sessionID,
		AccountID:   account.ID,
		Platform:    ctx.APIKey.AppPlatform,
		DeviceModel: ctx.ReqHeader.DeviceModel,
		DeviceID:    ctx.ReqHeader.DeviceID,
		UserAgent:   ctx.ReqHeader.UserAgent,
		IPAddress:   ipAddr,
		CreatedTime: nowMillis,
	}
	sessionToken := model.CstAccountSessionToken{
		ID:                 tokenID,
		SessionID:          sessionID,
		AccessToken:        accessTokenStr,
		AccessTokenExpiry:  accessExpiryMillis,
		RefreshToken:       refreshTokenStr,
		RefreshTokenExpiry: refreshExpiryMillis,
		CreatedTime:        nowMillis,
	}

	accStore := redisstore.NewCstAccountStore(redisConn)
	sessionStore := redisstore.NewCstAccountSessionStore(redisConn)
	tokenStore := redisstore.NewCstAccountSessionTokenStore(redisConn)

	// Delete old customer account sessions and tokens from Redis, if any.
	if deletedSessionIDs != nil && len(deletedSessionIDs) != 0 {
		sessionStore.SaveNilByIDs(deletedSessionIDs)
		tokenStore.SaveNilBySessionIDs(deletedSessionIDs)
	}

	account.LastLoginTime = session.CreatedTime
	account.LastActivityTime = session.CreatedTime

	// Save to Redis.
	accStore.Save(account)
	sessionStore.SaveSession(session)
	tokenStore.SaveToken(sessionToken)

	// Return the access token & refresh token.
	data := AccessTokenRequestResponseData{
		AccessToken:        accessTokenJWT,
		AccessTokenExpiry:  accessExpiryMillis,
		RefreshToken:       refreshTokenJWT,
		RefreshTokenExpiry: refreshExpiryMillis,
		Account:            account,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
