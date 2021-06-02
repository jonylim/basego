/**
 * @api           {post} /v1/client/register Register
 * @apiVersion    1.0.0
 * @apiName       Register
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Register a new customer account.
 *
 * @apiParam {string} fullName The full name.
 * @apiParam {string} email    The email address.
 * @apiParam {string} password The password.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "fullName": "John Doe",
 *       "email": "john@doe.com",
 *       "password": "this_is_password"
 *     }
 *
 * @apiSuccess {boolean} success       If the registration is successful.
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
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError ParamValidationFailed  The parameter validation failed.
 * @apiError EmailAlreadyRegistered The email address is already registered.
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
 * @apiErrorExample {json} EmailAlreadyRegistered:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "The email address is already registered",
 *         "field": "email"
 *       },
 *       "data": {}
 *     }
 */

/**
 * @api           {post} /v1/client/register Register
 * @apiVersion    1.2.2
 * @apiName       Register
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Register a new customer account.
 *
 * @apiParam {string}  fullName        The full name.
 * @apiParam {string}  email           The email address.
 * @apiParam {string}  password        The password.
 * @apiParam {string}  [isTOSAccepted] If the Terms of Service is accepted.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "fullName": "John Doe",
 *       "email": "john@doe.com",
 *       "password": "this_is_password",
 *       "isTOSAccepted": true
 *     }
 *
 * @apiSuccess {boolean} success       If the registration is successful.
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
 * @apiUse   ErrorClientHeaderValidationFailed
 * @apiError ParamValidationFailed  The parameter validation failed.
 * @apiError EmailAlreadyRegistered The email address is already registered.
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
 * @apiErrorExample {json} EmailAlreadyRegistered:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "The email address is already registered",
 *         "field": "email"
 *       },
 *       "data": {}
 *     }
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
	"github.com/jonylim/basego/internal/pkg/common/crypto/password"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"
	"github.com/jonylim/basego/internal/pkg/common/send/email"

	"github.com/julienschmidt/httprouter"
)

// RegisterRequestParam represents request body of Client API "Register".
type RegisterRequestParam struct {
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	IsTOSAccepted bool   `json:"isTOSAccepted"`
}

// RegisterResponseData represents response data of Client API "Register".
type RegisterResponseData struct {
	api.ResponseData
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	OTPID      int64  `json:"otpID"`
	OTPKey     string `json:"otpKey"`
	CodeLength int32  `json:"codeLength"`
}

const emailVerificationTTL = 86400

// Register registers a new customer account.
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.Register")

	var param RegisterRequestParam
	errReq := json.NewDecoder(r.Body).Decode(&param)
	if errReq != nil {
		logger.Error(ctx.ReqTag, errReq.Error())
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	var msg, field string
	if param.FullName == "" {
		msg = "Full name is required"
		field = "fullName"
	} else if param.Email == "" {
		msg = "Email address is required"
		field = "email"
	} else if err := helper.ValidateEmailFormat(param.Email); err != nil {
		msg = err.Error()
		field = "email"
	} else if param.Password == "" {
		msg = "Password is required"
		field = "password"
	} else if err := helper.ValidatePasswordFormat(param.Password); err != nil {
		msg = err.Error()
		field = "password"
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

	// Check if the email address has been registered.
	accRepo := repository.NewCstAccountRepo(redisConn)
	exists, err := accRepo.ExistsByEmail(param.Email)
	if exists {
		msg = "The email address is already registered"
		field = "email"
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	} else if err != nil {
		if err == accRepo.ErrDatabase {
			err = errDatabase
		} else {
			err = errInternal
		}
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Generate password salt and hash the password.
	pwdSalt := password.GenerateSalt()
	account := model.CstAccount{
		FullName:     param.FullName,
		Email:        param.Email,
		Password:     password.HashWithSalt(param.Password, pwdSalt),
		PasswordSalt: pwdSalt,
	}
	// if env != "production" {
	// 	account.IsEmailVerified = true
	// }

	// Begin database transaction.
	tx, err := db.Get().Begin()
	if err != nil {
		logger.Fatal("db.Begin", logger.FromError(err))
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}
	defer tx.Rollback()

	// Insert the new account to database.
	account.ID, account.CreatedTime, err = dao.NewCstAccountDAO().Insert(tx, account)
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	var accountTOS model.CstAccountTOS
	if param.IsTOSAccepted {
		// Insert the account's Terms of Service acceptance to database.
		accountTOS, err = dao.NewCstAccountTOSDAO().Insert(tx, account.ID)
		if err != nil {
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
			return
		}
	}

	// Generate OTP for email verification.
	otpKey, otpCode := otp.GenerateAlphanumeric()
	expiryTime := time.Now().Add(emailVerificationTTL * time.Second)
	otpData := model.CstAccountOTP{
		AccountID:  account.ID,
		Key:        otpKey,
		Code:       otpCode,
		Action:     otp.ActionVerifyEmail,
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
	accRepo.RedisStore().Save(account)
	redisstore.NewCstAccountOTPStore(redisConn).SaveOTPByAccountAndAction(otpData, emailVerificationTTL*2)
	accountTOSStore := redisstore.NewCstAccountTOSStore(redisConn)
	if param.IsTOSAccepted && accountTOS.ID != 0 {
		accountTOSStore.Save(accountTOS)
	} else {
		accountTOSStore.SaveNilByAccountID(account.ID)
	}

	if !account.IsEmailVerified {
		// Send verification email.
		go sendVerificationEmail(account, otpData)
	}

	// Return the response.
	data := RegisterResponseData{
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

func sendVerificationEmail(account model.CstAccount, otpData model.CstAccountOTP) {
	tokenData := emailVerificationToken{otpData.ID, otpData.Key, otpData.Code, account.Email}
	tokenString, err := tokenData.Encode()
	if err != nil {
		logger.Fatal("api", fmt.Sprintf("sendVerificationEmail: %v", err))
		return
	}
	q := url.Values{"token": []string{tokenString}}
	link := os.Getenv(envvar.FrontendURL) + "/verify/email?" + q.Encode()

	/* data := struct{ Name, Code, Link, TTLHours string }{
		Name:     account.FullName,
		Link:     os.Getenv(envvar.FrontendURL) + "/verify/email?" + q.Encode(),
		Code:     otpData.Code,
		TTLHours: helper.IntToString(emailVerificationTTL / 3600),
	}

	subject := "Please verify your email address"
	body := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html>
			<head>
				<meta name="viewport" content="width=device-width, initial-scale=1" />
			</head>
			<body>
				<div id="wrapper" style="text-align: center">
					<div id="content" style="background-color: #ffffff; border-radius: 8px; border: solid 1px #dcdcdc; font-size: 14px; padding: 32px 51px 24px; text-align: center; max-width: 483px; display: inline-block; box-sizing: border-box; font-family: Arial,Helvetica,sans-serif">
						<img alt="Logo" src="https://placeholder.com/wp-content/uploads/2018/10/placeholder.com-logo1.png" style="width: 120px; display: inline-block; margin-bottom: 32px" />
						<div style="letter-spacing: -0.4px; color: #191919; font-size: 20px; font-weight: bold">Verify your email address</div>
						<div style="margin-top: 20px; color: #191919">Hi, {{.Name}}!</div>
						<div style="margin-top: 10px; letter-spacing: -0.2px; color: #191919">To verify email address, please click the following button.</div>
						<a style="background-image: linear-gradient(to bottom, #ff9833, #ff7e00 100%); border: solid 1px #ff7e00; border-radius: 8px; box-shadow: 0 6px 6px 0 rgba(255, 126, 0, 0.2), 0 0 6px 0 rgba(255, 126, 0, 0.1); color: #ffffff; cursor: pointer; display: inline-block; font-size: 16px; margin-top: 20px; padding: 15px 0; text-align: center; text-decoration: none; width: 219px"
							href="{{.Link}}" target="_blank">
							Verify Now
						</a>
						<div style="margin-top: 24px; letter-spacing: -0.2px; color: #191919">The link will only be valid for {{.TTLHours}} hours.</div>
						<div style="border-top: solid 1px #dcdcdc; color: #9b9b9b; font-size: 12px; letter-spacing: -0.2px; line-height: 1.43; margin-top: 32px; padding-top: 24px; text-align: center">
							This email was sent to you because you have registered an account using this email address.
							If you did not register, please ignore this email.
						</div>
					</div>
				</div>
			</body>
		</html>`
	t, err := template.New("emailVerificationTemplate").Parse(body)
	if err != nil {
		logger.Fatal("api", fmt.Sprintf("sendVerificationEmail: %v", err))
		return
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		logger.Fatal("api", fmt.Sprintf("sendVerificationEmail: %v", err))
		return
	} */

	subject, body, err := emailtemplate.VerifyEmailAddress(account.FullName, link, otpData.Code, emailVerificationTTL/3600)
	if err == nil {
		email.Send(email.NewHTMLMessage(subject, body), email.Recipients{
			To: []string{account.Email},
		})
	}
}
