/**
 * @api           {post} /v1/account/profile/accept_tos Accept Terms of Service
 * @apiVersion    1.2.2
 * @apiName       AcceptTOS
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Accept Terms of Service.
 *
 * @apiParam {long}    createdTime      The account's created time, to validate the request.
 *
 * @apiParamExample {json} Request Example:
 *     {
 *       "createdTime": 1566452967522
 *     }
 *
 * @apiSuccess {boolean}  success          If accepted successfully.
 * @apiSuccess {string}   message          The message.
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
 *         "message": "Terms of Service accepted",
 *       }
 *     }
 *
 * @apiUse   ErrorAccountHeaderValidationFailed
 * @apiError ParamValidationFailed The parameter validation failed.
 *
 * @apiErrorExample {json} ParamValidationFailed:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40002",
 *         "message": "Created time is invalid",
 *         "field": "createdTime"
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
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// AccountProfileAcceptTOSRequestParam represents request data of Account API "Accept Terms of Service".
type AccountProfileAcceptTOSRequestParam struct {
	CreatedTime int64 `json:"createdTime"`
}

// AccountProfileAcceptTOSResponseData represents response data of Account API "Accept Terms of Service".
type AccountProfileAcceptTOSResponseData struct {
	api.ResponseData
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AccountProfileAcceptTOS accepts Terms of Service for a customer account.
func AccountProfileAcceptTOS(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.AccountProfileAcceptTOS")
	account := ctx.Account

	// Get the Redis connection and defer closing connection.
	redisConn := redis.GetConnection()
	defer redisConn.Close()

	var param AccountProfileAcceptTOSRequestParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		logger.Error(ctx.ReqTag, err.Error())
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.ReqParamValidationFailed, msg)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	var msg, field string
	if param.CreatedTime == 0 {
		msg = "Created time is required"
		field = "createdTime"
	} else if param.CreatedTime != account.CreatedTime {
		msg = "Created time is invalid"
		field = "createdTime"
	}
	if msg != "" {
		response := api.NewAPIResponseWithErrorField(ctx.ReqID, errcode.ReqParamValidationFailed, msg, field)
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	success := false
	message := ""
	// Get the account's terms of service details.
	tosRepo := repository.NewCstAccountTOSRepo(redisConn)
	_, err := tosRepo.GetByAccountID(ctx.Account.ID)
	if err != nil {
		if err == tosRepo.ErrNotFound {
			// Begin database transaction.
			tx, err := db.Get().Begin()
			if err != nil {
				logger.Fatal("db.Begin", logger.FromError(err))
				response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
				api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
				return
			}
			defer tx.Rollback()

			// Insert the Terms of Service acceptance to database.
			accountTOS, err := dao.NewCstAccountTOSDAO().Insert(tx, account.ID)
			if err != nil {
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

			// Save to Redis.
			tosRepo.RedisStore().Save(accountTOS)

			success = true
			message = "Terms of Service is accepted."
		} else {
			response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, err.Error())
			api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
			return
		}
	} else {
		success = false
		message = "Terms of Service is already accepted"
	}

	// Return the result.
	data := AccountProfileAcceptTOSResponseData{
		Success: success,
		Message: message,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
