// Package main demonstrates advanced Go interface concepts including:
// 1. Decoupling code using interfaces (Contracts).
// 2. Dependency Injection / Strategy Pattern (e.g., swapping payment gateways).
// 3. Implementing multiple interface methods.
// 4. Mocking dependencies for testing.
package main

import "fmt"

// PaymentGateway defines the contract for processing payments and refunds.
// Any struct that implements these two methods implicitly satisfies this interface.
type PaymentGateway interface {
	Pay(amount float32) bool
	Refund(amount float32, account string) bool
}

// PaymentService is a high-level service that depends on the PaymentGateway abstraction,
// NOT on concrete implementations (Dependency Inversion Principle).
type PaymentService struct {
	gateway PaymentGateway
}

// NewPaymentService is a constructor function for creating a PaymentService with an injected gateway.
func NewPaymentService(gateway PaymentGateway) *PaymentService {
	return &PaymentService{gateway: gateway}
}

// ProcessCheckout executes a payment using the injected gateway.
func (ps *PaymentService) ProcessCheckout(amount float32) {
	fmt.Printf("[Checkout] Attempting checkout of $%.2f...\n", amount)
	success := ps.gateway.Pay(amount)
	if success {
		fmt.Println("[Checkout] Payment successful!")
	} else {
		fmt.Println("[Checkout] Payment failed!")
	}
}

// ProcessRefund initiates a refund using the injected gateway.
func (ps *PaymentService) ProcessRefund(amount float32, account string) {
	fmt.Printf("[Refund] Processing refund of $%.2f to account %s...\n", amount, account)
	ps.gateway.Refund(amount, account)
}

// -------------------------------------------------------------
// Concrete Implementation 1: Razorpay
// -------------------------------------------------------------
type Razorpay struct {
	APIKey string
}

func (r Razorpay) Pay(amount float32) bool {
	fmt.Printf("  -> [Razorpay] Paid $%.2f via Razorpay Gateway (Key: %s)\n", amount, r.APIKey)
	return true
}

func (r Razorpay) Refund(amount float32, account string) bool {
	fmt.Printf("  -> [Razorpay] Refunded $%.2f to %s\n", amount, account)
	return true
}

// -------------------------------------------------------------
// Concrete Implementation 2: Stripe
// -------------------------------------------------------------
type Stripe struct {
	SecretKey string
}

func (s Stripe) Pay(amount float32) bool {
	fmt.Printf("  -> [Stripe] Paid $%.2f via Stripe (Secret: %s)\n", amount, s.SecretKey)
	return true
}

func (s Stripe) Refund(amount float32, account string) bool {
	fmt.Printf("  -> [Stripe] Refunded $%.2f to %s\n", amount, account)
	return true
}

// -------------------------------------------------------------
// Concrete Implementation 3: PayPal
// -------------------------------------------------------------
type PayPal struct {
	ClientID string
}

func (p PayPal) Pay(amount float32) bool {
	fmt.Printf("  -> [PayPal] Paid $%.2f via PayPal account\n", amount)
	return true
}

func (p PayPal) Refund(amount float32, account string) bool {
	fmt.Printf("  -> [PayPal] Refunded $%.2f to %s\n", amount, account)
	return true
}

// -------------------------------------------------------------
// Concrete Implementation 4: MockPaymentGateway (For Unit Testing)
// -------------------------------------------------------------
type MockPaymentGateway struct{}

func (m MockPaymentGateway) Pay(amount float32) bool {
	fmt.Printf("  -> [Mock Gateway] Simulated payment of $%.2f (No real money moved)\n", amount)
	return true
}

func (m MockPaymentGateway) Refund(amount float32, account string) bool {
	fmt.Printf("  -> [Mock Gateway] Simulated refund of $%.2f to %s\n", amount, account)
	return true
}

func main() {
	fmt.Println("=== 12: Advanced Interfaces & Dependency Injection ===")

	// 1. Using Razorpay gateway
	razorpayGateway := Razorpay{APIKey: "rzp_live_987654"}
	paymentApp := NewPaymentService(razorpayGateway)
	paymentApp.ProcessCheckout(150.75)
	paymentApp.ProcessRefund(50.00, "user_acc_123")

	fmt.Println("--------------------------------------------------")

	// 2. Swapping to Stripe gateway without changing PaymentService business logic!
	stripeGateway := Stripe{SecretKey: "sk_live_abcdef123456"}
	paymentApp = NewPaymentService(stripeGateway)
	paymentApp.ProcessCheckout(299.99)

	fmt.Println("--------------------------------------------------")

	// 3. Swapping to Mock Gateway (for tests / local dev)
	mockGateway := MockPaymentGateway{}
	testPaymentApp := NewPaymentService(mockGateway)
	testPaymentApp.ProcessCheckout(99.00)
}
