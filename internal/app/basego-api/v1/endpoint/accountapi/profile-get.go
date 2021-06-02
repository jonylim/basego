/**
 * @apiDefine SuccessAccountProfile
 *
 * @apiSuccess {object}  account                       The account's profile.
 * @apiSuccess {long}    account.id                    The account ID.
 * @apiSuccess {string}  account.fullName              The account's full name.
 * @apiSuccess {string}  account.email                 The email address.
 * @apiSuccess {boolean} account.isEmailVerified       If the email address is verified.
 * @apiSuccess {integer} account.countryID             The account's country ID.
 * @apiSuccess {string}  account.countryCallingCode    The country calling code for phone number.
 * @apiSuccess {string}  account.phone                 The phone number.
 * @apiSuccess {string}  account.phoneWithCode         The phone number with country calling code.
 * @apiSuccess {boolean} account.isPhoneVerified       If the phone number is verified.
 * @apiSuccess {object}  account.imageURL              The account's picture image URL.
 * @apiSuccess {string}  account.imageURL.thumbnail    Image URL for thumbnail picture.
 * @apiSuccess {string}  account.imageURL.fullsize     Image URL for fullsize picture.
 * @apiSuccess {long}    account.lastLoginTime         The account's last login, in Unix milliseconds.
 * @apiSuccess {long}    account.lastActivityTime      The account's last activity, in Unix milliseconds.
 * @apiSuccess {boolean} account.requireChangePassword If the account is required to change password.
 * @apiSuccess {long}    account.createdTime           The time the account was created, in Unix milliseconds.
 * @apiSuccess {long}    account.updatedTime           The time the account was last updated, in Unix milliseconds.
 * @apiSuccess {long}    account.deletedTime           The time the account was deleted, in Unix milliseconds.
 */

/**
 * @api           {post} /v1/account/profile/get Get Account Profile
 * @apiVersion    1.0.0
 * @apiName       GetAccountProfile
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Get the current account's profile.
 *
 * @apiUse     SuccessAccountProfile
 * @apiSuccess {object}   account                       The account's profile.
 * @apiSuccess {object}   tos                           The account's Terms of Service status.
 * @apiSuccess {boolean}  tos.isAccepted                If the Terms of Service has been accepted.
 * @apiSuccess {long}     tos.acceptedTime              The time the Terms of Service was accepted, in Unix milliseconds.
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
 *         "account": {
 *           "id": 8,
 *           "fullName": "Jony",
 *           "email": "",
 *           "isEmailVerified": false,
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
 *         },
 *         "tos": {
 *           "isAccepted": true,
 *           "acceptedTime": 1566452967572
 *         }
 *       }
 *     }
 *
 * @apiUse   ErrorAccountHeaderValidationFailed
 */

package accountapi

import (
	"net/http"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/repository"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// AccountProfileGetResponseData represents response data of Account API "Get Account Profile".
type AccountProfileGetResponseData struct {
	api.ResponseData
	Account model.CstAccount     `json:"account"`
	TOS     accountProfileGetTOS `json:"tos"`
}

type accountProfileGetTOS struct {
	IsAccepted   bool  `json:"isAccepted"`
	AcceptedTime int64 `json:"acceptedTime"`
}

// AccountProfileGet returns the logged in account's profile.
func AccountProfileGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.AccountProfileGet")

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	// Get the account's TOS status.
	var tosStatus accountProfileGetTOS
	tosRepo := repository.NewCstAccountTOSRepo(redisConn)
	tos, err := tosRepo.GetByAccountID(ctx.Account.ID)
	if err == nil {
		tosStatus.IsAccepted = true
		tosStatus.AcceptedTime = tos.CreatedTime
	} else if err != tosRepo.ErrNotFound {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Return the result.
	data := AccountProfileGetResponseData{
		Account: ctx.Account,
		TOS:     tosStatus,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
