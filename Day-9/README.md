# 🚀 GoLang Journey - Day 09

---

# 🎯 Goal of the Day

Today's goal was to begin my Backend Development journey using Go.

Instead of learning Go as just a programming language, I learned how real backend servers work, how HTTP requests travel from a client to a server, how routing works, and how to build my first REST API.

I also learned the professional architecture of backend applications and started organizing my project into packages.

---

# 📚 Topics Covered

## Backend Fundamentals

Learned what a backend server actually does.

Understood the complete journey of an HTTP request:

```

Browser

↓

Operating System

↓

Port

↓

Go Server

↓

Routing

↓

Handler

↓

Business Logic

↓

Response

↓

Browser

```

---

## HTTP Server

Built my first HTTP server using:

```go
http.ListenAndServe(":8080", nil)
```

Learned that:

- A server listens continuously for incoming requests.
- Port `8080` is commonly used during development.
- `ListenAndServe()` blocks the program and keeps the server running.

---

## Routing

Learned how Go connects URLs with functions using:

```go
http.HandleFunc("/students", StudentHandler)
```

Understood that Go internally maintains a routing table that maps:

```

URL

↓

Handler Function

```

Also understood why unknown routes return:

```
404 Not Found
```

---

## Handlers

Built multiple HTTP handlers.

Learned that a handler is simply a function responsible for processing one request.

Example:

```go
func StudentHandler(w http.ResponseWriter, r *http.Request)
```

Understood the purpose of:

- `http.ResponseWriter`
- `*http.Request`

---

## Request Lifecycle

Learned how every request follows the same lifecycle:

```

Receive Request

↓

Read Request

↓

Business Logic

↓

Prepare Response

↓

Write Response

```

---

## Request Object

Explored important fields inside `http.Request`.

Read:

```go
r.Method
```

```go
r.URL.Path
```

```go
r.Host
```

```go
r.UserAgent()
```

Learned that the Request object contains everything sent by the client.

---

## ResponseWriter

Learned that responses are sent using:

```go
fmt.Fprintln(w, ...)
```

instead of:

```go
fmt.Println(...)
```

Understood the difference between:

Terminal Output

vs

HTTP Response

---

## HTTP Methods

Learned the purpose of:

- GET
- POST

Understood that:

GET

↓

Reads data

POST

↓

Creates data

Also implemented method validation using:

```go
if r.Method != http.MethodGet {
    http.Error(...)
}
```

---

## HTTP Status Codes

Learned the meaning of:

- 200 OK
- 201 Created
- 400 Bad Request
- 404 Not Found
- 405 Method Not Allowed

Also learned when each status code should be returned.

---

## JSON

Worked with JSON for the first time.

Encoded Go structs into JSON:

```go
json.NewEncoder(w).Encode(student)
```

Decoded JSON requests:

```go
json.NewDecoder(r.Body).Decode(&student)
```

Learned why JSON acts as the common language between frontend and backend.

---

## Struct Tags

Used JSON tags:

```go
type Student struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Learned how struct tags map JSON fields to Go struct fields.

---

## Building My First REST API

Created a simple Student API.

Implemented:

```

GET /students

```

Returns all students.

```

POST /students

```

Creates a new student.

Used:

- Structs
- Slices
- JSON
- Request Decoding
- Response Encoding

---

## Validation

Implemented basic validation.

Checked:

- Empty Name
- Invalid Age
- Invalid JSON

Returned appropriate HTTP errors using:

```go
http.Error(...)
```

---

## Postman

Learned why browsers are not enough for backend testing.

Introduced to Postman.

Learned how Postman allows manual control over:

- HTTP Method
- URL
- Headers
- Request Body

Unlike browsers, Postman can send:

```

POST

PUT

PATCH

DELETE

```

requests directly.

---

## Professional Project Structure

Prepared a production-style project layout.

```

student-api/

├── main.go

├── go.mod

├── handlers/

├── models/

└── routes/

```

Understood why backend projects separate responsibilities instead of keeping everything inside one file.

---

## Go Packages (Introduction)

Learned that every folder represents a package.

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

Also understood why Go uses capitalization to control visibility across packages.

Example:

```go
Student
```

Exported.

```go
student
```

Package-private.

---

# 💻 Programs Written

- HTTP Server
- Hello Handler
- Student Handler
- GET API
- POST API
- JSON Encoder Example
- JSON Decoder Example
- Method Validation
- HTTP Status Code Handling
- Professional Project Structure

---

# 🧠 Important Concepts Learned

- Backend Architecture
- HTTP
- HTTP Server
- Routing
- Handlers
- Request Lifecycle
- ResponseWriter
- Request Object
- GET
- POST
- JSON Encoding
- JSON Decoding
- Struct Tags
- HTTP Status Codes
- REST APIs
- Postman
- Professional Folder Structure
- Go Packages

---

# ⚠️ Common Mistakes I Learned

- Using `fmt.Println()` instead of writing to `ResponseWriter`.
- Forgetting to validate HTTP methods.
- Thinking handlers are automatically called without routing.
- Forgetting to decode JSON before using request data.
- Returning incorrect HTTP status codes.
- Mixing models, handlers, and routes inside one file.
- Thinking browsers can properly test POST requests.

---

# 🎯 Interview Notes

## What is an HTTP Server?

A program that listens for incoming HTTP requests and sends responses back to clients.

---

## What is a Handler?

A function responsible for processing one HTTP request.

---

## Difference Between GET and POST

GET

Used for reading data.

POST

Used for creating data.

---

## What is JSON?

A language-independent data format used to exchange information between clients and servers.

---

## Why use ResponseWriter?

To write HTTP responses back to the client.

---

## What is http.Request?

An object containing everything sent by the client, including method, path, headers, and body.

---

## Why use Postman?

Because browsers mainly send GET requests, while Postman allows testing every HTTP method and custom request body.

---

## What is a REST API?

An API that exposes resources through URLs and uses HTTP methods like GET and POST to interact with them.

---

# 💡 Biggest Takeaways

Today I learned how backend applications actually work.

Instead of memorizing Go syntax, I understood the complete lifecycle of an HTTP request, how routing connects URLs to handlers, how requests are processed, how JSON is exchanged between client and server, and how REST APIs are built.

I also learned the importance of organizing backend applications using packages and professional project structures.

---

# 📈 Progress

Completed:

- ✅ HTTP Server
- ✅ Routing
- ✅ Handlers
- ✅ Request & Response
- ✅ GET
- ✅ POST
- ✅ JSON
- ✅ REST API
- ✅ Postman
- ✅ Status Codes
- ✅ Validation
- ✅ Professional Project Structure

---

# 🔥 Looking Ahead

Tomorrow (Day 10):

- Implement packages completely
- Connect handlers, models, and routes
- Build a professional multi-package backend
- Improve Student API architecture
- Prepare for database integration

---

# 💭 Reflection

Day 9 marked the beginning of my backend development journey.

Today I stopped thinking of Go as just a programming language and started understanding how backend systems actually communicate over HTTP. I built my first REST API, learned how requests and responses flow through a server, worked with JSON, explored HTTP methods and status codes, and prepared a professional project structure that will grow into a production-ready backend application in the coming days.
