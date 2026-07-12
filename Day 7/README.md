# 🚀 GoLang Journey - Day 07

---

# 🎯 Goal of the Day

Today's goal was to master Go's advanced type system by learning how interfaces, type assertions, type switches, and generics work together.

This day helped me understand not only **how** to write generic and flexible Go code, but also **why** these language features were designed the way they are.

---

# 📚 Topics Covered

## 1. Interfaces Deep Dive

Revisited interfaces from a compiler's perspective.

Learned:

- Interface Satisfaction
- Interface Variables
- Concrete Type
- Concrete Value
- Method Sets
- Pointer Receiver vs Value Receiver
- Nil Interfaces

One of the biggest realizations:

Interfaces don't care **what a type is**.

They care **what a type can do.**

---

## Interface Memory Model

Learned that an interface stores:

- Concrete Type
- Concrete Value

Example

```
Type : Rectangle

Value : Rectangle{}
```

This explained why:

```go
executeArea(rect)
```

works even though the function accepts a `Shape`.

The compiler automatically converts the concrete type into an interface value because it satisfies the interface.

---

## Duck Typing

Learned Go's philosophy.

> If a type has the required methods,
> it automatically satisfies the interface.

No

```java
implements
```

keyword is required.

---

## Type Assertions

Learned how to retrieve the original concrete value stored inside an interface.

Syntax:

```go
value.(int)
```

Also learned the safe version:

```go
number, ok := value.(int)
```

The comma-ok idiom prevents runtime panics.

---

## Type Switches

Learned how to determine the concrete type stored inside an interface.

Syntax:

```go
switch v := value.(type) {

case int:

case string:

case bool:

}
```

Unlike normal switch statements, Type Switches compare **types**, not values.

---

## Empty Interface (`any`)

Learned that:

```go
any
```

is simply an alias for

```go
interface{}
```

Because it has zero methods, every type automatically satisfies it.

---

## Why `any` Isn't Enough

Although `any` can store every type,

the compiler loses knowledge about the original type.

This prevents operations like:

```go
a + b
```

without performing type assertions.

---

# Generics

One of the biggest topics learned so far.

Basic Generic Function

```go
func Print[T any](value T)
```

Learned that:

- `T` represents a type.
- `any` means any type is allowed.
- The compiler preserves the original type.

Unlike:

```go
func Print(value any)
```

Generic functions maintain type safety.

---

# Identity Function

```go
func Identity[T any](value T) T
```

Learned that the return type is exactly the same as the input type.

Example:

```
Input  : int

Output : int
```

instead of returning `any`.

---

# Generic Constraints

Learned why this does NOT compile:

```go
func Add[T any](a,b T) T
```

because not every type supports `+`.

---

## Built-in Constraint

```go
comparable
```

Used for generic functions that require:

```
==
!=
```

Example:

```go
func Equal[T comparable](a,b T) bool
```

---

# Custom Constraints

Created custom interfaces representing groups of types.

Example:

```go
type Number interface {
    int | float64
}
```

Used in:

```go
func Add[T Number](a,b T) T
```

---

# Union Types

Learned that:

```go
int | float64
```

means

```
int

OR

float64
```

---

# Underlying Types

One of today's most advanced concepts.

Learned:

```go
~int
```

means

> Any type whose underlying type is int.

Example:

```go
type Age int

type Marks int
```

Both satisfy:

```go
~int
```

---

# Difference Between Interfaces and Constraints

Behavior Interface

```go
type Animal interface{
    Sound()
}
```

describes behavior.

Constraint Interface

```go
type Number interface{
    ~int | ~float64
}
```

describes a set of allowed types.

Same keyword.

Different purpose.

---

# Programs Written

- interface_intro.go
- pointer_interface.go
- nil_interface.go
- interface_design.go
- any_intro.go
- type_assertion.go
- type_switch.go
- generics_intro.go
- constraints.go
- generic_add.go

---

# 🧠 Important Concepts Learned

- Interface Satisfaction
- Duck Typing
- Interface Memory Model
- Concrete Type vs Concrete Value
- Type Assertions
- Safe Assertions
- Type Switches
- Empty Interface (`any`)
- Generic Functions
- Generic Type Parameters
- Generic Constraints
- `comparable`
- Custom Constraints
- Union Types (`|`)
- Underlying Types (`~int`)

---

# ⚠️ Common Mistakes I Learned

- Thinking interfaces require `implements`.
- Confusing an interface with the concrete value it stores.
- Using `any` when a concrete type is better.
- Using unsafe type assertions without checking.
- Thinking `switch value` and `switch value.(type)` are the same.
- Assuming Generics replace Interfaces.
- Forgetting that `~int` accepts custom types whose underlying type is `int`.

---

# 🎯 Interview Notes

## What does an interface store?

- Concrete Type
- Concrete Value

---

## What is Duck Typing?

A type satisfies an interface simply by implementing the required methods.

---

## What is `any`?

An alias for:

```go
interface{}
```

---

## Difference Between Type Assertion and Type Switch

Type Assertion

Checks one specific type.

Type Switch

Checks multiple possible types.

---

## Why were Generics introduced?

To preserve type safety while writing reusable code.

---

## What does `[T any]` mean?

`T` is a generic type parameter that can be any type.

---

## What is `comparable`?

A built-in generic constraint for types that support `==` and `!=`.

---

## What does `~int` mean?

Any type whose underlying type is `int`.

---

# 💡 Biggest Takeaways

Today I realized that Go's design is remarkably consistent.

Interfaces describe behavior.

Generics preserve type information.

Constraints reuse interfaces instead of introducing a completely new language feature.

This made Generics feel like a natural extension of the language instead of a separate system.

---

# 📈 Progress

Completed:

- ✅ Interface Deep Dive
- ✅ Interface Memory
- ✅ Duck Typing
- ✅ Type Assertions
- ✅ Type Switches
- ✅ Empty Interface (`any`)
- ✅ Generic Functions
- ✅ Generic Constraints
- ✅ `comparable`
- ✅ Custom Constraints
- ✅ Union Types
- ✅ Underlying Types (`~`)

---

# 🔥 Looking Ahead

Tomorrow begins Day 08.

Topics:

- Packages
- Modules
- Imports
- Project Structure
- Building Multi-file Applications
- Go Modules (`go.mod`)
- Creating My Own Packages

---

# 💭 Reflection

Day 7 completely changed the way I think about Go's type system.

Instead of memorizing syntax, I learned how interfaces store values, how the compiler resolves method calls, and how Generics preserve type information while remaining completely type-safe.

The biggest realization today was that Go keeps extending existing ideas rather than introducing entirely new concepts.

Interfaces power both behavior and generic constraints, making the language feel small, elegant, and internally consistent.

This was one of the most rewarding learning days of my Go journey so far.
