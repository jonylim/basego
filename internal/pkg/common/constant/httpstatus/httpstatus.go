package httpstatus

// HTTP status codes as registered with IANA.
// See: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	OK                   = 200 // RFC 7231, 6.3.1
	Created              = 201 // RFC 7231, 6.3.2
	Accepted             = 202 // RFC 7231, 6.3.3
	NonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	NoContent            = 204 // RFC 7231, 6.3.5
	ResetContent         = 205 // RFC 7231, 6.3.6
	PartialContent       = 206 // RFC 7233, 4.1
	Multi                = 207 // RFC 4918, 11.1
	AlreadyReported      = 208 // RFC 5842, 7.1
	IMUsed               = 226 // RFC 3229, 10.4.1

	BadRequest                  = 400 // RFC 7231, 6.5.1
	Unauthorized                = 401 // RFC 7235, 3.1
	Forbidden                   = 403 // RFC 7231, 6.5.3
	NotFound                    = 404 // RFC 7231, 6.5.4
	MethodNotAllowed            = 405 // RFC 7231, 6.5.5
	NotAcceptable               = 406 // RFC 7231, 6.5.6
	RequestTimeout              = 408 // RFC 7231, 6.5.7
	Conflict                    = 409 // RFC 7231, 6.5.8
	Gone                        = 410 // RFC 7231, 6.5.9
	LengthRequired              = 411 // RFC 7231, 6.5.10
	RequestEntityTooLarge       = 413 // RFC 7231, 6.5.11
	RequestURITooLong           = 414 // RFC 7231, 6.5.12
	UnsupportedMediaType        = 415 // RFC 7231, 6.5.13
	RangeNotSatisfiable         = 416 // RFC 7233, 4.4
	Locked                      = 423 // RFC 4918, 11.3
	UpgradeRequired             = 426 // RFC 7231, 6.5.15
	TooManyRequests             = 429 // RFC 6585, 4
	RequestHeaderFieldsTooLarge = 431 // RFC 6585, 5

	InternalServerError     = 500 // RFC 7231, 6.6.1
	NotImplemented          = 501 // RFC 7231, 6.6.2
	BadGateway              = 502 // RFC 7231, 6.6.3
	ServiceUnavailable      = 503 // RFC 7231, 6.6.4
	GatewayTimeout          = 504 // RFC 7231, 6.6.5
	HTTPVersionNotSupported = 505 // RFC 7231, 6.6.6
	InsufficientStorage     = 507 // RFC 4918, 11.5
)

const (
	// APIKeyInvalid means the API key specified in request header "API-Key" is missing or invalid.
	APIKeyInvalid = 491
)
