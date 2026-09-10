// Package main indicates that this file is the entry point for an executable program.
// In Go, programs meant to run as an executable must belong to package "main"
// and must contain a "func main()".
package main

// Import the "fmt" package from the standard library for formatted I/O.
import "fmt"

func main() {
	// fmt.Println prints text followed by a new line.
	fmt.Println("🎉 Welcome to Go (Golang)!")
	fmt.Println("----------------------------------------")
	fmt.Println("To run this file directly:")
	fmt.Println("  go run ./01-hello-world")
	fmt.Println("To compile it into an executable binary:")
	fmt.Println("  go build -o bin/hello.exe ./01-hello-world")
	fmt.Println("----------------------------------------")
	fmt.Println("Happy learning!")
}
