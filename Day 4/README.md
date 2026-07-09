# 🚀 GoLang Journey - Day 04

## 📅 Date

(Write today's date here)

---

# 🎯 Goal of the Day

Today's goal was to understand Go's philosophy of error handling.

Instead of learning syntax alone, I focused on understanding **why Go avoids exceptions**, how **errors are represented**, how **resources are cleaned up**, and when a program should **panic** versus **return an error**.

---

# 📚 Topics Covered

## Error Handling Philosophy

Learned why Go does not use Java-style exceptions.

Compared:

- Java → try/catch
- Go → explicit error values

Understood that Go prefers **explicit control flow** instead of hidden exception handling.

---

## The `error` Interface

Learned that:

```go
type error interface {
    Error() string
}
```

The built-in `error` type is an interface.

Any type implementing:

```go
Error() string
```

can be treated as an error.

---

## Creating Errors

Used:

```go
errors.New("Something went wrong")
```

Learned that:

- `errors.New()` returns an `error`
- it does **not** return a string

---

## `if err != nil`

Understood Go's most common idiom.

```go
if err != nil {
    return err
}
```

Learned:

- `nil` means "no value"
- errors should be checked immediately
- early returns reduce nesting
- guard clauses improve readability

---

## `defer`

Learned that:

```go
defer file.Close()
```

does not execute immediately.

Instead, Go schedules the function to run when the surrounding function exits.

---

## Deferred Stack

Learned that multiple deferred calls execute in **LIFO (Last In First Out)** order.

Example:

```go
defer fmt.Println("One")
defer fmt.Println("Two")
defer fmt.Println("Three")
```

Output:

```
Three
Two
One
```

---

## Argument Evaluation in `defer`

Learned an important concept.

```go
x := 10

defer fmt.Println(x)

x = 20
```

Output:

```
10
```

Because deferred function arguments are evaluated immediately.

Also compared it with:

```go
defer func() {
    fmt.Println(x)
}()
```

which prints:

```
20
```

because closures capture the variable, not the value.

---

## `panic`

Learned that `panic` is **not** Go's normal error handling mechanism.

Use `panic` only when the program reaches an unrecoverable state.

Examples:

- Missing configuration
- Impossible program state
- Internal programmer bugs

Avoid using `panic` for:

- Invalid password
- Validation failures
- HTTP 404
- Database query failures

Those should return an `error`.

---

## Stack Unwinding

Learned how `panic` unwinds the function call stack until the program crashes or a `recover()` handles it.

Also learned that deferred functions execute during stack unwinding.

---

## `recover()`

Learned how Go can recover from a panic.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println(r)
    }
}()
```

Also understood why `recover()` only works inside deferred functions.

---

# 🧠 Internal Working

Learned that Go follows this execution model:

Operation

↓

Return Result + Error

↓

Check Error

↓

Continue

or

↓

Return Error

For panic:

panic()

↓

Stack Unwinding

↓

Execute Deferred Calls

↓

recover() (optional)

↓

Crash or Continue

---

# 💻 Programs Written

### Creating Errors

```go
err := errors.New("Database connection failed")
```

---

### Error Checking

```go
if err != nil {
    return err
}
```

---

### defer

```go
defer file.Close()
```

---

### panic

```go
panic("Unexpected Failure")
```

---

### recover

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println(r)
    }
}()
```

---

# 🧪 Experiments

### Experiment 1

Printed:

```go
fmt.Println(err)
```

and

```go
fmt.Println(err.Error())
```

Observed identical output.

Learned that `fmt.Println()` internally uses the `Error()` method.

---

### Experiment 2

Compared deferred argument evaluation with deferred closures.

Observed different outputs.

---

### Experiment 3

Registered multiple deferred functions.

Observed reverse-order execution.

---

### Experiment 4

Triggered a panic and recovered from it.

Observed that the application continued instead of crashing.

---

# ⚠️ Common Mistakes I Learned

- Returning the wrong type from a function
- Thinking `error` is a string
- Forgetting to check `err`
- Calling `recover()` outside a deferred function
- Using `panic` for normal business logic
- Assuming deferred functions execute immediately
- Forgetting that deferred arguments are evaluated immediately

---

# 🎯 Interview Notes

### Why doesn't Go use exceptions?

Go prefers explicit error handling through return values to avoid hidden control flow.

---

### What is the `error` type?

An interface with one method:

```go
Error() string
```

---

### Why use `if err != nil` immediately?

To fail fast, avoid invalid state, and keep the happy path easy to read.

---

### Why does `recover()` only work inside `defer`?

Because deferred functions run during stack unwinding, which is when an active panic exists.

---

### When should `panic` be used?

Only for unrecoverable situations where the application cannot safely continue.

---

# 💡 Biggest Takeaways

Today I learned that Go treats errors as normal values instead of exceptions.

The Go philosophy is:

Expected failures

↓

Return `error`

Unexpected failures

↓

`panic`

Cleanup

↓

`defer`

Recovery (optional)

↓

`recover()`

This makes the execution flow explicit and predictable.

---

# 📈 Progress

Completed:

- ✅ Go Error Philosophy
- ✅ error Interface
- ✅ errors.New()
- ✅ if err != nil
- ✅ Guard Clauses
- ✅ defer
- ✅ Deferred Stack
- ✅ Argument Evaluation
- ✅ panic
- ✅ Stack Unwinding
- ✅ recover

---

# 🔥 Looking Ahead

Next Topics:

- Arrays
- Slices (Masterclass Begins)
- Slice Header
- append()
- len vs cap
- Underlying Arrays
- Memory Layout
- Production Slice Bugs

---

# 💭 Reflection

Today completely changed how I think about error handling.

Instead of relying on exceptions, I learned why Go makes failures explicit through returned error values.

Understanding `error`, `defer`, `panic`, and `recover` helped me see how Go balances simplicity, predictability, and reliability in production systems.

This was my first deep dive into Go's philosophy rather than just its syntax.
