package requestheader

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jonylim/basego/internal/pkg/common/platform"
)

// APIRequestHeader defines data passed to request header.
type APIRequestHeader struct {
	APIKey         string
	AppIdentifier  string
	Authorization  string
	ContentType    string
	DeviceID       string
	DeviceModel    string
	DevicePlatform string
	UserAgent      string
}

// Parse parses the API request headers.
func Parse(r *http.Request) APIRequestHeader {
	h := APIRequestHeader{
		APIKey:         r.Header.Get("API-Key"),
		AppIdentifier:  r.Header.Get("App-Identifier"),
		Authorization:  r.Header.Get("Authorization"),
		ContentType:    r.Header.Get("Content-Type"),
		DeviceID:       r.Header.Get("Device-Identifier"),
		DeviceModel:    r.Header.Get("Device-Model"),
		DevicePlatform: r.Header.Get("Device-Platform"),
		UserAgent:      r.UserAgent(),
	}
	if h.DevicePlatform == platform.WEB {
		h.AppIdentifier = r.Header.Get("Origin")
	}
	return h
}

// CheckRequired checks required request headers.
func CheckRequired(h APIRequestHeader) error {
	var keys []string
	if h.APIKey == "" {
		keys = append(keys, "API-Key")
	}
	if h.Authorization == "" {
		keys = append(keys, "Authorization")
	}
	if h.DevicePlatform != "" {
		if h.DevicePlatform != platform.WEB {
			if h.DeviceID == "" {
				keys = append(keys, "Device-Identifier")
			}
			if h.DeviceModel == "" {
				keys = append(keys, "Device-Model")
			}
		}
	} else {
		keys = append(keys, "Device-Platform")
	}
	if len(keys) != 0 {
		return errors.New("Request headers are required (" + strings.Join(keys, ", ") + ")")
	}
	if !platform.IsValidClient(h.DevicePlatform) {
		return errors.New("Request header is invalid (Device-Platform)")
	}
	return nil
}
