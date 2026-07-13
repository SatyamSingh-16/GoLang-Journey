# 🚀 GoLang Journey - Day 08

## 📅 Date

(Write today's date here)

---

# 🎯 Goal of the Day

Today's goal was to move from learning Go as a language to understanding how professional Go applications are organized.

Instead of writing single-file programs, I learned how Go projects are structured using packages, modules, dependencies, and multiple files.

I also built my first command-line application and learned how to read from and write to files.

---

# 📚 Topics Covered

## 1. Packages

Learned that a package is a collection of related Go files.

Example:

```go
package math
```

A package groups related functionality together.

---

## package main

Learned that:

```go
package main
```

is a special package.

Execution always starts from:

```go
func main()
```

Only `package main` creates an executable program.

---

## Importing Packages

Learned how packages are imported.

Example:

```go
import "fmt"
```

Realized that `fmt` is simply another Go package written by Go developers.

---

## Multiple Files in One Package

Created multiple Go files inside the same package.

Learned that files inside one package automatically know about each other.

Example:

```
main.go

math.go
```

Both:

```go
package main
```

No imports required.

---

## Go Compiles Packages

One of today's biggest lessons.

Go compiles **packages**, not individual files.

Example:

```bash
go run .
```

compiles every Go file inside the package.

---

## Creating My Own Package

Created my own package.

Example:

```go
package math
```

Learned how another package imports it using:

```go
import "calculator/math"
```

---

## Exported vs Unexported

Learned Go's visibility rule.

Capital Letter

```go
func Add()
```

Exported.

Visible outside the package.

Lowercase

```go
func add()
```

Visible only inside the package.

Unlike Java, Go doesn't use:

- public
- private
- protected

Instead it uses capitalization.

---

# Go Modules

Created my first module.

```bash
go mod init calculator
```

Generated:

```
go.mod
```

---

## Module

Learned that a module represents an entire Go project.

A module can contain many packages.

---

## Package vs Module

Module

```
Entire Project
```

Package

```
One logical unit inside the project
```

---

## go.mod

Learned that `go.mod` stores:

- Module name
- Go version
- Project dependencies

---

# Dependency Management

Installed external libraries using:

```bash
go get github.com/google/uuid
```

Learned that Go automatically:

- Downloads the dependency.
- Updates `go.mod`.
- Updates `go.sum`.

---

## go.sum

Learned that `go.sum` stores cryptographic checksums used to verify downloaded dependencies.

Never delete it.

Always commit it to Git.

---

## go mod tidy

Learned to clean dependency lists using:

```bash
go mod tidy
```

---

# Professional Project Structure

Learned the purpose of common folders.

## cmd/

Contains application entry points.

---

## internal/

Contains application-specific code.

---

## pkg/

Contains reusable packages.

---

Learned that these folders are conventions, not language requirements.

Small projects don't need them immediately.

---

# CLI Applications

Built my first command-line application.

Learned about:

```go
os.Args
```

Understood:

- Command-line arguments
- Argument validation
- switch-based command routing

Built a small calculator CLI.

---

# File Handling

Learned how Go works with files.

## Reading Files

```go
os.ReadFile()
```

Returns:

```go
[]byte
```

Converted using:

```go
string(data)
```

---

## Writing Files

```go
os.WriteFile()
```

Learned why files use bytes.

---

## Creating Files

```go
os.Create()
```

Used:

```go
defer file.Close()
```

to release operating system resources.

---

# Why File Operations Return Errors

Understood that files can fail because:

- File doesn't exist
- Permission denied
- Invalid path
- Disk errors

Therefore every file operation returns an `error`.

---

# Programs Written

- Package Examples
- Multi-file Package Demo
- Custom Package Demo
- Module Demo
- Dependency Demo
- CLI Calculator
- File Reader
- File Writer
- File Creator

---

# 🧠 Important Concepts Learned

- Packages
- Modules
- Imports
- go.mod
- go.sum
- Dependency Management
- Exported Identifiers
- Unexported Identifiers
- Professional Project Structure
- Command-line Arguments
- os.Args
- File Reading
- File Writing
- defer file.Close()

---

# ⚠️ Common Mistakes I Learned

- Importing files instead of packages.
- Thinking folders and modules are the same.
- Forgetting that Go compiles packages, not files.
- Running `go run main.go` instead of `go run .` for multi-file projects.
- Forgetting to close files.
- Ignoring `go.sum`.
- Overengineering small projects with unnecessary folder structures.

---

# 🎯 Interview Notes

## What is a Package?

A collection of related Go files.

---

## What is a Module?

An entire Go project identified by a `go.mod` file.

---

## Difference Between Package and Module

Package

Logical unit of code.

Module

Entire project.

---

## What is go.mod?

Stores:

- Module name
- Go version
- Dependencies

---

## What is go.sum?

Stores checksums used to verify downloaded dependencies.

---

## Difference Between Exported and Unexported

Capital letter

Exported.

Lowercase

Package-private.

---

## What does go get do?

Downloads a dependency and updates `go.mod` and `go.sum`.

---

## Why use go run . ?

Because Go compiles the entire package instead of a single file.

---

# 💡 Biggest Takeaways

Today I understood that Go organizes software using packages and modules instead of giant source files.

I learned how dependencies are managed, how projects are structured professionally, and how command-line applications work.

I also learned the basics of file handling and why proper resource management with `defer file.Close()` is important.

---

# 📈 Progress

Completed:

- ✅ Packages
- ✅ Multi-file Projects
- ✅ Custom Packages
- ✅ Modules
- ✅ Imports
- ✅ go.mod
- ✅ go.sum
- ✅ Dependencies
- ✅ CLI Applications
- ✅ File Handling

---

# 🔥 Looking Ahead

Tomorrow (Day 09):

- HTTP Server
- Routing
- HTTP Methods
- JSON
- REST APIs
- First Backend Application

---

# 💭 Reflection

Day 8 marked the transition from learning Go syntax to learning how professional Go applications are built.

I learned how Go projects are organized, how external dependencies are managed, and how to create applications that interact with the operating system through command-line arguments and file operations.

Today's lessons gave me the foundation needed before starting backend development with HTTP servers and REST APIs.
