// Package main demonstrates Go Channels and the Select statement.
//
// Key Concepts:
// 1. Channels (`chan T`): Typed conduits through which you can send and receive values between goroutines using `<-`.
// 2. Unbuffered Channels: Sends and receives block until both sender and receiver are ready (synchronous rendezvous).
// 3. Buffered Channels: `make(chan T, capacity)` allows sending values without blocking until the buffer is full.
// 4. Channel Direction: `chan<- T` (send-only) and `<-chan T` (receive-only) enforce type safety in functions.
// 5. Closing Channels: `close(ch)` signals that no more values will be sent. Ranging over channels (`for v := range ch`) stops automatically when closed.
// 6. Select Statement: Multiplexes across multiple channel operations, choosing the one that is ready.
package main

import (
	"fmt"
	"time"
)

// 1. Basic worker sending results to an unbuffered channel
func calculateSum(a, b int, resultChan chan<- int) {
	resultChan <- a + b
}

// 2. Producer function sending multiple items and closing channel
func emailProducer(emails chan<- string, count int) {
	for i := 1; i <= count; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		emails <- email
		time.Sleep(50 * time.Millisecond)
	}
	// Always close from sender side when done
	close(emails)
}

// 3. Consumer function receiving items until channel is closed
func emailConsumer(emails <-chan string, done chan<- bool) {
	for email := range emails {
		fmt.Printf("  -> [Email Worker] Sending notification to: %s\n", email)
	}
	done <- true
}

func main() {
	fmt.Println("=== 16: Channels & Select Multiplexing ===")

	// -------------------------------------------------------------
	// 1. Unbuffered Channel (Synchronous Send/Receive)
	// -------------------------------------------------------------
	fmt.Println("\n--- 1. Unbuffered Channel ---")
	sumChan := make(chan int)
	go calculateSum(25, 17, sumChan)
	sum := <-sumChan // Blocks until calculateSum sends the value
	fmt.Printf("Received calculation result: %d\n", sum)

	// -------------------------------------------------------------
	// 2. Buffered Channel & Range Loop
	// -------------------------------------------------------------
	fmt.Println("\n--- 2. Buffered Channel & Pipeline ---")
	emailQueue := make(chan string, 5) // Buffer size 5
	doneChan := make(chan bool)

	go emailProducer(emailQueue, 3)
	go emailConsumer(emailQueue, doneChan)

	<-doneChan // Wait for consumer to finish processing all emails

	// -------------------------------------------------------------
	// 3. Select Statement with Multiple Channels & Timeouts
	// -------------------------------------------------------------
	fmt.Println("\n--- 3. Multiplexing with Select ---")
	fastChan := make(chan string)
	slowChan := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		fastChan <- "Message from Fast Source"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		slowChan <- "Message from Slow Source"
	}()

	// Select on whichever channel receives first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-fastChan:
			fmt.Printf("[Select] Received: %s\n", msg1)
		case msg2 := <-slowChan:
			fmt.Printf("[Select] Received: %s\n", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("[Select] Timed out waiting for channel message!")
		}
	}

	// -------------------------------------------------------------
	// 4. Non-blocking Channel Operations with Default Case
	// -------------------------------------------------------------
	fmt.Println("\n--- 4. Non-blocking Select ---")
	nonBlockChan := make(chan int, 1)
	select {
	case val := <-nonBlockChan:
		fmt.Printf("Received: %d\n", val)
	default:
		fmt.Println("[Non-blocking] No message available immediately in channel, continuing...")
	}
}
