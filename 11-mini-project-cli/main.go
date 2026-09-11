package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const dataFile = "tasks.json"

// Task represents a single task item.
// Note the JSON struct tags (`json:"..."`) which control serialization.
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TaskList holds the slice of tasks
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// Load reads tasks from the local JSON file
func (tl *TaskList) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			tl.Tasks = []Task{}
			return nil
		}
		return fmt.Errorf("could not read file: %w", err)
	}

	if len(file) == 0 {
		tl.Tasks = []Task{}
		return nil
	}

	return json.Unmarshal(file, tl)
}

// Save writes the tasks list to disk as formatted JSON
func (tl *TaskList) Save(filename string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return fmt.Errorf("could not serialize tasks: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}

// Add appends a new task
func (tl *TaskList) Add(title string) Task {
	nextID := 1
	for _, t := range tl.Tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tl.Tasks = append(tl.Tasks, task)
	return task
}

// Complete marks a task as finished
func (tl *TaskList) Complete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			now := time.Now()
			tl.Tasks[i].Completed = true
			tl.Tasks[i].CompletedAt = &now
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// Delete removes a task by ID
func (tl *TaskList) Delete(id int) error {
	index := -1
	for i, t := range tl.Tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}
	// Idiomatic slice removal: append elements before index and elements after index
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
	return nil
}

// Print displays all tasks in a clean terminal table
func (tl *TaskList) Print() {
	if len(tl.Tasks) == 0 {
		fmt.Println("📭 No tasks yet! Add one using: -add \"Your task\"")
		return
	}

	fmt.Println("\n========================== TASK LIST ==========================")
	fmt.Printf("%-4s | %-6s | %-32s | %-16s\n", "ID", "STATUS", "TITLE", "CREATED")
	fmt.Println("-----+--------+----------------------------------+-----------------")
	for _, t := range tl.Tasks {
		status := " [ ] "
		if t.Completed {
			status = " [✓] "
		}
		created := t.CreatedAt.Format("02 Jan 15:04")
		fmt.Printf("%-4d | %s | %-32s | %s\n", t.ID, status, t.Title, created)
	}
	fmt.Println("===============================================================")
	fmt.Println()
}

func main() {
	// Define command-line flags
	addFlag := flag.String("add", "", "Add a new task title")
	completeFlag := flag.Int("done", 0, "Mark task ID as completed")
	deleteFlag := flag.Int("del", 0, "Delete task ID")
	listFlag := flag.Bool("list", false, "List all tasks")

	flag.Parse()

	var list TaskList
	if err := list.Load(dataFile); err != nil {
		fmt.Printf("Warning: failed to load tasks: %v\n", err)
	}

	switch {
	case *addFlag != "":
		task := list.Add(*addFlag)
		if err := list.Save(dataFile); err != nil {
			fmt.Printf("Error saving: %v\n", err)
			return
		}
		fmt.Printf(" Added task #%d: %q\n", task.ID, task.Title)

	case *completeFlag > 0:
		if err := list.Complete(*completeFlag); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}
		if err := list.Save(dataFile); err != nil {
			fmt.Printf("Error saving: %v\n", err)
			return
		}
		fmt.Printf(" Marked task #%d as completed!\n", *completeFlag)

	case *deleteFlag > 0:
		if err := list.Delete(*deleteFlag); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}
		if err := list.Save(dataFile); err != nil {
			fmt.Printf("Error saving: %v\n", err)
			return
		}
		fmt.Printf("🗑️  Deleted task #%d\n", *deleteFlag)

	case *listFlag:
		list.Print()

	default:
		// Default action: show help and list
		fmt.Println("⚡ Mini CLI Task Tracker in Go")
		fmt.Println("Usage:")
		fmt.Println("  go run ./11-mini-project-cli -list")
		fmt.Println("  go run ./11-mini-project-cli -add \"Learn Go concurrency\"")
		fmt.Println("  go run ./11-mini-project-cli -done 1")
		fmt.Println("  go run ./11-mini-project-cli -del 1")
		list.Print()
	}
}
