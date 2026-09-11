// Package user manages user profile models and operations.
package user

import "fmt"

// User represents an application user.
// In Go structs:
// - Fields starting with CAPITAL letters (e.g., ID, Name, Email) are EXPORTED and accessible externally.
// - Fields starting with lowercase letters (e.g., role) are UNEXPORTED and hidden outside this package.
type User struct {
	ID    int
	Name  string
	Email string
	role  string // unexported internal field
}

// NewUser is a constructor function for creating a User with default role settings.
func NewUser(id int, name, email string) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
		role:  "standard_user",
	}
}

// DisplayProfile prints user details.
func (u User) DisplayProfile() {
	fmt.Printf("User Profile: [ID: %d | Name: %s | Email: %s | Role: %s]\n",
		u.ID, u.Name, u.Email, u.role)
}
