package api

import (
	"encoding/json"
	"net/http"

	"github.com/jonylim/basego/internal/pkg/common/api/errcode"
	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// RequestParam defines interface for request parameters.
type RequestParam interface {
	Validate() (msg, field string)
}

// DecodeBodyJSON decodes JSON-encoded body from an HTTP request into v and validates the parameters.
func DecodeBodyJSON(w http.ResponseWriter, r *http.Request, reqID string, v RequestParam) (ok bool) {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		logger.Error("api:"+reqID, err.Error())
		msg := "Request body format is invalid"
		response := NewAPIResponseWithError(reqID, errcode.ReqParamValidationFailed, msg)
		SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
	} else if msg, field := v.Validate(); msg != "" {
		response := NewAPIResponseWithErrorField(reqID, errcode.ReqParamValidationFailed, msg, field)
		SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
	} else {
		ok = true
	}
	return
}
