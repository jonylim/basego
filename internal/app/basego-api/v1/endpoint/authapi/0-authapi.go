/**
 * @apiDefine AuthAPI Authentication API
 *
 * #### HTTP Request Headers
 * | **Header Name**   | **Required** | **Description** |
 * |-------------------|:------------:|-----------------|
 * | API-Key           | ✓ | API key for accessing the API. |
 * | App-Identifier    |   | The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web). |
 * | Authorization     | ✓ | Authorization type and credentials, e.g.: basic credentials or refresh token to request new access token.<br>Format: <code><i>&lt;type&gt; &lt;credentials&gt;</i></code> |
 * | Content-Type      |   | Content type of the request body. |
 * | Device-Identifier | ✓ | The device ID (optional for web). |
 * | Device-Model      | ✓ | Model name of the device (optional for web). |
 * | Device-Platform   | ✓ | The device's platform. Values are `android`, `ios`, or `web`. |
 * | User-Agent        |   | The user agent of the client accessing the API. |
 *
 * #### HTTP Response Status Codes
 * | **Code** | **Description**                                                                                        |
 * |:--------:|--------------------------------------------------------------------------------------------------------|
 * |   200    | OK, request proceed without error.                                                                     |
 * |   204    | Request proceed successfully and is not returning any content (empty `data`).                          |
 * |   400    | Bad request, either the request header or the API parameter validation failed.                         |
 * |   401    | Unauthorized access, either the HTTP request header `Authorization` is missing or invalid.             |
 * |   403    | Forbidden, the user does not have access to the requested resource.                                    |
 * |   404    | Not found, requested resource does not exist.                                                          |
 * |   429    | Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.         |
 * |   491    | The API key specified in HTTP request header `API-Key` is invalid.                                     |
 * |   500    | Internal server error while processing the request.                                                    |
 *
 * #### API Error Codes
 * | **Code** | **Description**                                                                                        |
 * |:--------:|--------------------------------------------------------------------------------------------------------|
 * |  40001   | HTTP request header validation failed.                                                                 |
 * |  40002   | API request parameter validation failed.                                                               |
 * |  40101   | HTTP request header `Authorization` is not provided when required.                                     |
 * |  40102   | HTTP request header `Authorization` format is invalid.                                                 |
 * |  40103   | Failed to parse the token or the token is not found or invalid. Client should get a new token.         |
 * |  40104   | The token has expired. Client should get a new token.                                                  |
 * |  40105   | The token owner does not belong to the client's info. Client should get a new token.                   |
 * |  40106   | User is not found for the specified token.                                                             |
 * |  40107   | The user's account has not been verified.                                                              |
 * |  40301   | The user does not have access to the requested resource or action.                                     |
 * |  40401   | The requested resource is not found.                                                                   |
 * |  49101   | The API key is not provided.                                                                           |
 * |  49102   | Failed to parse the API key, or the API key is invalid.                                                |
 * |  49103   | The provided API key is not found.                                                                     |
 * |  49104   | The provided API key is not intended to be used with the client's app platform.                        |
 * |  49105   | The provided API key is not intended to be used with the client's app identifier.                      |
 * |  49106   | The API key has expired.                                                                               |
 * |  49107   | The API key is disabled.                                                                               |
 * |  50001   | An error occurred while validating the API key.                                                        |
 * |  99999   | Other errors, usually without specific reason or action.                                               |
 */

/**
 * @apiDefine ErrorAuthHeaderValidationFailed
 * @apiVersion 1.0.0
 *
 * @apiError HeaderValidationFailed The request header validation failed.
 * @apiErrorExample {json} HeaderValidationFailed:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 400,
 *       "error": {
 *         "code": "40001",
 *         "message": "Request headers are required (API-Key, Authorization, Device-Platform)",
 *         "field": ""
 *       },
 *       "data": {}
 *     }
 */

package authapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/jonylim/basego/internal/app/basego-api/v1/endpoint/authapi/requestheader"
	"github.com/jonylim/basego/internal/app/basego-api/v1/requestvalidator"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

type (
	// Context contains a request's context.
	Context struct {
		context.Context
		ReqID     string
		ReqTag    string
		Path      string
		ReqHeader requestheader.APIRequestHeader
		APIKey    model.XAPIKey
	}

	// Handle handles requests for auth APIs.
	Handle func(http.ResponseWriter, *http.Request, httprouter.Params, Context)
)

var (
	errDatabase = errors.New("An error occurred while processing your request")
	errInternal = errors.New("An error occurred while processing your request")
)

var env string

// Init initializes required variables.
func Init() {
	env = os.Getenv(envvar.Environment)
}

// HandleRequest handles a request for auth APIs.
func HandleRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params, handle Handle) {
	reqID := api.CreateReqID()

	// Validate request headers.
	reqHeader := requestheader.Parse(r)
	if err := requestheader.CheckRequired(reqHeader); err != nil {
		response := api.NewAPIResponseWithError(reqID, errcode.ReqHeaderValidationFailed, err.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return
	}

	// Validate request's API key.
	validator := requestvalidator.New(w, r, reqID)
	apiKey, ok := validator.ValidateAPIKey(reqHeader.APIKey, reqHeader.AppIdentifier, reqHeader.DevicePlatform)
	if !ok {
		return
	}

	// OK!
	reqTag := fmt.Sprintf("api:%s", reqID)
	path := r.URL.Path
	logger.Trace(reqTag, "Path: "+path)
	handle(w, r, p, Context{
		Context:   r.Context(),
		ReqID:     reqID,
		ReqTag:    reqTag,
		Path:      path,
		ReqHeader: reqHeader,
		APIKey:    apiKey,
	})
}
