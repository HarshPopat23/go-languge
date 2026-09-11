// Package main demonstrates File System operations and File I/O in Go.
//
// Key Concepts:
// 1. Reading & Writing full files at once with `os.ReadFile` and `os.WriteFile` (ideal for small files).
// 2. Stream-based Reading & Writing using `os.Open`, `os.Create`, `bufio.Reader`, and `bufio.Writer` (ideal for large files).
// 3. Inspecting File Metadata with `os.Stat` and `os.FileInfo` (Name, Size, Permissions, ModTime).
// 4. Directory Operations: Reading directory contents with `os.ReadDir`.
// 5. Cleanup: Deleting files with `os.Remove`.
// 6. Best Practice: Always `defer file.Close()` right after opening/creating a file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== 18: File Handling & File I/O in Go ===")

	filename := "example.txt"
	copyFilename := "example_copy.txt"

	// -------------------------------------------------------------
	// 1. Creating and Writing to a File
	// -------------------------------------------------------------
	fmt.Println("\n--- 1. Creating and Writing File ---")
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}

	// Write bytes and strings to the open file
	_, err = file.WriteString("Hello, Go File System!\n")
	if err != nil {
		file.Close()
		panic(err)
	}

	byteData := []byte("Line 2: Writing raw bytes to file.\nLine 3: File I/O in Go is fast and simple.\n")
	_, err = file.Write(byteData)
	if err != nil {
		file.Close()
		panic(err)
	}

	// Explicitly close before reopening
	file.Close()
	fmt.Printf("Successfully created and wrote to '%s'\n", filename)

	// -------------------------------------------------------------
	// 2. Inspecting File Information & Metadata (os.Stat)
	// -------------------------------------------------------------
	fmt.Println("\n--- 2. File Metadata (os.Stat) ---")
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File Name:        %s\n", info.Name())
	fmt.Printf("File Size:        %d bytes\n", info.Size())
	fmt.Printf("Permissions/Mode: %v\n", info.Mode())
	fmt.Printf("Last Modified:    %v\n", info.ModTime())
	fmt.Printf("Is Directory?     %t\n", info.IsDir())

	// -------------------------------------------------------------
	// 3. Reading Entire File (os.ReadFile)
	// -------------------------------------------------------------
	fmt.Println("\n--- 3. Reading Entire File (os.ReadFile) ---")
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Contents:\n" + string(content))

	// -------------------------------------------------------------
	// 4. Line-by-Line Reading (bufio.Scanner)
	// -------------------------------------------------------------
	fmt.Println("--- 4. Reading Line-by-Line (bufio.Scanner) ---")
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// -------------------------------------------------------------
	// 5. Copying a File Stream (io.Copy)
	// -------------------------------------------------------------
	fmt.Println("\n--- 5. Streaming File Copy (io.Copy) ---")
	srcFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(copyFilename)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Copied %d bytes from '%s' to '%s'\n", bytesCopied, filename, copyFilename)

	// -------------------------------------------------------------
	// 6. Reading Directory Entries (os.ReadDir)
	// -------------------------------------------------------------
	fmt.Println("\n--- 6. Listing Directory Contents (os.ReadDir) ---")
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("Directory files:")
	for _, entry := range entries {
		fmt.Printf("  - %s (Dir? %t)\n", entry.Name(), entry.IsDir())
	}

	// -------------------------------------------------------------
	// 7. Cleanup / Deleting Files (os.Remove)
	// -------------------------------------------------------------
	fmt.Println("\n--- 7. Cleaning Up Test Files ---")
	_ = os.Remove(filename)
	_ = os.Remove(copyFilename)
	fmt.Println("Cleaned up temporary test files.")
}
