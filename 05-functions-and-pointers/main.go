package main

import "fmt"

// Go functions can return multiple values (commonly result and error)
func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Variadic function: accepts any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Pointers in Go:
// Everything in Go is passed by value (copied).
// To modify the caller's original variable, pass a pointer (*T).
func incrementByValue(val int) {
	val += 10 // modifies only the local copy
}

func incrementByPointer(val *int) {
	*val += 10 // dereference pointer and modify the original memory
}

func main() {
	fmt.Println("=== 05: Functions and Pointers ===")

	// 1. Multiple return values
	res, ok := divide(10, 2)
	fmt.Printf("10 / 2 = %.2f (success: %t)\n", res, ok)

	_, okZero := divide(10, 0)
	fmt.Printf("10 / 0 success: %t\n", okZero)

	// 2. Variadic function
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))

	// 3. First-class functions & Closures
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}
	triple := multiplier(3)
	fmt.Println("Closure triple(7):", triple(7))

	// 4. Pointers (& to get memory address, * to read/modify value at address)
	num := 20
	fmt.Println("Original num:", num)

	incrementByValue(num)
	fmt.Println("After incrementByValue:", num) // Still 20!

	incrementByPointer(&num)
	fmt.Println("After incrementByPointer:", num) // Changed to 30!
}
