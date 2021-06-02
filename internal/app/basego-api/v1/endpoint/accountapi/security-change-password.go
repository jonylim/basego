/**
 * @api           {post} /v1/account/security/change_password Security - Change Password
 * @apiVersion    1.0.0
 * @apiName       ChangePassword
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Change the account's password.
 *
 * @apiParam {string} password    The account's current password, to verify the request.
 * @apiParam {string} newPassword The new password.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "password": "this_is_password",
 *       "newPassword": "this_is_new_password"
 *     }
 *
 * @apiSuccess {boolean} success If the password is changed successfully.
 * @apiSuccess {string}  message The message, if failed.
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
 *         "message": "Password changed successfully"
 *       }
 *     }
 *
 * @apiUse   ErrorAccountHeaderValidationFailed
 * @apiError PasswordInvalid          The current password is invalid.
 * @apiError NewPasswordFormatInvalid The new password format is invalid.
 *
 * @apiErrorExample {json} PasswordInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Current password is invalid",
 *         "field": "password"
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} NewPasswordFormatInvalid:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Password must contain at least 1 lowercase, uppercase, and special characters and 1 number",
 *         "field": "newPassword"
 *       },
 *       "data": {}
 *     }
 */

package accountapi

import (
	"encoding/json"
	"net/http"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
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

// SecurityChangePasswordRequestParam represents request body of Account API "Change Password".
type SecurityChangePasswordRequestParam struct {
	CurrentPassword string `json:"password"`
	NewPassword     string `json:"newPassword"`
}

// SecurityChangePasswordResponseData represents response data of Account API "Change Password".
type SecurityChangePasswordResponseData struct {
	api.ResponseData
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SecurityChangePassword changes the account's password.
func SecurityChangePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.SecurityChangePassword")

	var param SecurityChangePasswordRequestParam
	errReq := json.NewDecoder(r.Body).Decode(&param)
	if errReq != nil {
		logger.Error(ctx.ReqTag, errReq.Error())
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	var msg, field string
	if param.CurrentPassword == "" {
		msg = "Current password is required"
		field = "password"
	} else if chk := password.HashWithSalt(param.CurrentPassword, ctx.Account.PasswordSalt); chk != ctx.Account.Password {
		msg = "Current password is invalid"
		field = "password"
	} else if param.NewPassword == "" {
		msg = "New password is required"
		field = "newPassword"
	} else if err := helper.ValidatePasswordFormat(param.NewPassword); err != nil {
		msg = err.Error()
		field = "newPassword"
	}
	if msg != "" {
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	// Generate new password salt and hash the password.
	pwdHash, pwdSalt := password.Hash(param.NewPassword)

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Save the new password to database.
	success, err := dao.NewCstAccountDAO().ChangePassword(tx, ctx.Account.ID, pwdHash, pwdSalt)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	} else if !success {
		response := api.NewAPIResponse(ctx.ReqID)
		response.SetData(SecurityChangePasswordResponseData{
			Success: false,
			Message: "Failed to change password",
		})
		api.SendResponseJSON(w, response)
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

	go func(accountID int64) {
		// Get the Redis connection and defer closing connection.
		redisConn := redis.GetConnection()
		defer redisConn.Close()

		// Sync to Redis.
		repository.NewCstAccountRepo(redisConn).SyncByID(accountID)
	}(ctx.Account.ID)

	// Return the result.
	data := SecurityChangePasswordResponseData{
		Success: true,
		Message: "Password changed successfully",
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
