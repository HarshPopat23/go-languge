package main

import "fmt"

// User defines a custom data structure.
// Fields starting with an uppercase letter are exported (public); lowercase are unexported (private).
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Constructor function: Idiomatic Go pattern to create and initialize structs
func NewUser(id int, name, email string, age int) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

// Value receiver method: operates on a COPY of the User struct (read-only)
func (u User) DisplayInfo() {
	fmt.Printf("[User #%d] %s <%s>, Age: %d\n", u.ID, u.Name, u.Email, u.Age)
}

// Pointer receiver method: operates on the ACTUAL struct in memory (can mutate state)
func (u *User) CelebrateBirthday() {
	u.Age++
}

// Struct embedding (Composition over inheritance):
// Admin embeds User, inheriting all its fields and methods!
type Admin struct {
	User
	Role string
}

func main() {
	fmt.Println("=== 06: Structs and Methods ===")

	// 1. Creating a struct using constructor
	user1 := NewUser(1, "Alice Smith", "alice@example.com", 28)
	user1.DisplayInfo()

	// 2. Calling pointer receiver method
	user1.CelebrateBirthday()
	fmt.Println("After birthday:")
	user1.DisplayInfo()

	// 3. Struct embedding (Composition)
	admin := Admin{
		User: User{
			ID:    99,
			Name:  "Super Admin",
			Email: "admin@corp.internal",
			Age:   35,
		},
		Role: "Platform Engineer",
	}

	// Admin can directly access embedded User fields and methods
	admin.DisplayInfo()
	fmt.Printf("Admin Role: %s\n", admin.Role)
}
