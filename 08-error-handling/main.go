package main

import (
	"errors"
	"fmt"
)

// Sentinel errors (well-known static error values)
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrNegativeAmount = errors.New("amount must be greater than zero")

type BankAccount struct {
	Owner   string
	Balance float64
}

// Idiomatic Go function returning (result, error)
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		// Wrapping an error with additional context using %w
		return fmt.Errorf("withdraw failed: %w", ErrNegativeAmount)
	}
	if amount > b.Balance {
		return fmt.Errorf("withdraw of $%.2f failed: %w (current balance: $%.2f)",
			amount, ErrInsufficientBalance, b.Balance)
	}
	b.Balance -= amount
	return nil
}

// Demonstrating `defer`: Deferred statements run LIFO when surrounding function returns
func demonstrateDefer() {
	fmt.Println("Starting function with deferred calls...")
	defer fmt.Println("  -> Defer 1 (runs last)")
	defer fmt.Println("  -> Defer 2 (runs second)")
	defer fmt.Println("  -> Defer 3 (runs first)")
	fmt.Println("Doing work inside function...")
}

// Demonstrating `recover` from a panic (rarely used except for boundary guards/web middleware)
func safeDivision(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Safely recovered from panic: %v\n", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("division by zero is not allowed!")
	}
	return a / b
}

func main() {
	fmt.Println("=== 08: Error Handling, Defer, Panic & Recover ===")

	account := BankAccount{Owner: "Alice", Balance: 100.0}

	// 1. Handling standard errors
	err := account.Withdraw(150.0)
	if err != nil {
		fmt.Printf("Error encountered: %v\n", err)

		// errors.Is unwraps and checks for sentinel errors
		if errors.Is(err, ErrInsufficientBalance) {
			fmt.Println("-> Caught ErrInsufficientBalance! Advise user to add funds.")
		}
	}

	// Successful operation
	err = account.Withdraw(40.0)
	if err == nil {
		fmt.Printf("Withdrawal successful! Remaining balance: $%.2f\n", account.Balance)
	}

	// 2. Defer execution demonstration
	fmt.Println("\nDefer demonstration:")
	demonstrateDefer()

	// 3. Panic and recover
	fmt.Println("\nPanic & Recover demonstration:")
	val := safeDivision(10, 0)
	fmt.Println("safeDivision returned:", val)
}
