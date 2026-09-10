# 🚀 Complete Go (Golang) Learning Roadmap & Environment

Welcome to your Go learning workspace! Your development environment is fully configured with the Go compiler, official Language Server (`gopls`), Delve interactive debugger (`dlv`), and IDE settings for auto-formatting and organizing imports on save.

---

## 🛠️ Environment Status

| Component | Status | Path / Details |
| :--- | :--- | :--- |
| **Go Compiler** | `go1.26.5 windows/amd64` | `D:\go1.26.5.windows-amd64\go\bin\go.exe` |
| **GOPATH** | Configured | `D:\go` |
| **Language Server (`gopls`)** | v0.23.0 installed | `D:\go\bin\gopls.exe` |
| **Delve Debugger (`dlv`)** | v1.27.1 installed | `D:\go\bin\dlv.exe` |
| **User PATH** | Updated | Includes `D:\go\bin` |
| **IDE Settings** | Configured | Format on save (`gofmt`), organize imports, F5 launch debugger |

---

## 🧭 Step-by-Step Learning Curriculum

Each directory is a self-contained, working lesson. Follow them in order:

### 1. [01-hello-world](file:///d:/ASL/remain-Go/01-hello-world/main.go)
- **Concepts**: `package main`, `func main()`, `import "fmt"`, standard output.
- **Run**:
  ```bash
  go run ./01-hello-world
  ```

### 2. [02-variables-and-types](file:///d:/ASL/remain-Go/02-variables-and-types/main.go)
- **Concepts**: `var` declaration, short assignment `:=`, basic types (`int`, `float64`, `string`, `bool`), zero values, explicit type conversion, `const`, and `iota`.
- **Run**:
  ```bash
  go run ./02-variables-and-types
  ```

### 3. [03-control-flow](file:///d:/ASL/remain-Go/03-control-flow/main.go)
- **Concepts**: `if` with short statement (`if x := ...; x > 0`), `switch` (no implicit fallthrough, switch on expression or condition), `for` loops (standard, while-style, `range`).
- **Run**:
  ```bash
  go run ./03-control-flow
  ```

### 4. [04-slices-and-maps](file:///d:/ASL/remain-Go/04-slices-and-maps/main.go)
- **Concepts**: Arrays vs slices, dynamic resizing with `append()`, length vs capacity (`make`), sub-slices `[low:high]`, hash maps (`map[K]V`), and comma-ok lookup idiom.
- **Run**:
  ```bash
  go run ./04-slices-and-maps
  ```

### 5. [05-functions-and-pointers](file:///d:/ASL/remain-Go/05-functions-and-pointers/main.go)
- **Concepts**: Multi-value returns `(T, error)`, variadic arguments `...T`, first-class functions & closures, pointers (`&` address-of, `*` dereference), pass-by-value vs mutating state.
- **Run**:
  ```bash
  go run ./05-functions-and-pointers
  ```

### 6. [06-structs-and-methods](file:///d:/ASL/remain-Go/06-structs-and-methods/main.go)
- **Concepts**: Struct definition, exported vs unexported fields (capitalization), constructor functions (`New...`), value receiver vs pointer receiver methods, composition via struct embedding.
- **Run**:
  ```bash
  go run ./06-structs-and-methods
  ```

### 7. [07-interfaces](file:///d:/ASL/remain-Go/07-interfaces/main.go)
- **Concepts**: Implicit interfaces (duck typing), polymorphism, `any` type (`interface{}`), type assertions, type switches.
- **Run**:
  ```bash
  go run ./07-interfaces
  ```

### 8. [08-error-handling](file:///d:/ASL/remain-Go/08-error-handling/main.go)
- **Concepts**: Idiomatic explicit error handling, sentinel errors (`errors.New`), error wrapping (`fmt.Errorf("%w")`), `errors.Is`, `defer` execution order, `panic` & `recover`.
- **Run**:
  ```bash
  go run ./08-error-handling
  ```

### 9. [09-concurrency](file:///d:/ASL/remain-Go/09-concurrency/main.go)
- **Concepts**: Goroutines (`go func()`), channels (`chan`), worker pool pattern, channel closing, `sync.WaitGroup`, `sync.Mutex` for thread safety, and non-blocking `select` with timeouts.
- **Run**:
  ```bash
  go run ./09-concurrency
  ```

### 10. [10-testing](file:///d:/ASL/remain-Go/10-testing/mathutil_test.go)
- **Concepts**: Writing unit tests (`*testing.T`), Table-Driven Testing (industry standard in Go), sub-tests (`t.Run`), and performance benchmarks (`*testing.B`).
- **Run**:
  ```bash
  # Run all unit tests
  go test -v remain-go/10-testing

  # Run benchmarks
  go test -v -bench="." remain-go/10-testing
  ```

### 11. [11-mini-project-cli](file:///d:/ASL/remain-Go/11-mini-project-cli/main.go)
- **Concepts**: Real-world CLI application combining structs, slices, JSON file persistence (`os.ReadFile`, `os.WriteFile`, `encoding/json`), and CLI flags with the `flag` package.
- **Run**:
  ```bash
  # View help & existing tasks
  go run ./11-mini-project-cli

  # Add a task
  go run ./11-mini-project-cli -add "Master Go concurrency"

  # List tasks
  go run ./11-mini-project-cli -list

  # Complete task #1
  go run ./11-mini-project-cli -done 1

  # Delete task #1
  go run ./11-mini-project-cli -del 1
  ```

---

## ⚡ Essential Go Commands Cheat Sheet

| Command | Purpose |
| :--- | :--- |
| `go run <file-or-dir>` | Compiles and executes code in memory directly |
| `go build -o <output> <package>` | Compiles package into a standalone binary `.exe` |
| `go test -v ./...` | Runs all tests recursively in the workspace |
| `go vet ./...` | Examines code and reports suspicious constructs / bugs |
| `go fmt ./...` | Formats all Go files according to official Go standard style |
| `go mod tidy` | Adds missing module requirements and cleans unused ones |
| `go doc <package>.<Symbol>` | Views documentation for any standard library function in terminal |

---

## 📚 Recommended Official Learning Resources

1. [A Tour of Go](https://go.dev/tour/) - The interactive official browser tutorial.
2. [Go by Example](https://gobyexample.com/) - Practical annotated code examples for all concepts.
3. [Effective Go](https://go.dev/doc/effective_go) - Essential guide to writing idiomatic Go code.
4. [Standard Library Documentation](https://pkg.go.dev/std) - Official documentation of all built-in packages.
