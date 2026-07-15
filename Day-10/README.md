# 🚀 GoLang Journey - Day 10

## 📅 Date

(Write today's date here)

---

# 🎯 Goal of the Day

Today's goal was to move beyond basic HTTP servers and start understanding how professional backend applications are designed.

Instead of simply writing REST APIs, I learned how backend developers organize large projects using packages, how HTTP requests carry different kinds of information, and how CRUD operations work in real-world applications.

---

# 📚 Topics Covered

## Go Packages

Learned why Go applications are divided into packages.

Instead of writing everything inside one file, responsibilities are separated into different packages.

Example:

```go
package models
```

```go
package handlers
```

```go
package routes
```

---

## Why package main is Special

Learned that Go always starts execution from:

```go
package main

func main()
```

Execution begins from `main()` inside `package main`.

---

## Imports

Learned how Go imports packages.

Example:

```go
import "student-api/models"
```

Understood that Go imports **packages**, not individual files.

---

## Exported vs Unexported Identifiers

Learned Go's visibility rule.

Exported

```go
Student
```

Visible outside the package.

Unexported

```go
student
```

Accessible only inside the same package.

Unlike Java, Go doesn't use `public` or `private`; it uses capitalization.

---

## Backend Project Architecture

Learned how professional backend applications are divided into responsibilities.

```
main
   ↓
routes
   ↓
handlers
   ↓
models
```

Understood that each package has a single responsibility.

---

## CRUD Operations

Learned the four basic operations used by almost every backend application.

- Create
- Read
- Update
- Delete

Mapped them to HTTP methods.

| CRUD   | HTTP Method |
| ------ | ----------- |
| Create | POST        |
| Read   | GET         |
| Update | PUT / PATCH |
| Delete | DELETE      |

---

## PUT

Learned how PUT replaces an entire resource.

Example:

```http
PUT /students/1
```

Replaces the complete student object.

---

## PATCH

Learned how PATCH updates only selected fields.

Example:

```http
PATCH /students/1
```

Only modifies the provided fields.

---

## DELETE

Learned how resources are deleted.

Also understood how slices are modified when removing elements.

---

## Path Parameters

Learned how URLs identify specific resources.

Example:

```http
GET /students/5
```

Extracted values using:

```go
r.URL.Path
```

Used:

- strings.Split()
- strconv.Atoi()

to convert path values into integers.

---

## Query Parameters

Learned how APIs filter resources.

Example:

```http
GET /students?age=21
```

Read query parameters using:

```go
r.URL.Query().Get("age")
```

---

## HTTP Headers

Learned that requests contain metadata.

Read request headers using:

```go
r.Header.Get("Content-Type")
```

Also explored:

- Authorization
- User-Agent
- Content-Type

---

## Response Headers

Learned that servers also send headers back to clients.

Example:

```go
w.Header().Set("Content-Type", "application/json")
```

Also learned:

```go
w.WriteHeader()
```

for sending HTTP status codes.

---

## Clean Handler Design

Learned why professional projects don't keep all logic inside one huge handler.

Instead of:

```go
StudentHandler()
```

Split responsibilities into:

```go
GetStudents()

CreateStudent()

UpdateStudent()

PatchStudent()

DeleteStudent()
```

Making code easier to read and maintain.

---

# 💻 Programs Written

- Package Introduction
- Import Examples
- Exported Identifier Demo
- Backend Architecture Demo
- PUT Handler
- PATCH Handler
- DELETE Handler
- Path Parameter Example
- Query Parameter Example
- Header Reading Example

---

# 🧠 Important Concepts Learned

- Packages
- Imports
- Exported Identifiers
- CRUD
- PUT
- PATCH
- DELETE
- Path Parameters
- Query Parameters
- HTTP Headers
- Response Headers
- Clean Backend Architecture
- Single Responsibility

---

# ⚠️ Common Mistakes I Learned

- Importing files instead of packages.
- Confusing PUT with PATCH.
- Using POST to update data.
- Treating path parameters like query parameters.
- Ignoring HTTP headers.
- Writing all backend logic inside one handler.

---

# 🎯 Interview Notes

## What is a Package?

A folder containing related Go files.

---

## Difference Between PUT and PATCH

PUT

Replaces the entire resource.

PATCH

Updates only selected fields.

---

## What are Path Parameters?

Values inside the URL used to identify a specific resource.

Example:

```
/students/5
```

---

## What are Query Parameters?

Used to filter resources.

Example:

```
/students?age=21
```

---

## What are HTTP Headers?

Metadata about a request or response.

Example:

```
Content-Type
Authorization
User-Agent
```

---

## Why Split Handlers?

To keep backend code maintainable, readable, and easier to debug.

---

# 💡 Biggest Takeaways

Today I learned how professional backend applications are organized.

I understood the difference between packages and responsibilities, learned how CRUD operations map to HTTP methods, explored path parameters, query parameters, headers, and saw how backend projects become easier to maintain by separating responsibilities into small functions and packages.

---

# 📈 Progress

Completed:

- ✅ Packages
- ✅ Imports
- ✅ Exported Identifiers
- ✅ CRUD
- ✅ PUT
- ✅ PATCH
- ✅ DELETE
- ✅ Path Parameters
- ✅ Query Parameters
- ✅ HTTP Headers
- ✅ Response Headers
- ✅ Clean Handler Design

---

# 🔥 Looking Ahead

Tomorrow (Day 11):

- Refactor Student API into packages
- Student IDs
- Better Routing
- Better CRUD
- Cleaner Backend Architecture
- More Real-World Coding

---

# 💭 Reflection

Day 10 helped me move from understanding HTTP to understanding backend architecture.

Instead of simply learning new HTTP methods, I understood how backend engineers organize projects, why packages exist, how requests carry information through paths, queries, headers, and bodies, and how clean code is achieved by dividing responsibilities into smaller functions and packages.
