package mathutil

import (
	"errors"
	"testing"
)

// Unit Test using Table-Driven Testing pattern (Idiomatic Go)
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero value", 0, 5, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -5, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// Test with error handling
func TestSafeDivide(t *testing.T) {
	// Sub-test 1: Successful division
	t.Run("valid division", func(t *testing.T) {
		got, err := SafeDivide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 5.0 {
			t.Errorf("SafeDivide(10, 2) = %f; want 5.0", got)
		}
	})

	// Sub-test 2: Division by zero
	t.Run("division by zero", func(t *testing.T) {
		_, err := SafeDivide(10, 0)
		if err == nil {
			t.Fatal("expected error for division by zero, got nil")
		}
		if !errors.Is(err, ErrDivideByZero) {
			t.Errorf("got error %v; want %v", err, ErrDivideByZero)
		}
	})
}

// Benchmark test: measures execution time and allocations
// Run with: go test -bench=. ./10-testing
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(15)
	}
}
