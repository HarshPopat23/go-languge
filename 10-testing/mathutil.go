package mathutil

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Factorial calculates the factorial of n (n!)
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// SafeDivide divides a by b, returning an error if b == 0
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
