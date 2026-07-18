# 🚀 GoLang Journey - Day 14 & Day 15

---

# 🎯 Goal of the Day

Today's goal was to transform the Student API from an in-memory application into a production-style backend powered by SQLite.

I learned how backend applications communicate with databases, perform complete CRUD operations, validate incoming requests, organize code into clean layers, and build REST APIs that follow professional backend engineering practices.

By the end of these lessons, my Student API became a complete CRUD application with persistent storage, standardized responses, proper validation, and a much cleaner architecture.

---

# 📚 Topics Covered

## SQLite Integration

Connected the Go application with SQLite and replaced in-memory slice storage with persistent database storage.

Learned:

- Opening database connections
- Creating tables
- Executing SQL statements
- Persisting data

---

## SQL Fundamentals

Worked with the four fundamental SQL operations.

### INSERT

Learned:

- INSERT INTO
- DB.Exec()
- LastInsertId()

---

### SELECT

Learned:

- Query()
- QueryRow()
- rows.Next()
- Scan()
- Reading multiple rows
- Reading a single row

---

### UPDATE

Learned:

- UPDATE statement
- WHERE clause
- RowsAffected()
- Detecting missing records

---

### DELETE

Learned:

- DELETE statement
- Safe deletion using WHERE
- Returning appropriate responses when records do not exist

---

## Database Package

Created reusable database functions.

Implemented:

- InsertStudent()
- GetStudents()
- GetStudentByID()
- UpdateStudent()
- DeleteStudent()

Learned why SQL should remain inside the database package instead of handlers.

---

## CRUD API

Completed the Student API by implementing all CRUD operations.

Endpoints:

```

POST /students

GET /students

GET /students/{id}

PUT /students/{id}

DELETE /students/{id}

```

Learned how every HTTP method maps to a specific database operation.

---

## Request Lifecycle

Understood how an HTTP request travels through the application.

```

Client

↓

Router

↓

Middleware

↓

Handler

↓

Database Package

↓

SQLite

↓

Database Package

↓

Handler

↓

JSON Response

↓

Client

```

This was one of the biggest architecture lessons so far.

---

## Validation

Added request validation before performing database operations.

Validated:

- Empty names
- Invalid ages
- Invalid JSON
- Invalid IDs

Learned that validation should happen before reaching the database.

---

## Validation vs Business Logic

Understood an important backend engineering concept.

Validation answers:

> Is the request structurally correct?

Business Logic answers:

> Should the system allow this request?

This distinction made backend design much clearer.

---

## Professional API Responses

Created reusable response helper functions.

Implemented:

```go
WriteJSON()

WriteError()
```

Instead of repeating response logic in every handler, all responses now follow a consistent structure.

---

## Constants

Created a dedicated constants package.

Moved all repeated strings into constants.

Examples:

- Success messages
- Error messages
- Validation messages

Learned why production applications avoid magic strings.

---

## Project Architecture

Current project structure:

```
student-api/

database/

handlers/

middleware/

models/

response/

constants/

main.go

student.db
```

Learned that every package should have one clear responsibility.

---

# 💻 Programs Written

- SQLite Connection
- Database Initialization
- Table Creation
- INSERT Operations
- SELECT Operations
- UPDATE Operations
- DELETE Operations
- CRUD Database Functions
- Request Validation
- Response Helpers
- Constants Package
- Complete Student CRUD API

---

# 🧠 Important Concepts Learned

- SQLite
- SQL
- INSERT
- SELECT
- UPDATE
- DELETE
- Exec()
- Query()
- QueryRow()
- Scan()
- RowsAffected()
- LastInsertId()
- Validation
- Business Logic
- CRUD APIs
- Request Lifecycle
- Response Helpers
- Constants
- Separation of Concerns

---

# ⚠️ Common Mistakes I Learned

- Forgetting the WHERE clause during UPDATE or DELETE.
- Writing SQL directly inside handlers.
- Allowing invalid requests to reach the database.
- Repeating response code across multiple handlers.
- Hardcoding strings instead of using constants.
- Confusing validation with business logic.
- Ignoring RowsAffected() after UPDATE and DELETE operations.

---

# 🎯 Interview Notes

## Difference Between Query() and QueryRow()

Query() returns multiple rows.

QueryRow() returns exactly one row.

---

## Why Use Exec()?

Exec() is used for SQL statements that do not return rows.

Examples:

- INSERT
- UPDATE
- DELETE

---

## Why Use RowsAffected()?

To determine whether UPDATE or DELETE actually modified any records.

---

## Why Validate Before Database Operations?

To prevent invalid data from reaching persistent storage and to fail fast.

---

## Difference Between Validation and Business Logic?

Validation checks whether input is structurally correct.

Business Logic determines whether the requested operation is allowed according to application rules.

---

## Why Separate Database Logic From Handlers?

To keep responsibilities isolated, improve maintainability, and make the codebase easier to scale.

---

# 💡 Biggest Takeaways

Today I learned that backend development is much more than connecting APIs to a database.

Professional applications separate responsibilities into different layers.

Handlers should process HTTP requests.

Database packages should execute SQL.

Validation should protect the database.

Response helpers should standardize API responses.

Constants should eliminate duplicated strings.

These small architectural decisions make applications significantly easier to maintain and scale.

---

# 📈 Progress

Completed:

- ✅ SQLite Integration
- ✅ SQL Fundamentals
- ✅ INSERT
- ✅ SELECT
- ✅ UPDATE
- ✅ DELETE
- ✅ CRUD API
- ✅ Database Package
- ✅ Request Validation
- ✅ Validation vs Business Logic
- ✅ Response Helpers
- ✅ Constants
- ✅ Professional Project Structure

---

# 🔥 Looking Ahead

Next Steps:

- Configuration Management
- Environment Variables
- Logging
- UUIDs
- Graceful Shutdown
- PostgreSQL
- Authentication (JWT)
- Docker
- Redis
- Gin Framework

---

# 💭 Reflection

These two lessons completely changed how I think about backend engineering.

Before, I only knew how to create HTTP handlers.

Now I understand how requests travel through an application, how data flows between handlers and databases, why validation protects persistent storage, and how clean architecture keeps large projects maintainable.

Building a complete CRUD API with SQLite made the Student API feel like a real backend service rather than a simple learning project.

More importantly, I realized that writing production-ready backend code is not about adding more code—it's about organizing responsibilities correctly.

This milestone marks the transition from learning Go syntax to learning backend engineering.
