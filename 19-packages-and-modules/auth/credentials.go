// Package auth provides user authentication and session management utilities.
// In Go, package names should be short, lowercase, and concise.
package auth

import "fmt"

// LoginWithCredentials authenticates a user with their username and password.
// In Go, identifiers starting with a CAPITAL letter (PascalCase) are EXPORTED (public)
// and can be accessed from outside this package.
func LoginWithCredentials(username, password string) bool {
	// Internal validation logic
	if username == "" || password == "" {
		fmt.Println("[auth] Error: username and password cannot be empty")
		return false
	}

	fmt.Printf("[auth] Successfully logged in user: '%s'\n", username)
	return true
}

// validatePassword is an UNEXPORTED (private) helper function.
// Identifiers starting with a lowercase letter are only accessible within the 'auth' package.
func validatePassword(password string) bool {
	return len(password) >= 6
}
