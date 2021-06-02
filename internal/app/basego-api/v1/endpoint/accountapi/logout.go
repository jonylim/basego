/**
 * @api           {post} /v1/account/logout Logout
 * @apiVersion    1.0.0
 * @apiName       Logout
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Log out of an account session.
 *
 * @apiSuccess {boolean} success       If the logout is successful.
 * @apiSuccess {string}  message       The message.
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
 *         "success": true,
 *         "message": "You have logged out"
 *       }
 *     }
 *
 * @apiUse   ErrorAccountHeaderValidationFailed
 */

package accountapi

import (
	"net/http"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// LogoutResponseData represents response data of Account API "Logout".
type LogoutResponseData struct {
	api.ResponseData
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Logout deletes the account session.
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.Logout")

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Delete the account session from database.
	sessionDB := dao.NewCstAccountSessionDAO()
	if _, err = sessionDB.DeleteSessionByID(tx, ctx.AccountSession.ID, true); err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Delete the account session's tokens from database.
	if _, err = sessionDB.DeleteSessionTokenBySessionID(tx, ctx.AccountSession.ID); err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
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

	go func(sessionID int64) {
		// Get the Redis connection and defer closing connection.
		redisConn := redis.GetConnection()
		defer redisConn.Close()

		// Delete the account session and tokens from Redis, if any.
		sessionStore := redisstore.NewCstAccountSessionStore(redisConn)
		tokenStore := redisstore.NewCstAccountSessionTokenStore(redisConn)
		sessionStore.SaveNilByID(sessionID)
		tokenStore.SaveNilBySessionID(sessionID)
	}(ctx.AccountSession.ID)

	// Return the result.
	data := LogoutResponseData{
		Success: true,
		Message: "You have logged out",
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
