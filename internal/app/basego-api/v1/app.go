package v1

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/jonylim/basego/internal/app/basego-api/v1/endpoint/accountapi"
	"github.com/jonylim/basego/internal/app/basego-api/v1/endpoint/authapi"
	"github.com/jonylim/basego/internal/app/basego-api/v1/endpoint/clientapi"
	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"

	"github.com/julienschmidt/httprouter"
)

// APIPrefix defines the prefix path for API URLs.
const APIPrefix = "/v1/"

// Define endpoints and handles here.
var authAPIs = map[string]authapi.Handle{
	"access_token/request": authapi.AccessTokenRequest,
	"access_token/refresh": authapi.AccessTokenRefresh,
}
var clientAPIs = map[string]clientapi.Handle{
	"server_time":                       clientapi.ServerTime,
	"register":                          clientapi.Register,
	"account_verification/resend_email": clientapi.AccountVerificationResendEmail,
	"account_verification/submit":       clientapi.AccountVerificationSubmit,
	"reset_password/request_token":      clientapi.ResetPasswordRequestToken,
	"reset_password/verify_token":       clientapi.ResetPasswordVerifyToken,
	"reset_password/set_password":       clientapi.ResetPasswordSetPassword,
}
var accountAPIs = map[string]accountapi.Handle{
	"countries":                accountapi.Countries,
	"time_zones":               accountapi.TimeZones,
	"profile/get":              accountapi.AccountProfileGet,
	"profile/accept_tos":       accountapi.AccountProfileAcceptTOS,
	"security/change_password": accountapi.SecurityChangePassword,
	"logout":                   accountapi.Logout,
}
var mapAPIs = map[string]interface{}{
	"auth":    authAPIs,
	"client":  clientAPIs,
	"account": accountAPIs,
}

var (
	errDatabase = errors.New("An error occurred while processing your request")
	errInternal = errors.New("An error occurred while processing your request")
)

var allowOriginURL string

// RouteAPIs configure the router for APIs.
func RouteAPIs(router *httprouter.Router) {
	allowOriginURL = strings.Trim(os.Getenv(envvar.FrontendURL), "/")

	// Init APIs.
	authapi.Init()
	clientapi.Init()
	accountapi.Init()

	for apiType, apiList := range mapAPIs {
		apiPrefix := APIPrefix + apiType + "/"

		switch apiType {
		case "auth":
			for apiName, apiHandle := range apiList.(map[string]authapi.Handle) {
				var h = apiHandle
				router.OPTIONS(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
				})
				router.POST(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
					authapi.HandleRequest(w, r, p, h)
				})
			}
			break

		case "client":
			for apiName, apiHandle := range apiList.(map[string]clientapi.Handle) {
				var h = apiHandle
				router.OPTIONS(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
				})
				router.POST(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
					clientapi.HandleRequest(w, r, p, h)
				})
			}
			break

		case "account":
			for apiName, apiHandle := range apiList.(map[string]accountapi.Handle) {
				var h = apiHandle
				router.OPTIONS(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
				})
				router.POST(apiPrefix+apiName, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					allowCORS(w, r)
					accountapi.HandleRequest(w, r, p, h)
				})
			}
			break
		}
	}
}

func allowCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", allowOriginURL)
	if acrh := r.Header.Get("Access-Control-Request-Headers"); acrh != "" {
		w.Header().Set("Access-Control-Allow-Headers", acrh)
	} else {
		w.Header().Set("Access-Control-Allow-Headers", "*")
	}
	w.Header().Set("Vary", "Origin")
}
