// Package main demonstrates Generics (Parametric Polymorphism) in Go (introduced in Go 1.18+).
// Generics allow writing reusable, type-safe functions and data structures without
// sacrificing compile-time type safety or resorting to empty interface (`any`) type assertions.
package main

import (
	"errors"
	"fmt"
)

// -------------------------------------------------------------
// 1. Generic Functions with Built-in Constraints
// -------------------------------------------------------------

// PrintSlice prints each item of a slice of any type.
// [T any] allows T to be any type (alias for interface{}).
func PrintSlice[T any](label string, items []T) {
	fmt.Printf("%s: [ ", label)
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
	fmt.Println("]")
}

// FindIndex searches for a target element in a slice and returns its index.
// [T comparable] allows T to be any type supporting equality operators (`==` and `!=`).
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// -------------------------------------------------------------
// 2. Custom Type Constraints
// -------------------------------------------------------------

// Number is a constraint interface that permits any integer or floating-point type.
// The tilde `~` allows custom types whose underlying type is int or float64.
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Sum adds all elements in a numeric slice together.
func Sum[T Number](numbers []T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// -------------------------------------------------------------
// 3. Generic Data Structures (Generic Stack)
// -------------------------------------------------------------

// Stack is a Last-In-First-Out (LIFO) generic container.
type Stack[T any] struct {
	elements []T
}

// Push adds an item onto the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T // zero value for type T
		return zero, errors.New("stack is empty")
	}
	lastIdx := len(s.elements) - 1
	item := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]
	return item, nil
}

// Peek returns the top item without removing it.
func (s *Stack[T]) Peek() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// Len returns the number of items currently in the stack.
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func main() {
	fmt.Println("=== 14: Generics in Go (Go 1.18+) ===")

	// 1. Generic Functions with type inference
	intSlice := []int{10, 20, 30, 40}
	strSlice := []string{"Go", "Rust", "TypeScript", "Python"}
	boolSlice := []bool{true, false, true}

	PrintSlice("Integers", intSlice)
	PrintSlice("Strings", strSlice)
	PrintSlice("Booleans", boolSlice)

	// 2. Using FindIndex with comparable constraint
	fmt.Println("\nSearching with FindIndex:")
	fmt.Printf("  Index of 'TypeScript': %d\n", FindIndex(strSlice, "TypeScript"))
	fmt.Printf("  Index of 30: %d\n", FindIndex(intSlice, 30))
	fmt.Printf("  Index of 99 (missing): %d\n", FindIndex(intSlice, 99))

	// 3. Custom Number constraint
	intTotal := Sum([]int{1, 2, 3, 4, 5})
	floatTotal := Sum([]float64{1.5, 2.5, 3.5})
	fmt.Printf("\nSum of ints: %d\n", intTotal)
	fmt.Printf("Sum of floats: %.2f\n", floatTotal)

	// 4. Generic Stack data structure
	fmt.Println("\nGeneric Stack Demo:")
	strStack := Stack[string]{}
	strStack.Push("first")
	strStack.Push("second")
	strStack.Push("third")

	fmt.Printf("  Stack size: %d\n", strStack.Len())
	top, _ := strStack.Peek()
	fmt.Printf("  Top item (Peek): %s\n", top)

	for strStack.Len() > 0 {
		item, _ := strStack.Pop()
		fmt.Printf("  Popped: %s (Remaining size: %d)\n", item, strStack.Len())
	}
}
