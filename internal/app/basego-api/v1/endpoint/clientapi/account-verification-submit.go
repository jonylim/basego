/**
 * @api           {post} /v1/client/account_verification/submit Account Verification - Submit
 * @apiVersion    1.0.0
 * @apiName       AccountVerification_Submit
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Submit code for account verification.
 *
 * @apiParam {long}   otpID   The OTP ID.
 * @apiParam {string} otpKey  The OTP key.
 * @apiParam {string} otpCode The OTP code.
 * @apiParam {string} email   The email address of the account to be verified.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "otpID": 128,
 *       "otpKey": "49ff390684515734c4c645df3884ed7c",
 *       "otpCode": "123456",
 *       "email": "john@doe.com"
 *     }
 *
 * @apiSuccess {boolean} success If the account verification is successful.
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
 *         "success": true,
 *         "message": "Your account verification is successful"
 *       }
 *     }
 *
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError ParamValidationFailed The parameter validation failed.
 * @apiError EmailNotRegistered    The email address is not registered.
 *
 * @apiErrorExample {json} ParamValidationFailed:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Email address format is invalid",
 *         "field": "email"
 *       },
 *       "data": {}
 *     }
 *
 * @apiErrorExample {json} EmailNotRegistered:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "The email address is not registered",
 *         "field": "email"
 *       },
 *       "data": {}
 *     }
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

// AccountVerificationSubmitRequestParam represents request body of Client API "Account Verification - Submit".
type AccountVerificationSubmitRequestParam struct {
	Email   string `json:"email"`
	OTPID   int64  `json:"otpID"`
	OTPKey  string `json:"otpKey"`
	OTPCode string `json:"otpCode"`
}

// AccountVerificationSubmitResponseData represents response data of Client API "Account Verification - Submit".
type AccountVerificationSubmitResponseData struct {
	api.ResponseData
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AccountVerificationSubmit verifies the OTP and email for account verification.
func AccountVerificationSubmit(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.AccountVerificationSubmit")

	var param AccountVerificationSubmitRequestParam
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
		msg = "Verification ID is required"
		field = "otpID"
	} else if param.OTPKey == "" {
		msg = "Verification key is required"
		field = "otpKey"
	} else if param.OTPCode == "" {
		msg = "Verification code is required"
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
			msg = "The email address is not registered"
			field = "email"
			response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
			return
		}
		if err == accRepo.ErrDatabase {
			err = errDatabase
		} else {
			err = errInternal
		}
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Check if the email address has been verified.
	if account.IsEmailVerified {
		data := AccountVerificationSubmitResponseData{
			Success: false,
			Message: "The account has already been verified",
		}
		response := api.NewAPIResponse(ctx.ReqID)
		response.SetData(data)
		api.SendResponseJSON(w, response)
		return
	}

	// Get active OTP's details.
	otpRepo := repository.NewCstAccountOTPRepo(redisConn)
	otpData, err := otpRepo.GetActiveOTPByAccountAndAction(account.ID, otp.ActionVerifyEmail)
	if err != nil {
		if err == otpRepo.ErrNotFound {
			msg = "There is no pending verification found, or the code has expired"
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
		msg = "Verification ID is invalid"
		field = "otpID"
		isValidID = false
	} else if otpData.Key != param.OTPKey {
		msg = "Verification key is invalid"
		field = "otpKey"
	} else if otpData.Code != param.OTPCode {
		msg = "Verification code is incorrect"
		field = "otpCode"
	} else if otpData.Email != account.Email {
		msg = "Email address has changed, you should request a new verification code"
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

		// Increment the OTP's verification attempt count.
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
				otpRepo.RedisStore().SaveOTPByAccountAndAction(otpData, emailVerificationTTL)
			}
		}
	}
	if msg != "" {
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	// The verification is successful.
	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Mark the OTP as verified.
	attemptCount, isVerified, err := dao.NewCstAccountOTPDAO().SetVerified(tx, otpData.ID)
	if attemptCount == 0 || err != nil {
		otpRepo.RedisStore().DeleteOTPByID(otpData.ID)
		otpRepo.RedisStore().DeleteOTPByAccountAndAction(otpData.AccountID, otpData.Action)

		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	otpData.AttemptCount = attemptCount
	otpData.IsVerified = isVerified
	otpData.UpdatedTime = helper.UnixMillisecond(time.Now())

	// Mark the account's email address as verified.
	_, err = dao.NewCstAccountDAO().SetVerifiedEmail(tx, account.ID)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	account.IsEmailVerified = true

	// Commit database transaction.
	err = tx.Commit()
	if err != nil {
		logger.Fatal("tx.Commit", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Update to Redis.
	if err := accRepo.RedisStore().Save(account); err != nil {
		accRepo.RedisStore().Delete(account)
	}
	otpRepo.RedisStore().SaveOTPByAccountAndAction(otpData, emailVerificationTTL)

	// Return the response.
	data := AccountVerificationSubmitResponseData{
		Success: true,
		Message: "Your account verification is successful",
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
