package errcode

// API error codes.
const (
	ReqHeaderValidationFailed    = "40001"
	ReqParamValidationFailed     = "40002"
	AuthorizationEmpty           = "40101"
	AuthorizationFormatInvalid   = "40102"
	AuthorizationTokenInvalid    = "40103"
	AuthorizationTokenExpired    = "40104"
	AuthorizationNotTokenOwner   = "40105"
	AuthorizationUserNotFound    = "40106"
	AuthorizationUserNotVerified = "40107"
	PermissionDenied             = "40301"
	ItemNotFound                 = "40401"
	FileNotFound                 = "40401"

	APIKeyEmpty                = "49101"
	APIKeyInvalid              = "49102"
	APIKeyNotFound             = "49103"
	APIKeyAppPlatformInvalid   = "49104"
	APIKeyAppIdentifierInvalid = "49105"
	APIKeyExpired              = "49106"
	APIKeyDisabled             = "49107"

	InternalAPIKeyValidationFailed = "50001"
	InternalIllegalArgument        = "50002"

	Other = "99999"
)
