# 🚀 GoLang Journey - Day 29

---

# 🏗️ Theme

Today marked another major milestone in my Go backend journey as I built my first complete **Create API** using **Gin** and **PostgreSQL**.

Unlike previous days where the backend only retrieved data from the database, today I learned how professional backend systems receive data from clients, validate incoming requests, transform JSON into Go objects, execute database INSERT operations, and return proper HTTP responses.

The biggest realization of the day was that creating data is much more than simply writing an INSERT query. Professional backend systems carefully separate external client requests from internal database models using DTOs, validate incoming data before it reaches the business layer, and ensure every layer performs exactly one responsibility.

By the end of the day, I had implemented my first production-style **POST /students** endpoint following clean architecture principles while gaining a much deeper understanding of request processing inside Gin.

---

# 🎯 Goal of the Day

Today's goal was to understand how production backend applications create resources through REST APIs.

The objective was to build a complete **POST /students** endpoint while learning every concept involved, including Request Bodies, DTOs, JSON Binding, Validation, Repository INSERT operations, HTTP Status Codes, and API testing using Postman.

By the end of the day, I successfully completed my first write operation using the layered architecture built over the previous days.

---

# 📚 Topics Covered

## Understanding POST Requests

Studied how HTTP POST differs from GET.

Covered:

- POST Requests
- Request Body
- Resource Creation
- HTTP Methods
- Client Requests
- REST Principles

Learned that POST requests create new resources and send data inside the HTTP request body instead of URL parameters.

---

## HTTP Request Body

Studied how HTTP requests carry client data.

Covered:

- Request Body
- JSON Payload
- HTTP Structure
- Headers
- Content-Type
- Client Data

Learned that POST requests transport structured JSON data inside the request body while the URL simply identifies the resource.

---

## Data Transfer Objects (DTO)

Created the first Request DTO.

Covered:

- DTO
- API Contracts
- Request Objects
- Data Isolation
- Client Models
- Layer Separation

Learned why production applications never bind client JSON directly into database models and instead define exactly what data is allowed to cross the API boundary.

---

## JSON Binding

Studied how Gin converts JSON into Go structs.

Covered:

- ShouldBindJSON()
- Reflection
- JSON Parsing
- Struct Population
- Pointer Semantics
- Request Binding

Learned that Gin reads the raw HTTP body, parses JSON, matches fields using reflection, and fills Go structs automatically.

---

## Struct Tags

Implemented JSON mapping.

Covered:

- json Tags
- Field Mapping
- Reflection Metadata
- External Representation
- Go Struct Tags

Learned how struct tags define how JSON field names map to Go struct fields.

---

## Validation

Implemented automatic request validation.

Covered:

- Validation Tags
- Required Fields
- Email Validation
- Range Validation
- Input Validation
- Request Verification

Learned that Gin automatically validates request data during JSON binding before the request reaches business logic.

---

## DTO Validation vs Business Validation

Studied where validation belongs.

Covered:

- API Validation
- Business Rules
- Layer Responsibilities
- Input Validation
- Service Validation

Learned the difference between validating request structure inside DTOs and enforcing business rules inside the Service layer.

---

## Repository INSERT Operations

Implemented the first database write operation.

Covered:

- INSERT Queries
- QueryRowContext()
- RETURNING
- Scan()
- PostgreSQL
- Parameterized Queries

Learned how PostgreSQL returns generated IDs using RETURNING and how Go retrieves them through QueryRowContext().

---

## QueryRowContext() vs ExecContext()

Studied the proper database APIs for write operations.

Covered:

- ExecContext()
- QueryRowContext()
- INSERT
- UPDATE
- DELETE
- RETURNING

Learned when each database method should be used depending on whether SQL returns rows.

---

## HTTP Responses

Completed the POST endpoint.

Covered:

- 201 Created
- HTTP Status Codes
- JSON Responses
- Response Body
- Error Responses

Learned why resource creation should return HTTP 201 instead of HTTP 200.

---

## API Testing

Verified the endpoint using Postman.

Covered:

- Postman
- Request Body
- JSON Requests
- Response Validation
- Database Verification
- API Testing

Learned how to test production APIs, verify inserted database records, and debug invalid requests.

---

# 💻 Concepts Learned

- POST Requests
- HTTP Request Body
- JSON Payload
- DTO (Data Transfer Object)
- Request DTO
- API Contract
- ShouldBindJSON()
- Reflection
- Struct Tags
- JSON Mapping
- Validation Tags
- Required Validation
- Email Validation
- DTO Validation
- Business Validation
- INSERT Query
- RETURNING id
- QueryRowContext()
- ExecContext()
- Scan()
- Parameterized Queries
- PostgreSQL INSERT
- HTTP 201 Created
- Postman API Testing
- Clean Request Flow

---

# 🧠 Important Concepts Learned

- POST requests send data inside the request body instead of the URL.
- DTOs define exactly what data clients are allowed to send.
- Database models should never be exposed directly to clients.
- `ShouldBindJSON()` performs both JSON binding and validation.
- Struct tags define how external JSON maps to Go structs.
- Validation belongs at the API boundary.
- Business rules belong inside the Service layer.
- PostgreSQL's `RETURNING` clause allows INSERT statements to return generated values.
- `QueryRowContext()` is used whenever SQL returns exactly one row.
- `ExecContext()` is used when SQL modifies data without returning rows.
- `201 Created` is the proper success response for resource creation.
- Every layer communicates only with its neighboring layer.

---

# ⚠️ Common Mistakes I Learned

- Binding client JSON directly into database models.
- Returning HTTP 200 instead of HTTP 201 after creating resources.
- Forgetting to validate incoming requests.
- Performing business validation inside the Repository.
- Writing SQL inside Handlers.
- Returning database errors directly to HTTP clients.
- Using `ExecContext()` when SQL returns data.
- Forgetting to pass pointers into `ShouldBindJSON()`.
- Assuming validation belongs inside the database.
- Mixing API models with persistence models.

---

# 🎯 Interview Notes

## Why Do We Need DTOs?

DTOs define the contract between clients and the backend, allowing only approved data to enter the application while hiding internal database structures.

---

## Why ShouldBindJSON() Receive A Pointer?

Gin needs to modify the original struct during JSON binding. Passing a value would only modify a copy.

---

## What Does ShouldBindJSON() Actually Do?

It reads the HTTP request body, parses JSON, maps fields using reflection, validates struct tags, and populates the target struct.

---

## Why Doesn't The Repository Validate Requests?

Repositories should only communicate with the database. Input validation belongs to the API boundary while business validation belongs to the Service layer.

---

## When Should QueryRowContext() Be Used?

Whenever SQL returns exactly one row, including PostgreSQL INSERT statements using `RETURNING`.

---

## When Should ExecContext() Be Used?

Whenever SQL modifies data without returning rows, such as INSERT, UPDATE, or DELETE statements without `RETURNING`.

---

## Why Return 201 Created?

HTTP 201 explicitly communicates that a new resource has been successfully created, making the API more expressive than simply returning HTTP 200.

---

# 🏛️ Architecture Reinforced Today

```
Client

↓

POST /students

↓

Gin Router

↓

StudentHandler

↓

ShouldBindJSON()

↓

DTO Validation

↓

StudentService

↓

StudentRepository

↓

INSERT INTO students

↓

RETURNING id

↓

PostgreSQL

↓

Repository

↓

Service

↓

Handler

↓

201 Created

↓

JSON Response

↓

Client
```

---

# 💡 Biggest Takeaways

Today completely changed how I think about creating resources in backend systems.

Initially, I believed creating an API simply meant writing an INSERT query. Today I realized that production systems involve many responsibilities before data ever reaches the database. Incoming requests must be validated, transformed into DTOs, checked against business rules, passed through multiple architectural layers, and finally persisted using safe parameterized SQL queries.

The most valuable realization was understanding that DTOs and validation are not simply conveniences—they form the protective boundary between external clients and the application's internal implementation.

I also developed a much stronger understanding of Gin's request binding, Go reflection, PostgreSQL's RETURNING clause, and proper HTTP semantics for resource creation.

---

# 📈 Progress

Completed:

- ✅ POST Requests
- ✅ HTTP Request Body
- ✅ DTO
- ✅ JSON Binding
- ✅ Struct Tags
- ✅ Validation
- ✅ DTO Validation
- ✅ Business Validation
- ✅ INSERT Queries
- ✅ RETURNING id
- ✅ QueryRowContext()
- ✅ HTTP 201 Created
- ✅ Complete POST Endpoint
- ✅ Postman Testing

---

# 🔥 Looking Ahead

Next Steps:

- PUT Requests
- PATCH Requests
- UPDATE Queries
- ExecContext()
- RowsAffected()
- Updating Existing Records
- Delete APIs
- Middleware
- Authentication
- Authorization
- Transactions
- Logging
- Production Error Handling

---

# 💭 Reflection

Day 41 has been another transformative step in my backend journey.

Until now, my application only retrieved information from PostgreSQL. Today I built my first complete write operation using a professional layered architecture. More importantly, I learned that creating data safely involves much more than inserting rows into a database.

I now understand why production applications introduce DTOs, validation, and Service layers before interacting with persistence. Each layer protects the application in a different way while keeping responsibilities isolated.

The distinction between input validation and business validation was one of the biggest lessons of the day. I also gained a much deeper appreciation for Gin's JSON binding system, Go's reflection capabilities, PostgreSQL's RETURNING clause, and HTTP semantics such as 201 Created.

With Day 41 complete, my Student API now supports both reading and creating resources using a clean production-style architecture. The next step is implementing update operations and moving even closer to a complete CRUD backend.
