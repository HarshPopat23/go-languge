package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function for channels
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// Simulate computation time
		time.Sleep(50 * time.Millisecond)
		results <- j * 2
	}
}

// Thread-safe counter using sync.Mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("=== 09: Concurrency (Goroutines, Channels, WaitGroup, Mutex) ===")

	// 1. Basic Goroutine (`go func()`)
	fmt.Println("\n-- 1. Basic Goroutines & sync.WaitGroup --")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			fmt.Printf("Task %d is executing concurrently\n", taskID)
		}(i)
	}
	wg.Wait() // Wait until all 3 goroutines call wg.Done()

	// 2. Channels (Worker Pool Pattern)
	fmt.Println("\n-- 2. Worker Pool with Channels --")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var poolWg sync.WaitGroup
	// Spawn 2 workers
	for w := 1; w <= 2; w++ {
		poolWg.Add(1)
		go worker(w, jobs, results, &poolWg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Closing signals to workers that no more jobs will arrive

	poolWg.Wait()
	close(results)

	// Read results
	for r := range results {
		fmt.Printf("Result received: %d\n", r)
	}

	// 3. Mutex: Preventing Race Conditions
	fmt.Println("\n-- 3. Thread-safe Mutex Counter --")
	counter := SafeCounter{}
	var counterWg sync.WaitGroup

	// 100 concurrent increments
	for i := 0; i < 100; i++ {
		counterWg.Add(1)
		go func() {
			defer counterWg.Done()
			counter.Increment()
		}()
	}
	counterWg.Wait()
	fmt.Printf("Final Counter Value (expected 100): %d\n", counter.Value())

	// 4. Select Statement (multiplexing channels with timeout)
	fmt.Println("\n-- 4. Select with Timeout --")
	ch := make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "Fast response!"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timed out waiting for message!")
	}
}
