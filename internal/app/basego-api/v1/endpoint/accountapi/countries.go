/**
 * @api           {post} /v1/account/countries Get Country List
 * @apiVersion    1.0.0
 * @apiName       Countries
 * @apiGroup      AccountAPI
 * @apiPermission account
 *
 * @apiDescription Get the country list.
 *
 * @apiSuccess {object[]} countries              The list of countries.
 * @apiSuccess {id}       countries.id           The country ID.
 * @apiSuccess {string}   countries.commonName   The country's common name.
 * @apiSuccess {string}   countries.officialName The country's official name.
 * @apiSuccess {string}   countries.iso2Code     The country code based on ISO 3166-1 alpha-2.
 * @apiSuccess {string}   countries.iso3Code     The country code based on ISO 3166-1 alpha-3.
 * @apiSuccess {string}   countries.callingCode  The country's international calling code for phone number.
 * @apiSuccess {string}   countries.currencyCode The country's currency (e.g.: "IDR").
 * @apiSuccess {boolean}  countries.isEnabled    If the country is enabled.
 * @apiSuccess {boolean}  countries.isHidden     If the country is hidden.
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
 *         "countries": [
 *            {
 *              "id": 1,
 *              "commonName": "Indonesia",
 *              "officialName": "Republic of Indonesia",
 *              "iso2Code": "ID",
 *              "iso3Code": "IDN",
 *              "callingCode": "62",
 *              "currencyCode": "IDR",
 *              "isEnabled": true,
 *              "isHidden": false
 *            },
 *            ...
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

// CountriesResponseData represents response data of Account API "Get Country List".
type CountriesResponseData struct {
	Countries []model.XCountry `json:"countries"`
}

// Countries returns list of countries.
func Countries(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx Context) {
	logger.Trace(ctx.ReqTag, "Handle: accountapi.Countries")

	// Get the country list.
	countries, err := dao.NewXCountryDAO().GetActiveCountryList()
	if err != nil {
		response := api.NewAPIResponseWithError(ctx.ReqID, errcode.Other, errDatabase.Error())
		api.SendResponseJSONWithStatusCode(w, response, httpstatus.InternalServerError)
		return
	}

	// Return the response.
	data := CountriesResponseData{
		Countries: countries,
	}
	response := api.NewAPIResponse(ctx.ReqID)
	response.SetData(data)
	api.SendResponseJSON(w, response)
}
