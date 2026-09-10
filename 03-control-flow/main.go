package main

import "fmt"

func main() {
	fmt.Println("=== 03: Control Flow (if, switch, for) ===")

	// 1. If / Else with Short Statement
	// Go allows initializing a variable before the condition: `if init; condition`
	// The variable scope is limited to the if/else blocks.
	if score := 88; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Printf("Grade: B (Score: %d)\n", score)
	} else {
		fmt.Println("Grade: C or below")
	}

	// 2. Switch Statement (Notice: NO break statement needed! Go switches do not fall through by default)
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Printf("%s is a weekday.\n", day)
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend!\n", day)
	default:
		fmt.Println("Invalid day.")
	}

	// Switch with no condition (cleaner alternative to long if-else-if chains)
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// 3. For loop (Go's ONLY looping construct! It does everything: standard, while, infinite, range)
	fmt.Print("Standard for loop: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// "While"-style loop
	counter := 3
	fmt.Print("While-style for loop: ")
	for counter > 0 {
		fmt.Printf("%d ", counter)
		counter--
	}
	fmt.Println()

	// Loop over a collection with `for ... range`
	skills := []string{"Go", "Concurrency", "Web Services", "Microservices"}
	fmt.Println("Skills to master:")
	for index, skill := range skills {
		fmt.Printf("  [%d] %s\n", index+1, skill)
	}
}
