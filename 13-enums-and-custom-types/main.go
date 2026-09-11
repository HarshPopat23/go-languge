// Package main demonstrates how to model Enums (Enumerations) and Custom Types in Go.
// Go does not have a dedicated `enum` keyword. Instead, Go idioms use:
// 1. Defined custom types (type MyEnum int or type MyEnum string)
// 2. Constants grouped with `const (...)`
// 3. The `iota` identifier for auto-incrementing integer enums
// 4. Methods attached to custom types (e.g., validation, Stringer interface)
package main

import "fmt"

// -------------------------------------------------------------
// Pattern 1: String-based Enum
// -------------------------------------------------------------
// OrderStatus represents the current state of a customer order.
type OrderStatus string

const (
	StatusReceived  OrderStatus = "Received"
	StatusConfirmed OrderStatus = "Confirmed"
	StatusPrepared  OrderStatus = "Prepared"
	StatusDelivered OrderStatus = "Delivered"
	StatusCancelled OrderStatus = "Cancelled"
)

// IsTerminal returns true if the order status is in a final state.
func (s OrderStatus) IsTerminal() bool {
	return s == StatusDelivered || s == StatusCancelled
}

// UpdateOrderStatus demonstrates receiving a strongly-typed enum parameter.
func UpdateOrderStatus(orderID int, newStatus OrderStatus) {
	fmt.Printf("[Order #%d] Status updated to: %s (Is Completed? %t)\n",
		orderID, newStatus, newStatus.IsTerminal())
}

// -------------------------------------------------------------
// Pattern 2: Integer-based Enum using `iota`
// -------------------------------------------------------------
// Priority represents task importance.
type Priority int

const (
	PriorityLow      Priority = iota // 0
	PriorityMedium                   // 1
	PriorityHigh                     // 2
	PriorityCritical                 // 3
)

// Implementing the fmt.Stringer interface converts integer enums to human-readable strings when printed.
func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "LOW"
	case PriorityMedium:
		return "MEDIUM"
	case PriorityHigh:
		return "HIGH"
	case PriorityCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

func main() {
	fmt.Println("=== 13: Enums and Custom Types in Go ===")

	// 1. Working with String Enums
	orderID := 101
	UpdateOrderStatus(orderID, StatusReceived)
	UpdateOrderStatus(orderID, StatusConfirmed)
	UpdateOrderStatus(orderID, StatusPrepared)
	UpdateOrderStatus(orderID, StatusDelivered)

	fmt.Println("--------------------------------------------------")

	// 2. Working with Integer/Iota Enums
	fmt.Println("Priority Levels:")
	fmt.Printf("  PriorityLow: %d -> %s\n", PriorityLow, PriorityLow)
	fmt.Printf("  PriorityMedium: %d -> %s\n", PriorityMedium, PriorityMedium)
	fmt.Printf("  PriorityHigh: %d -> %s\n", PriorityHigh, PriorityHigh)
	fmt.Printf("  PriorityCritical: %d -> %s\n", PriorityCritical, PriorityCritical)

	currentPriority := PriorityHigh
	if currentPriority >= PriorityHigh {
		fmt.Println("⚠️  Urgent action required for High/Critical priority!")
	}
}
