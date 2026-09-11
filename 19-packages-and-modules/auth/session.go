// Package auth continues in this file.
// Notice that all .go files in the same directory MUST declare the exact same package name!
// They share package-level scope, so unexported functions from credentials.go are visible here.
package auth

// extractSessionToken is an UNEXPORTED function only accessible within the auth package.
func extractSessionToken() string {
	return "sess_token_987654321_active"
}

// GetSession is an EXPORTED function returning the current active session token.
func GetSession() string {
	return extractSessionToken()
}
