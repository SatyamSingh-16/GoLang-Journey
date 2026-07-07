# 🚀 GoLang Journey - Day 02

## 📅 Date

(Write today's date here)

---

# 📚 Topics Covered

## Variables

- Why variables exist
- Variables as labels to memory
- Memory visualization
- Variable naming conventions
- Variable declaration using `var`
- Type inference
- Short variable declaration (`:=`)
- Assignment operator (`=`)
- Difference between `:=` and `=`

---

## Zero Values

- What are zero values?
- Why Go initializes variables automatically
- Zero value of `int`
- Zero value of `string`
- Zero value of `bool`
- Why zero values make programs safer

---

## Data Types

### Integer (`int`)

- Whole numbers
- Compile-time type inference
- Binary representation (Introduction)
- Why Go uses `int`

### String (`string`)

- Sequence of characters
- String literals
- String concatenation using `+`
- Difference between `"21"` and `21`

### Boolean (`bool`)

- `true` and `false`
- Why boolean exists
- Real-world applications
- Boolean naming conventions

---

# 🧠 Key Learnings

- A variable is a **named label** that refers to data stored in memory.
- Variables and string literals are completely different things.
- `:=` creates a **new variable** and assigns a value.
- `=` updates the value of an **existing variable**.
- Go performs **compile-time type inference**.
- Every Go variable always has a value, even if we don't initialize it.
- Strings are enclosed inside double quotes.
- `+` means addition for integers but concatenation for strings.
- Go has a strict type system and prevents mixing incompatible types.
- Boolean variables should be named like questions (`isAdmin`, `hasPermission`, etc.).

---

# 🧪 Experiments Performed

### Experiment 1

Removed `import "fmt"`

Observed compiler error:

```text
undefined: fmt
```

### Experiment 2

Imported `fmt` but didn't use it.

Observed compiler error:

```text
"fmt" imported and not used
```

### Experiment 3

Checked zero values.

```go
var age int
var name string
var isStudent bool
```

Observed:

```text
0

false
```

Learned that an empty string prints as a blank line.

---

# 💻 Programs Written

### Variable Declaration

```go
var age int = 21
```

---

### Type Inference

```go
var salary = 50000
```

---

### Short Variable Declaration

```go
name := "Satyam"
```

---

### Updating Variables

```go
score := 80

score = 95

score = 100
```

---

### String Concatenation

```go
firstName := "Satyam"
lastName := "Singh"

fullName := firstName + " " + lastName
```

---

### Boolean Variables

```go
isStudent := true
isAdmin := false
```

---

# ⚠️ Common Mistakes I Learned

- Forgetting quotes around strings.
- Thinking `"21"` and `21` are the same.
- Using `:=` to update an existing variable.
- Mixing integers and strings.
- Assuming variables without initialization contain garbage values.
- Forgetting that `=` updates an existing variable instead of creating a new one.

---

# 🎯 Interview Notes

### Difference between `:=` and `=`

- `:=` → Declare + Assign
- `=` → Assign to an existing variable

---

### What is a variable?

A variable is a **named reference to data stored in memory**.

---

### What is type inference?

The compiler automatically determines the variable's type from the assigned value during compilation.

---

### What is a zero value?

The default value automatically assigned by Go when a variable is declared but not explicitly initialized.

---

### Why does Go have a strict type system?

To catch bugs during compilation instead of runtime.

---

# 💡 My Biggest Learnings Today

- Think like the **compiler**, not just like a programmer.
- Every value has a **type**.
- Every variable has a **memory location**.
- Assignment first evaluates the right-hand side expression and then updates the left-hand side variable.
- Go favors explicitness and safety over hidden behavior.

---

# 📈 Progress

- ✅ Variables
- ✅ Zero Values
- ✅ Integer (`int`)
- ✅ String (`string`)
- ✅ Boolean (`bool`)

---

# 🚀 What's Next

- Constants
- Operators
- Control Flow (`if`, `else`)
- Loops
- Functions with Parameters
- Arrays
- Slices
- Maps
- Structs
- Concurrency

---

# 💭 Reflection

Today was the day I truly started understanding **how Go handles data**.

Instead of memorizing syntax, I learned:

- why variables exist,
- how Go infers types,
- how values are stored,
- why zero values make Go safer,
- and why Go's strict type system catches bugs before the program even runs.

This day strengthened my foundation for writing reliable Go programs.
