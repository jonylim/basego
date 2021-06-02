/**
 * @api           {post} /v1/account/time_zones Get Time Zone List
 * @apiVersion    1.0.0
 * @apiName       TimeZones
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Get the list of time zones.
 *
 * @apiSuccess {object}  timeZones           The list of time zones.
 * @apiSuccess {string}  timeZones.name      The time zone's name.
 * @apiSuccess {string}  timeZones.abbrev    The time zone's abbrev.
 * @apiSuccess {string}  timeZones.utcOffset The offset from UTC (format: `"+HH:mm:ss"`).
 * @apiSuccess {boolean} timeZones.isDST     If the time zone is currently in DST.
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
 *         "timeZones": [
 *           {
 *             "name": "Asia/Jakarta",
 *             "abbrev": "WIB",
 *             "utcOffset": "07:00:00",
 *             "isDST": false
 *           },
 *           {
 *             "name": "America/New_York",
 *             "abbrev": "EDT",
 *             "utcOffset": "-04:00:00",
 *             "isDST": true
 *           }
 *         ]
 *       }
 *     }
 *
 * @apiUse   ErrorAccountHeaderValidationFailed
 */

package accountapi

import (
	"net/http"

	"github.com/jonylim/basego/internal/pkg/basego-api/v1/data/dao"
	"github.com/jonylim/basego/internal/pkg/basego-api/v1/model"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	"github.com/julienschmidt/httprouter"
)

// TimeZonesResponseData represents response data of Account API "Get Time Zones".
type TimeZonesResponseData struct {
	api.ResponseData
	TimeZones []model.PgTimeZone `json:"timeZones"`
}

// TimeZones returns the list of time zones.
func TimeZones(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.TimeZones")

	// Get timezones list.
	timeZones, err := dao.NewPgTimeZoneDAO().GetAll(ctx, "")
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Return the response.
	data := TimeZonesResponseData{
		TimeZones: timeZones,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
