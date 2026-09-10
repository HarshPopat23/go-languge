package main

import "fmt"

// Constants cannot be modified after declaration
const (
	AppName = "GoLearningLab"
	Pi      = 3.14159
	StatusActive = iota // iota starts at 0 and increments: 0
	StatusPending       // 1
	StatusClosed        // 2
)

func main() {
	fmt.Println("=== 02: Variables, Constants, and Types ===")

	// 1. Explicit variable declaration with `var`
	var age int = 25
	var name string = "Alice"
	var isDeveloper bool = true

	// 2. Type inference with `var`
	var language = "Go"

	// 3. Short variable declaration syntax `:=` (only valid inside functions)
	version := 1.26
	framework := "Standard Library"

	// 4. Zero values: Variables declared without an explicit value get initialized to zero values
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultString string // "" (empty string)
	var defaultBool bool     // false

	fmt.Printf("Developer: %s (age %d), learning %s (v%.2f)\n", name, age, language, version)
	fmt.Printf("Is dev: %t | Framework: %s\n", isDeveloper, framework)
	fmt.Printf("Zero values -> int: %d, float: %f, string: %q, bool: %t\n",
		defaultInt, defaultFloat, defaultString, defaultBool)

	// 5. Type conversion (Go NEVER performs implicit type casting!)
	var count int = 42
	var countFloat float64 = float64(count)
	fmt.Printf("Converted count: %d -> %.1f\n", count, countFloat)

	// 6. Constants & iota enum pattern
	fmt.Printf("App: %s, Pi: %.4f\n", AppName, Pi)
	fmt.Printf("Status Enums -> Active: %d, Pending: %d, Closed: %d\n",
		StatusActive, StatusPending, StatusClosed)
}
