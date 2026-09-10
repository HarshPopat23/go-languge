package main

import "fmt"

func main() {
	fmt.Println("=== 04: Slices, Arrays, and Maps ===")

	// 1. Arrays (fixed length, rarely used directly in Go)
	var fixedArray [3]int = [3]int{10, 20, 30}
	fmt.Println("Fixed Array:", fixedArray, "Length:", len(fixedArray))

	// 2. Slices (dynamically-sized views into underlying arrays - used everywhere in Go)
	// Declared using []T without a length
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Initial slice:", fruits, "Len:", len(fruits), "Cap:", cap(fruits))

	// Appending elements (append returns a new slice header)
	fruits = append(fruits, "Dragonfruit", "Elderberry")
	fmt.Println("After append:", fruits, "Len:", len(fruits), "Cap:", cap(fruits))

	// Creating slices with make(type, length, capacity)
	numbers := make([]int, 3, 5) // len=3, cap=5
	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300
	fmt.Println("Made slice:", numbers, "Len:", len(numbers), "Cap:", cap(numbers))

	// Slicing operator: slice[start:end] (end index is exclusive)
	subset := fruits[1:3]
	fmt.Println("Subset [1:3]:", subset)

	// 3. Maps (hash tables / key-value stores)
	// Declaring and initializing a map
	scores := map[string]int{
		"Alice": 95,
		"Bob":   88,
	}

	// Adding/updating keys
	scores["Charlie"] = 92

	// Checking if a key exists using the comma-ok idiom
	score, ok := scores["David"]
	if ok {
		fmt.Println("David's score:", score)
	} else {
		fmt.Println("David is not in the scores map!")
	}

	// Deleting a key
	delete(scores, "Bob")

	// Iterating over a map (maps are unordered in Go)
	fmt.Println("Current Scores:")
	for student, pts := range scores {
		fmt.Printf("  %s: %d\n", student, pts)
	}
}
