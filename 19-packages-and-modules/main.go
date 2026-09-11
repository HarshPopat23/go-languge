// Package main demonstrates organizing Go code into Packages and Modules.
//
// Key Concepts:
// 1. Modules: A Go module (`go.mod`) is a collection of Go packages that are versioned together.
// 2. Package Organization: Each subdirectory forms an isolated package named after the directory.
// 3. Exported Identifiers (Public): Start with a Capital letter (e.g., `auth.LoginWithCredentials`).
// 4. Unexported Identifiers (Private): Start with a lowercase letter (e.g., `extractSessionToken`).
// 5. Imports: Packages within the same module are imported using `<module_name>/<package_path>`.
package main

import (
	"fmt"

	"github.com/HarshPopat23/go-languge/19-packages-and-modules/auth"
	"github.com/HarshPopat23/go-languge/19-packages-and-modules/user"
)

func main() {
	fmt.Println("=== 19: Packages and Modules in Go ===")

	// 1. Calling exported functions from the `auth` package
	fmt.Println("\n--- 1. Authentication Package ---")
	isLoggedIn := auth.LoginWithCredentials("harsh", "superSecretPass123")
	if isLoggedIn {
		token := auth.GetSession()
		fmt.Printf("Active Session Token: %s\n", token)
	}

	// 2. Using types and methods from the `user` package
	fmt.Println("\n--- 2. User Package ---")
	newUser := user.NewUser(1, "Harsh Popat", "harsh@example.com")
	newUser.DisplayProfile()

	// Accessing exported struct fields directly
	fmt.Printf("Direct access -> Name: %s, Email: %s\n", newUser.Name, newUser.Email)
}
