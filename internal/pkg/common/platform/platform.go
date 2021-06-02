package platform

// Platform is string describing an app platform.
const (
	SERVER  = "server"
	ANDROID = "android"
	IOS     = "ios"
	WEB     = "web"
)

// IsValidAll checks if the given platform is valid.
func IsValidAll(platform string) bool {
	return platform == SERVER || platform == ANDROID || platform == IOS || platform == WEB
}

// IsValidServer checks if the given platform is valid for server.
func IsValidServer(platform string) bool {
	return platform == SERVER
}

// IsValidClient checks if the given platform is valid for client.
func IsValidClient(platform string) bool {
	return platform == ANDROID || platform == IOS || platform == WEB
}
