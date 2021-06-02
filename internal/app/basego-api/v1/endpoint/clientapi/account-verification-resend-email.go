/**
 * @api           {post} /v1/client/account_verification/resend_email Account Verification - Resend Email
 * @apiVersion    1.0.0
 * @apiName       AccountVerification_ResendEmail
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Resend email for account verification.
 *
 * @apiParam {string} email The account's email address.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "email": "john@doe.com"
 *     }
 *
 * @apiSuccess {boolean} success       If the email is resent successfully.
 * @apiSuccess {string}  message       The message, if failed.
 * @apiSuccess {long}    otpID         The OTP ID for OTP verification.
 * @apiSuccess {string}  otpKey        The OTP key.
 * @apiSuccess {integer} codeLength    The OTP code's length.
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
 *         "message": "",
 *         "otpID": 128,
 *         "otpKey": "49ff390684515734c4c645df3884ed7c",
 *         "codeLength": 6
 *       }
 *     }
 *
 * @apiSuccessExample {json} Success Response (Already Verified):
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "error": {
 *         "code": "",
 *         "message": "",
 *         "field": ""
 *       },
 *       "data": {
 *         "success": false,
 *         "message": "The account has already been verified"
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
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
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

// AccountVerificationResendEmailRequestParam represents request body of Client API "Account Verification - Resend Email".
type AccountVerificationResendEmailRequestParam struct {
	Email string `json:"email"`
}

// AccountVerificationResendEmailResponseData represents response data of Client API "Account Verification - Resend Email".
type AccountVerificationResendEmailResponseData struct {
	api.ResponseData
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	OTPID      int64  `json:"otpID,omitempty"`
	OTPKey     string `json:"otpKey,omitempty"`
	CodeLength int32  `json:"codeLength,omitempty"`
}

// AccountVerificationResendEmail resends account verification email to the specified email address.
func AccountVerificationResendEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.AccountVerificationResendEmail")

	var param AccountVerificationResendEmailRequestParam
	errReq := json.NewDecoder(r.Body).Decode(&param)
	if errReq != nil {
		logger.Error(ctx.ReqTag, errReq.Error())
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	var msg, field string
	if param.Email == "" {
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
		data := AccountVerificationResendEmailResponseData{
			Success: false,
			Message: "The account has already been verified",
		}
		response := api.NewAPIResponse(ctx.ReqID)
		response.SetData(data)
		api.SendResponseJSON(w, response)
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

	// Generate OTP for email verification.
	otpKey, otpCode := otp.GenerateAlphanumeric()
	expiryTime := time.Now().Add(emailVerificationTTL * time.Second)
	otpData := model.CstAccountOTP{
		AccountID:  account.ID,
		Key:        otpKey,
		Code:       otpCode,
		Action:     otp.ActionVerifyEmail,
		Method:     otp.MethodEmail,
		Email:      account.Email,
		ExpiryTime: helper.UnixMillisecond(expiryTime),
		SendCount:  1,
	}

	// Delete currently active OTP by account and action.
	otpDAO := dao.NewCstAccountOTPDAO()
	deletedID, lastSendCount, err := otpDAO.DeleteActiveOTPByAccountAndAction(tx, otpData.AccountID, otpData.Action)
	if err != nil {
		// NOTE: Error deleting active OTP can be ignored.
		// response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		// api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		// return
	}

	// Set next send count.
	otpData.SendCount = lastSendCount + 1

	// Insert the new OTP to database.
	otpID, otpCreatedMillis, err := otpDAO.InsertOTP(tx, otpData)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	otpData.ID, otpData.CreatedTime = otpID, otpCreatedMillis

	// Commit database transaction.
	err = tx.Commit()
	if err != nil {
		logger.Fatal("tx.Commit", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Save to Redis.
	otpStore := redisstore.NewCstAccountOTPStore(redisConn)
	otpStore.DeleteOTPByAccountAndAction(otpData.AccountID, otpData.Action)
	otpStore.SaveNilByID(deletedID, emailVerificationTTL)
	otpStore.SaveOTPByAccountAndAction(otpData, emailVerificationTTL*2)

	// Send verification email.
	go sendVerificationEmail(account, otpData)

	// Return the response.
	data := AccountVerificationResendEmailResponseData{
		Success:    true,
		Message:    "",
		OTPID:      otpID,
		OTPKey:     otpKey,
		CodeLength: otp.Length,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
