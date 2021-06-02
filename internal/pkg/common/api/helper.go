package api

import (
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

// GetClientIPAddress returns the client's original IP address.
func GetClientIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		items := strings.Split(ip, ",")
		for i := len(items) - 1; i >= 0; i-- {
			item := strings.Trim(items[i], " ")
			if item != "127.0.0.1" && !strings.HasPrefix(item, "10.") {
				return item
			}
		}
	}
	// Try parsing remote address.
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	// Fallback method.
	if strings.HasPrefix(r.RemoteAddr, "[") {
		parts := strings.Split(r.RemoteAddr, "]")
		return parts[0][1:]
	}
	parts := strings.Split(r.RemoteAddr, ":")
	return parts[0]
}

// CreateReqID returns a new generated request ID.
func CreateReqID() string {
	return uuid.NewV4().String()[:8]
}

// ValidateRouterParamsFromPath parses a router path and check the router parameter values.
// All parsed parameter names must not be empty. Otherwise, an HTTP response 404 is automatically sent.
// The boolean is false if the validation fails and the request should not be processed any further.
func ValidateRouterParamsFromPath(w http.ResponseWriter, p httprouter.Params, routerPath string) bool {
	for _, n := range getParamNamesFromRouterPath(routerPath) {
		if p.ByName(n) == "" {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return false
		}
	}
	return true
}

func getParamNamesFromRouterPath(path string) []string {
	segments := strings.Split(path, "/")
	names := make([]string, 0)
	for _, s := range segments {
		if len(s) > 1 && s[0] == ':' {
			names = append(names, s[1:])
		}
	}
	return names
}
