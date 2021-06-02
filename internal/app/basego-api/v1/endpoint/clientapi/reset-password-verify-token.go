/**
 * @api           {post} /v1/client/reset_password/verify_token Reset Password - Verify Token
 * @apiVersion    1.0.0
 * @apiName       ResetPassword_VerifyToken
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Verify a reset password token.
 *
 * @apiParam {long}   otpID   The OTP ID of the reset password token.
 * @apiParam {string} otpKey  The OTP key.
 * @apiParam {string} otpCode The OTP code.
 * @apiParam {string} email   The account's email address.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "otpID": 128,
 *       "otpKey": "49ff390684515734c4c645df3884ed7c",
 *       "otpCode": "123456",
 *       "email": "john@doe.com"
 *     }
 *
 * @apiSuccess {boolean} isValid If the reset password token is valid.
 * @apiSuccess {string}  message The message.
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
 *         "isValid": true,
 *         "message": ""
 *       }
 *     }
 *
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError ParamValidationFailed The parameter validation failed.
 *
 * @apiErrorExample {json} ParamValidationFailed:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Token is required",
 *         "field": "otpCode"
 *       },
 *       "data": {}
 *     }
 *
 */

package clientapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/token/otp"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// ResetPasswordVerifyTokenRequestParam represents request body of Client API "Reset Password - Verify Token".
type ResetPasswordVerifyTokenRequestParam struct {
	Email   string `json:"email"`
	OTPID   int64  `json:"otpID"`
	OTPKey  string `json:"otpKey"`
	OTPCode string `json:"otpCode"`
}

// ResetPasswordVerifyTokenResponseData represents response data of Client API "Reset Password - Verify Token".
type ResetPasswordVerifyTokenResponseData struct {
	api.ResponseData
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

// ResetPasswordVerifyToken verifies a reset password token.
func ResetPasswordVerifyToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.ResetPasswordVerifyToken")

	var param ResetPasswordVerifyTokenRequestParam
	errReq := json.NewDecoder(r.Body).Decode(&param)
	if errReq != nil {
		logger.Error(ctx.ReqTag, errReq.Error())
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	var msg, field string
	if param.OTPID == 0 {
		msg = "Token is required"
		field = "otpID"
	} else if param.OTPKey == "" {
		msg = "Token is required"
		field = "otpKey"
	} else if param.OTPCode == "" {
		msg = "Token is required"
		field = "otpCode"
	} else if param.Email == "" {
		msg = "Email address is required"
		field = "email"
	} else if err := helper.ValidateEmailFormat(param.Email); err != nil {
		msg = err.Error()
		field = "email"
	}
	if msg != "" {
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Check if the email address is registered.
	accRepo := repository.NewCstAccountRepo(redisConn)
	account, err := accRepo.GetByEmail(param.Email)
	if err != nil {
		if err == accRepo.ErrNotFound {
			response := api.NewAPIResponse(ctx.ReqID)
			response.SetData(ResetPasswordVerifyTokenResponseData{
				IsValid: false,
				Message: "Invalid request",
			})
			api.SendResponseJSON(w, response)
			return
		} else if err == accRepo.ErrDatabase {
			err = errDatabase
		} else {
			err = errInternal
		}
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Get active OTP's details.
	otpRepo := repository.NewCstAccountOTPRepo(redisConn)
	otpData, err := otpRepo.GetActiveOTPByAccountAndAction(account.ID, otp.ActionResetPassword)
	if err != nil {
		if err == otpRepo.ErrNotFound {
			msg = "There is no reset password request found, or the code has expired"
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
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

	// Validate the submitted OTP.
	isValidID := true
	isValidCode := false
	if otpData.ID != param.OTPID {
		msg = "Token is invalid"
		field = "otpID"
		isValidID = false
	} else if otpData.Key != param.OTPKey {
		msg = "Token is invalid"
		field = "otpKey"
	} else if otpData.Code != param.OTPCode {
		msg = "Token is incorrect"
		field = "otpCode"
	} else if otpData.Email != account.Email {
		msg = "Token is invalid"
		field = "email"
	} else {
		isValidCode = true
	}
	if isValidID && !isValidCode {
		// Begin database transaction.
		tx, err := db.Get().Begin()
		if err != nil {
			logger.Fatal("db.Begin", logger.FromError(err))
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
			return
		}
		defer tx.Rollback()

		// Increment the OTP's reset password attempt count.
		attemptCount, err := dao.NewCstAccountOTPDAO().IncrementAttemptCountByID(tx, otpData.ID)
		if err != nil {
			otpRepo.RedisStore().DeleteOTPByID(otpData.ID)
			otpRepo.RedisStore().DeleteOTPByAccountAndAction(otpData.AccountID, otpData.Action)
		} else {
			// Commit database transaction.
			err = tx.Commit()
			if err != nil {
				logger.Fatal("tx.Commit", logger.FromError(err))
			} else {
				// Update to Redis.
				otpData.AttemptCount = attemptCount
				otpData.UpdatedTime = helper.UnixMillisecond(time.Now())
				otpRepo.RedisStore().SaveOTPByAccountAndAction(otpData, emailResetPasswordTTL)
			}
		}
	}
	if msg != "" {
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	// Return the response.
	data := ResetPasswordVerifyTokenResponseData{
		IsValid: true,
		Message: "",
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
