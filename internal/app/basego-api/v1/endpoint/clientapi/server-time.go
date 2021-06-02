/**
 * @api           {post} /v1/client/server_time Get Server Time
 * @apiVersion    1.0.0
 * @apiName       ServerTime
 * @apiGroup      ClientAPI
 * @apiPermission client
 *
 * @apiDescription Get the server's current timestamp.
 *
 * @apiSuccess {object} timestamp              The server's current timestamp.
 * @apiSuccess {long}   timestamp.seconds      The timestamp in seconds.
 * @apiSuccess {long}   timestamp.milliseconds The timestamp in milliseconds.
 * @apiSuccess {long}   timestamp.nanoseconds  The timestamp in nanoseconds.
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
 *         "timestamp": {
 *           "seconds": 1551841641,
 *           "milliseconds": 1551841641095,
 *           "nanoseconds": 1551841641095244400
 *         }
 *       }
 *     }
 *
 * @apiUse   ErrorClientHeaderValidationFailed
 */

package clientapi

import (
	"net/http"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// ServerTimeResponseData represents response data of Client API "Get Server Time".
type ServerTimeResponseData struct {
	api.ResponseData
	Timestamp serverTimeResponseDataTimestamp `json:"timestamp"`
}
type serverTimeResponseDataTimestamp struct {
	Seconds      int64 `json:"seconds"`
	Milliseconds int64 `json:"milliseconds"`
	Nanoseconds  int64 `json:"nanoseconds"`
}

// ServerTime returns the server's current timestamp.
func ServerTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: clientapi.ServerTime")

	// Get current time.
	now := time.Now()

	// Return the response.
	data := ServerTimeResponseData{
		Timestamp: serverTimeResponseDataTimestamp{
			Seconds:      now.Unix(),
			Milliseconds: helper.UnixMillisecond(now),
			Nanoseconds:  now.UnixNano(),
		},
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
