/**
 * @api           {post} /v1/auth/access_token/refresh Refresh Access Token
 * @apiVersion    1.0.0
 * @apiName       AccessToken_Refresh
 * @apiGroup      AuthAPI
 * @apiPermission client
 *
 * @apiDescription Request a new access token using a refresh token.
 *
 * The refresh token is passed via request header `Authorization` using the following format:
 *
 * ```
 * Authorization: Bearer <refresh_token>
 * ```
 *
 * A new refresh token will be generated and returned along with the new access token.
 * An access token is valid for 24 hours, after which it must be refreshed using a refresh token.
 * The old access token and refresh token will no longer be usable.
 *
 * This API has the same response structure as API [Request Access Token](#api-AuthAPI-AccessToken_Request).
 *
 * @apiParamExample {json} Request Header Example:
 * Content-Type: application/json
 * API-Key: YjQzYjQ4NzQ1ZGZhMGU0NGsxOk16STVYekV1TVYvOWhjSEJmYTJWNVgybGtYMkZ1WkQvb3hOVE0yT1RrM09EYzNNakkwTnpJNA==
 * Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiJkMjlhNTIzZWZjNzQwNzYzNDdmM2FlOGIwOTdhNTUwM2RiZWU1YjA2ZGYxNjAwYWU2NTY4N2Y5NmFkOGNiYWRlODYzYTNlOTk0MzFiMDkxNzc0YjE1ZTNhODk1ZDkzM2Y0NDQ2NzNkNmJlNWRhZDA5M2EwYjAyMWMyMWNiNTdhNSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NjcxMzk3MiwianRpIjoiMTU2NDEyMTk3MnIzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ip_Uf2rMfgrwUB6oLJ2TVI-NT-9IpxxibdUittg3CKU
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
 *         "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiI3OWFlMmIzZmVlZGYyOTM1ZDNjODdkYWE5OGE0YjJjNTU1NmUyNjBjZDAxNjk3M2EyMjczMDA1MjhlZGQ4MGQyOTgxNWYzNWRhODk1ZTNiZDk3NTBlYWE5NTk0NGVlNWU0OGMwZGQ0NzE2MzY5OTkyYmM5MjE5Y2UzM2ExNmQwOCIsInRpZCI6Nywic2lkIjo0LCJ1aWQiOjgsImV4cCI6MTU2NDMwNTA1NiwianRpIjoiMTU2NDIxODY1NmE3IiwiaWF0IjoxNTY0MjE4NjU2LCJpc3MiOiJjb25zb2xlIn0.lLEwnoIDuAzGASvnLlZE581ljShXYPbWxtl6If1NXlc",
 *         "accessTokenExpiry": 1564305056000,
 *         "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiI4MjJmMDAxMDhkYzE4NmM1NzI3ZWI2OGExZTY5NGE2MWUxMTc2MTBkZDkzYTg0NDUwNWI5NDg4MmU2OGE5YzljMmVlNWFiMDRiODBmZDZkYjY4ZDUzMzVlNmFhNjJjM2UxNzA4NjViOTEwMzE4ODU1MTdhZmM5OGQ0YmJiZDg2YyIsInRpZCI6Nywic2lkIjo0LCJ1aWQiOjgsImV4cCI6MTU2NjgxMDY1NiwianRpIjoiMTU2NDIxODY1NnI3IiwiaWF0IjoxNTY0MjE4NjU2LCJpc3MiOiJjb25zb2xlIn0.tn5HHwENMmNUj4KwLXW3KPkMFljfsvBaTmy82vsMoeY",
 *         "refreshTokenExpiry": 1566810656000,
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
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError AuthorizationFormatInvalid The header `Authorization` format is invalid.
 * @apiError RefreshTokenInvalid        The refresh token is invalid, or has already been used.
 * @apiError RefreshTokenExpired        The refresh token has expired.
 * @apiError DeviceInvalid              The refresh token does not belong to the device.
 * @apiError AccountNotFound            The refresh token is valid, but the associated account is not found.
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
 * @apiErrorExample {json} RefreshTokenInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40103",
 *         "message": "Token is not found",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} RefreshTokenExpired:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40104",
 *         "message": "Refresh token is expired",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} DeviceInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 401,
 *       "error": {
 *         "code": "40105",
 *         "message": "Refresh token does not belong to the device",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 */

package authapi

import (
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
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// AccessTokenRefreshData represents response data of Auth API "Refresh Access Token".
type AccessTokenRefreshData struct {
	AccessTokenRequestResponseData
}

// AccessTokenRefresh requests a new access token using a refresh token.
func AccessTokenRefresh(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: authapi.AccessTokenRefresh")

	// Get refresh token from header.
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
	} else if parts[0] != "Bearer" {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationFormatInvalid, "Authorization type is invalid or not supported")
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}

	// Parse the refresh token, get the token details.
	claims, code, err := refreshtoken.ParseJWT(parts[1])
	if err != nil {
		var errCode string
		switch code {
		case refreshtoken.ErrParseFailed:
			errCode = errcode.AuthorizationTokenInvalid
		case refreshtoken.ErrTokenExpired:
			errCode = errcode.AuthorizationTokenExpired
		}
		response := api.NewAPIResponseWithError(ctx.ReqID, errCode, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
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
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.AuthorizationTokenInvalid, msg)
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		} else {
			if err == sessionRepo.ErrDatabase {
				err = errDatabase
			} else {
				err = errInternal
			}
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		}
		return
	}

	// Validate the refresh token.
	deviceID := ctx.ReqHeader.DeviceID
	if code, err := claims.ValidateState(session, sessionToken, ctx.APIKey, deviceID); err != nil {
		var errCode string
		switch code {
		case refreshtoken.ErrTokenInvalid:
			errCode = errcode.AuthorizationTokenInvalid
		case refreshtoken.ErrDeviceInvalid:
			errCode = errcode.AuthorizationNotTokenOwner
		case refreshtoken.ErrNotOwner:
			errCode = errcode.AuthorizationNotTokenOwner
		case refreshtoken.ErrTokenExpired:
			errCode = errcode.AuthorizationTokenExpired
		}
		response := api.NewAPIResponseWithError(ctx.ReqID, errCode, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.Unauthorized)
		return
	}

	// Get the account data.
	accRepo := repository.NewCstAccountRepo(redisConn)
	account, err := accRepo.GetByID(session.AccountID)
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

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Update the account's last activity time.
	now := time.Now()
	nowSeconds := now.Unix()
	nowMillis := helper.UnixMillisecond(now)
	dao.NewCstAccountDAO().UpdateLastActivity(tx, account.ID, now)
	account.LastActivityTime = nowMillis

	// Generate new access token & refresh tokens, calculate expiry times.
	accessTokenStr := accesstoken.GenerateAccessToken(session.ID)
	refreshTokenStr := refreshtoken.GenerateRefreshToken(session.ID)
	accessExpiryTime := now.Add(accesstoken.TokenTTL * time.Second)
	accessExpirySeconds := accessExpiryTime.Unix()
	accessExpiryMillis := accessExpirySeconds * 1000
	refreshExpiryTime := now.Add(refreshtoken.TokenTTL * time.Second)
	refreshExpirySeconds := refreshExpiryTime.Unix()
	refreshExpiryMillis := refreshExpirySeconds * 1000

	// Save the new tokens to database.
	sessionDAO := dao.NewCstAccountSessionDAO()
	tokenID, err := sessionDAO.InsertSessionToken(tx, session.ID, accessTokenStr, accessExpiryMillis, refreshTokenStr, refreshExpiryMillis)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Delete old tokens from database.
	if _, err = sessionDAO.DeleteSessionTokenByID(tx, claims.TokenID, session.ID); err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Generate JWT for the access token and refresh token.
	accessTokenJWT, err := accesstoken.GenerateJWT(tokenID, accessTokenStr, nowSeconds, accessExpirySeconds, session.ID, account.ID)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errInternal.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	refreshTokenJWT, err := refreshtoken.GenerateJWT(tokenID, refreshTokenStr, nowSeconds, refreshExpirySeconds, session.ID, account.ID)
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

	sessionToken = model.CstAccountSessionToken{
		ID:                 tokenID,
		SessionID:          session.ID,
		AccessToken:        accessTokenStr,
		AccessTokenExpiry:  accessExpiryMillis,
		RefreshToken:       refreshTokenStr,
		RefreshTokenExpiry: refreshExpiryMillis,
		CreatedTime:        nowMillis,
	}

	// Save to Redis.
	redisstore.NewCstAccountStore(redisConn).Save(account)
	redisstore.NewCstAccountSessionTokenStore(redisConn).SaveToken(sessionToken)

	// Return the access token & refresh token.
	data := AccessTokenRefreshData{
		AccessTokenRequestResponseData{
			AccessToken:        accessTokenJWT,
			AccessTokenExpiry:  accessExpiryMillis,
			RefreshToken:       refreshTokenJWT,
			RefreshTokenExpiry: refreshExpiryMillis,
			Account:            account,
		},
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
