package main

import "fmt"

// func printslice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printslice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printslice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printstrslice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T any] struct {
	elements []T
}

func main() {

	myStack := stack[string]{
		elements: []string{"hello","world"},
	}

	fmt.Println(myStack)

	// nums := []int{1,2,3}

	// nums:= []string{"hello","world"}

	// printslice(nums)
}