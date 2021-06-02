/**
 * @api           {post} /v1/client/reset_password/request_token Reset Password - Request Token
 * @apiVersion    1.0.0
 * @apiName       ResetPassword_RequestToken
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Request token for reset password.
 *
 * @apiParam {string} email The account's email address.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "email": "john@doe.com",
 *     }
 *
 * @apiSuccess {boolean} success       If email containing the reset password token is sent successfully.
 * @apiSuccess {string}  message       The message.
 * @apiSuccess {long}    otpID         The OTP ID for reset paassword token.
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
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError ParamValidationFailed  The parameter validation failed.
 *
 * @apiErrorExample {json} ParamValidationFailed:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Email address is invalid",
 *         "field": "email"
 *       },
 *       "data": {}
 *     }
 *
 */

package clientapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/redisstore"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/emailtemplate"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/token/otp"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"
	"github.com/jonylim/basego/internal/pkg/common/send/email"

	"github.com/julienschmidt/httprouter"
)

// ResetPasswordRequestTokenRequestParam represents request body of Client API "Reset Password - Request Token".
type ResetPasswordRequestTokenRequestParam struct {
	Email string `json:"email"`
}

// ResetPasswordRequestTokenResponseData represents response data of Client API "Reset Password - Request Token".
type ResetPasswordRequestTokenResponseData struct {
	api.ResponseData
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	OTPID      int64  `json:"otpID"`
	OTPKey     string `json:"otpKey"`
	CodeLength int32  `json:"codeLength"`
}

const emailResetPasswordTTL = 3600

// ResetPasswordRequestToken sends email containing request token for reset password.
func ResetPasswordRequestToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.ResetPasswordRequestToken")

	var param ResetPasswordRequestTokenRequestParam
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
	param.Email = strings.ToLower(param.Email)

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Check if the email address has not been registered.
	accRepo := repository.NewCstAccountRepo(redisConn)
	account, err := accRepo.GetByEmail(param.Email)
	if err != nil {
		if err == accRepo.ErrNotFound {
			msg = "The email address is invalid"
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

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Generate OTP for reset password.
	otpKey, otpCode := otp.GenerateAlphanumeric()
	expiryTime := time.Now().Add(emailResetPasswordTTL * time.Second)
	otpData := model.CstAccountOTP{
		AccountID:  account.ID,
		Key:        otpKey,
		Code:       otpCode,
		Action:     otp.ActionResetPassword,
		Method:     otp.MethodEmail,
		Email:      param.Email,
		ExpiryTime: helper.UnixMillisecond(expiryTime),
		SendCount:  1,
	}

	// Insert the OTP to database.
	otpID, otpCreatedMillis, err := dao.NewCstAccountOTPDAO().InsertOTP(tx, otpData)
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
	redisstore.NewCstAccountOTPStore(redisConn).SaveOTPByAccountAndAction(otpData, emailResetPasswordTTL*2)

	// Send reset password email.
	go sendResetPasswordEmail(account, otpData)

	// Return the response.
	data := ResetPasswordRequestTokenResponseData{
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

func sendResetPasswordEmail(account model.CstAccount, otpData model.CstAccountOTP) {
	tokenData := emailResetPasswordToken{otpData.ID, otpData.Key, otpData.Code, account.Email}
	tokenString, err := tokenData.Encode()
	if err != nil {
		logger.Fatal("api", fmt.Sprintf("sendResetPasswordEmail: %v", err))
		return
	}
	q := url.Values{"token": []string{tokenString}}
	link := os.Getenv(envvar.FrontendURL) + "/reset-password?" + q.Encode()

	/* data := struct{ Name, Code, Link, TTLHours string }{
			Name:     account.FullName,
			Link:     os.Getenv(envvar.FrontendURL) + "/reset-password?" + q.Encode(),
			Code:     otpData.Code,
			TTLHours: helper.IntToString(emailResetPasswordTTL / 3600),
		}

		subject := "Reset password request"
		body := `<!DOCTYPE html
	  PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
			<html>

			<head>
				<title>{{.Title}}</title>
				<meta name="viewport" content="width=device-width, initial-scale=1" />
			</head>

			<body>
				<div id="wrapper" style="text-align: center">
					<div id="content"
						style="background-color: #ffffff; border-radius: 8px; border: solid 1px #dcdcdc; font-size: 14px; padding: 32px 51px 24px; text-align: center; max-width: 483px; display: inline-block; box-sizing: border-box; font-family: Arial,Helvetica,sans-serif">
						<img alt="Logo" src="https://placeholder.com/wp-content/uploads/2018/10/placeholder.com-logo1.png" style="width: 120px; display: inline-block; margin-bottom: 32px" />
						<div style="letter-spacing: -0.4px; color: #191919; font-size: 20px; font-weight: bold">Reset Password
						</div>
						<div style="margin-top: 20px; color: #191919">Hi, {{.Name}}!</div>
						<div style="margin-top: 10px; letter-spacing: -0.2px; color: #191919">
							To reset your password, please click the following button.</div>
						<a style="background-image: linear-gradient(to bottom, #ff9833, #ff7e00 100%); border: solid 1px #ff7e00; border-radius: 8px; box-shadow: 0 6px 6px 0 rgba(255, 126, 0, 0.2), 0 0 6px 0 rgba(255, 126, 0, 0.1); color: #ffffff; cursor: pointer; display: inline-block; font-size: 16px; margin-top: 20px; padding: 15px 0; text-align: center; text-decoration: none; width: 219px"
							href="{{.Link}}" target="_blank">Reset Password Now</a>
						<div style="margin-top: 24px; letter-spacing: -0.2px; color: #191919">The link will only be valid for
							{{.TTLHours}} hours.</div>
						<div
							style="border-top: solid 1px #dcdcdc; color: #9b9b9b; font-size: 12px; letter-spacing: -0.2px; line-height: 1.43; margin-top: 32px; padding-top: 24px; text-align: center">
							This email was sent to you because you have registered an account using this email address.
							If you did not register, please ignore this email.
						</div>
					</div>
				</div>
			</body>

			</html>`
		t, err := template.New("emailResetPasswordTemplate").Parse(body)
		if err != nil {
			logger.Fatal("api", fmt.Sprintf("sendResetPasswordEmail: %v", err))
			return
		}

		var buf bytes.Buffer
		err = t.Execute(&buf, data)
		if err != nil {
			logger.Fatal("api", fmt.Sprintf("sendResetPasswordEmail: %v", err))
			return
		} */

	subject, body, err := emailtemplate.ResetPassword(account.FullName, link, otpData.Code, emailResetPasswordTTL/3600)
	if err == nil {
		email.Send(email.NewHTMLMessage(subject, body), email.Recipients{
			To: []string{account.Email},
		})
	}
}
