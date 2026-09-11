// Package main demonstrates Goroutines and synchronization using sync.WaitGroup.
//
// Key Concepts:
// 1. Goroutine: An extremely lightweight, non-blocking execution thread managed by the Go runtime scheduler (m:n scheduler).
// 2. Launching: Any function can be run as a concurrent goroutine using the `go` keyword (e.g., `go task()`).
// 3. sync.WaitGroup: The standard way to wait for a collection of concurrent goroutines to complete.
//   - `wg.Add(n)`: Increments the counter by n (number of tasks to wait for).
//   - `wg.Done()`: Decrements the counter by 1 (typically called with `defer wg.Done()`).
//   - `wg.Wait()`: Blocks execution until the counter reaches 0.
package main

import (
	"fmt"
	"sync"
	"time"
)

// performTask simulates an asynchronous worker performing a job.
// Note: sync.WaitGroup MUST ALWAYS be passed as a pointer (*sync.WaitGroup) to avoid copying the internal state!
func performTask(taskID int, duration time.Duration, wg *sync.WaitGroup) {
	// defer ensures wg.Done() is called even if the function panics or returns early
	defer wg.Done()

	fmt.Printf("[Worker %d] Started task (duration: %v)\n", taskID, duration)
	time.Sleep(duration)
	fmt.Printf("[Worker %d] Finished task!\n", taskID)
}

func main() {
	fmt.Println("=== 15: Goroutines & sync.WaitGroup ===")

	var wg sync.WaitGroup

	// 1. Running multiple named concurrent tasks
	fmt.Println("\n--- Starting Concurrent Named Workers ---")
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		// Increment WaitGroup counter BEFORE launching the goroutine
		wg.Add(1)
		go performTask(i, time.Duration(i*100)*time.Millisecond, &wg)
	}

	// 2. Running Anonymous Goroutines
	fmt.Println("\n--- Starting Anonymous Goroutines ---")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// IMPORTANT: Pass loop variable `i` as an argument to avoid concurrency variable capture issues
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  -> [Anon Routine %d] Running concurrently...\n", id)
			time.Sleep(150 * time.Millisecond)
			fmt.Printf("  -> [Anon Routine %d] Done!\n", id)
		}(i)
	}

	// wg.Wait() blocks the main goroutine until all Add() counters are decremented to 0 by Done()
	fmt.Println("\n[Main] Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("[Main] All workers completed successfully! Program exiting.")
}
