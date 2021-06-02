package api

import (
	"encoding/json"
	"net/http"

	"github.com/jonylim/basego/internal/pkg/common/constant/httpstatus"
	"github.com/jonylim/basego/internal/pkg/common/logger"
)

// Response represents API response payload.
type Response struct {
	Status int           `json:"status"`
	ReqID  string        `json:"reqID"`
	Err    ResponseError `json:"error"`
	Data   interface{}   `json:"data"`
}

// ResponseError is an error in an API request.
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

// ResponseData is the payload data to be returned for an API request.
type ResponseData struct{}

// NewAPIResponse returns new instance of API response.
func NewAPIResponse(reqID string) *Response {
	r := &Response{
		Status: httpstatus.OK,
		ReqID:  reqID,
		Err:    ResponseError{},
		Data:   ResponseData{},
	}
	return r
}

// NewAPIResponseWithError returns new instance of API response with error.
func NewAPIResponseWithError(reqID, code, msg string) *Response {
	return &Response{
		Status: httpstatus.OK,
		ReqID:  reqID,
		Err:    ResponseError{code, msg, ""},
		Data:   ResponseData{},
	}
}

// NewAPIResponseWithErrorField returns new instance of API response with error.
func NewAPIResponseWithErrorField(reqID, code, msg, field string) *Response {
	return &Response{
		Status: httpstatus.OK,
		ReqID:  reqID,
		Err:    ResponseError{code, msg, field},
		Data:   ResponseData{},
	}
}

// SetStatus sets the HTTP response status code to the API response.
func (r *Response) SetStatus(status int) {
	r.Status = status
}

// SetError sets the error message to the API response.
func (r *Response) SetError(code, msg string) {
	r.Err = ResponseError{code, msg, ""}
}

// SetErrorField sets the error message to the API response.
func (r *Response) SetErrorField(code, msg, field string) {
	r.Err = ResponseError{code, msg, field}
}

// SetData sets the API response data. This will overwrite previously set data.
func (r *Response) SetData(data interface{}) {
	r.Data = data
}

// SendResponseJSON writes API response into JSON.
func SendResponseJSON(w http.ResponseWriter, r *Response) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		logger.Error("api", logger.FromError(err))
	}
	return err
}

// SendResponseJSONWithStatusCode writes API response into JSON.
func SendResponseJSONWithStatusCode(w http.ResponseWriter, r *Response, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	r.Status = statusCode
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		logger.Error("api", logger.FromError(err))
	}
	return err
}
