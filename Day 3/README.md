# 🚀 GoLang Journey - Day 03

## 📅 Date

(Write today's date here)

---

# 🎯 Goal of the Day

Today's objective was to make Go programs **think and make decisions**.

Instead of writing programs that execute every line sequentially, I learned how to control the program flow using conditions, logical operators, loops, and functions.

---

# 📚 Topics Covered

## Control Flow

- Why control flow exists
- Decision making in programs
- Boolean expressions
- `if`
- `else`
- `else if`

---

## Comparison Operators

- `==`
- `!=`
- `>`
- `<`
- `>=`
- `<=`

Learned that comparison operators always produce a boolean value.

---

## Logical Operators

- `&&`
- `||`
- `!`

Learned how to combine multiple conditions.

---

## Short Circuit Evaluation

Learned that Go evaluates only as much of a logical expression as necessary.

Examples:

```go
false && checkDatabase()
```

The function is never called.

```go
true || expensiveOperation()
```

The expensive function is skipped.

This improves performance.

---

## Switch

Covered:

- Basic switch
- Multiple values in one case
- Expression-less switch
- `fallthrough`
- Automatic break behavior

Compared it with Java's switch statement.

---

## Loops

Learned that Go has only one looping keyword:

```go
for
```

Covered:

- Traditional for loop
- While-style loop
- Infinite loop

Understood why Go removed the `while` keyword.

---

## Functions

Covered:

- Function declaration
- Parameters
- Return values
- Calling functions
- Function naming
- Why Go doesn't support function overloading

---

## Multiple Return Values

One of Go's unique features.

Learned how a function can return:

```go
return result, err
```

Introduced the blank identifier:

```go
_
```

for intentionally ignoring values.

---

# 🧠 Compiler's Perspective

Today I started thinking like the Go compiler.

The compiler:

- Evaluates comparison expressions first.
- Produces boolean values.
- Uses those boolean values inside `if`.
- Converts high-level `if` statements into branching instructions.
- Automatically inserts an exit after every `switch` case unless `fallthrough` is used.
- Verifies that the number of returned values matches the number of receiving variables.

---

# 🧠 Key Learnings

- Every `if` ultimately depends on a boolean value.
- Comparison operators do not modify variables; they only produce boolean results.
- Go does not support JavaScript's truthy/falsy behavior.
- Short-circuit evaluation avoids unnecessary work.
- Go's `switch` automatically breaks after a matching case.
- One `for` keyword can replace Java's `for`, `while`, and `do...while`.
- Go functions exist without classes.
- Go does not support function overloading.
- Functions can return multiple values.
- The blank identifier (`_`) explicitly ignores unwanted return values.

---

# 💻 Programs Written

### Decision Making

```go
if isLoggedIn {
	fmt.Println("Welcome Buddy!")
}
```

---

### Eligibility Checker

```go
if age >= 18 {
	fmt.Println("Eligible")
}
```

---

### Boolean Logic

```go
isLoggedIn && isVerified

isAdmin || isModerator

!isBlocked
```

---

### Switch

```go
switch day {
case 6, 7:
	fmt.Println("Weekend")
}
```

---

### For Loop

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

---

### Function

```go
func add(a, b int) int {
	return a + b
}
```

---

### Multiple Returns

```go
func calculate(a, b int) (int, int) {
	return a+b, a*b
}
```

---

# 🧪 Experiments

### Experiment 1

Forgot braces in an `if` statement.

Observed compiler error.

Learned that Go requires braces for every block.

---

### Experiment 2

Compared a string with an integer.

```go
age := "20"

if age >= 18
```

Compiler rejected it because Go does not compare different types.

---

### Experiment 3

Removed the increment from a `for` loop.

Observed an infinite loop.

---

### Experiment 4

Used multiple return values.

Ignored one value using `_`.

---

# ⚠️ Common Mistakes I Learned

- Forgetting braces after `if`.
- Using `=` instead of `==`.
- Comparing incompatible data types.
- Assuming comparisons modify variables.
- Forgetting to update the loop variable.
- Trying to overload functions.
- Receiving the wrong number of return values.

---

# 🎯 Interview Notes

### Why doesn't Go support truthy/falsy values?

Go requires explicit boolean expressions to improve readability and catch bugs during compilation.

---

### Why is there no `while` loop?

Go's `for` loop is flexible enough to express traditional `for`, `while`, and infinite loops.

---

### Why doesn't Go support function overloading?

To keep the language simple and avoid ambiguity.

---

### Why are multiple return values useful?

They enable explicit error handling without exceptions.

---

### What is short-circuit evaluation?

Logical expressions stop evaluating as soon as the final result is known.

---

# 💡 Biggest Takeaways

Today I realized that programming is not just writing statements.

A program follows this execution pattern:

Receive Input

↓

Compare Values

↓

Produce Boolean

↓

Take Decision

↓

Execute Code

↓

Return Output

This is the foundation of backend development.

---

# 📈 Progress

Completed:

- ✅ if
- ✅ else
- ✅ else if
- ✅ Comparison Operators
- ✅ Logical Operators
- ✅ Short Circuit Evaluation
- ✅ switch
- ✅ for Loop
- ✅ Functions
- ✅ Multiple Return Values

---

# 💭 Reflection

Today was one of the most important milestones in my Go journey.

I moved beyond simply storing data and started writing programs that can make decisions, repeat tasks, and organize logic into reusable functions.

The biggest mindset shift was understanding **why Go makes certain design choices**—explicit booleans, a single loop construct, automatic `switch` breaks, and multiple return values.

I'm no longer memorizing syntax. I'm beginning to understand the philosophy behind Go's design.
