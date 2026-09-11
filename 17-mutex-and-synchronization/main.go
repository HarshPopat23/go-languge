// Package main demonstrates Mutual Exclusion (sync.Mutex), Read-Write Mutex (sync.RWMutex),
// and Atomic Operations (sync/atomic) in Go to prevent Race Conditions.
//
// Key Concepts:
// 1. Race Condition: When multiple goroutines concurrently read and write to the same shared memory without synchronization.
// 2. sync.Mutex: Provides mutual exclusion. Only one goroutine can hold the lock at a time via `mu.Lock()` and `mu.Unlock()`.
// 3. sync.RWMutex: Allows multiple concurrent readers (`mu.RLock()`), but only one writer at a time (`mu.Lock()`).
// 4. sync/atomic: Low-level atomic memory operations that execute in a single CPU instruction without mutex locking overhead.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// -------------------------------------------------------------
// 1. sync.Mutex: Protecting Shared Mutable State
// -------------------------------------------------------------
type Post struct {
	Title string
	Views int
	mu    sync.Mutex // Mutex guarding the Views field
}

// IncrementViews safely increments views concurrently
func (p *Post) IncrementViews(wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire lock before accessing shared state
	p.mu.Lock()
	p.Views++ // Critical section
	p.mu.Unlock()
}

// -------------------------------------------------------------
// 2. sync.RWMutex: Concurrent Readers & Exclusive Writer
// -------------------------------------------------------------
type SafeCache struct {
	store map[string]string
	rw    sync.RWMutex
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		store: make(map[string]string),
	}
}

// Set writes a key-value pair (Exclusive Lock)
func (c *SafeCache) Set(key, value string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.store[key] = value
}

// Get reads a value for a key (Shared Read Lock - Multiple goroutines can read simultaneously)
func (c *SafeCache) Get(key string) (string, bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	val, ok := c.store[key]
	return val, ok
}

func main() {
	fmt.Println("=== 17: Mutex & Concurrency Synchronization ===")

	// 1. Mutex Protection Demo
	fmt.Println("\n--- 1. Testing sync.Mutex ---")
	post := &Post{Title: "Mastering Concurrency in Go", Views: 0}
	var wg sync.WaitGroup

	numConcurrentViews := 1000
	for i := 0; i < numConcurrentViews; i++ {
		wg.Add(1)
		go post.IncrementViews(&wg)
	}

	wg.Wait()
	fmt.Printf("Post '%s' Final View Count: %d (Expected: %d)\n", post.Title, post.Views, numConcurrentViews)

	// 2. RWMutex Demo
	fmt.Println("\n--- 2. Testing sync.RWMutex Cache ---")
	cache := NewSafeCache()
	cache.Set("session_id", "xyz-789-token")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			val, _ := cache.Get("session_id")
			fmt.Printf("  -> [Reader %d] Cached session_id = %s\n", readerID, val)
		}(i)
	}
	wg.Wait()

	// 3. sync/atomic Demo
	fmt.Println("\n--- 3. Testing sync/atomic Operations ---")
	var atomicCounter int64 = 0

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Atomic increment without needing a Mutex lock
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg.Wait()

	// Atomic load ensures memory visibility across CPU cores
	finalCount := atomic.LoadInt64(&atomicCounter)
	fmt.Printf("Atomic Counter Value: %d (Expected: 500)\n", finalCount)
	time.Sleep(10 * time.Millisecond)
}
