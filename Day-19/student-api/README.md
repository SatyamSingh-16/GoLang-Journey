# 🚀 GoLang Journey - Day 19

---

# 🏗️ Theme

Today marked the transition from building backend APIs to building a secure authentication system. Instead of simply processing requests and managing data, the application began verifying user identities, protecting sensitive information, enforcing authentication rules, and preparing APIs for secure access. The focus shifted from CRUD operations to implementing the security foundations that every modern backend relies on.

---

# 🎯 Goal of the Day

Today's goal was to build the foundation of a secure authentication system for the Student API.

Instead of focusing on CRUD operations, I learned how modern backend applications identify users, securely store credentials, validate incoming requests, authenticate users during login, and prepare APIs for token-based authentication.

By the end of the day, I understood how authentication systems work internally, why passwords should never be stored directly, how bcrypt protects user credentials, how login verification works, and the complete theory behind JWT-based authentication.

---

# 📚 Topics Covered

## Authentication vs Authorization

Learned the difference between proving identity and controlling access.

Covered:

- Authentication
- Authorization
- Public Routes
- Protected Routes
- Authentication Flow
- Authorization Flow
- Real-world examples

Understood that authentication answers **"Who are you?"**, while authorization answers **"What are you allowed to do?"**

---

## Users Table

Designed a dedicated table for user authentication.

Covered:

- Users Table
- SQLite
- User Schema
- UUID-based User IDs
- Email Uniqueness
- Password Hash Storage

Created fields for:

- ID
- Name
- Email
- Password Hash

Learned why passwords should never be stored directly inside the database.

---

## User Models

Created dedicated models for authentication.

Covered:

- User Model
- RegisterRequest
- LoginRequest
- Request Models
- Database Models

Learned why HTTP request models and database models should have separate responsibilities.

---

## User Registration

Built the first authentication endpoint.

Implemented:

```
POST /register
```

Covered:

- JSON Decoding
- Request Validation
- HTTP Status Codes
- Error Handling
- User Registration Flow

Learned how backend services safely receive user registration requests.

---

## Input Validation

Learned why backend applications should never trust client input.

Validated:

- Name
- Email
- Password

Covered:

- Required Fields
- Bad Request Responses
- Validation Flow
- Defensive Programming

Understood that validation should happen before any business logic or database operations.

---

## Password Hashing

Implemented secure password hashing using bcrypt.

Covered:

- bcrypt
- Password Hashing
- One-way Hash Functions
- Hash Verification
- Password Security
- Default Cost

Learned why passwords are hashed instead of encrypted and why plaintext passwords should never be stored.

---

## Database Integration

Connected registration logic with the database layer.

Covered:

- SQL INSERT
- User Creation
- UUID Generation
- Database Layer
- Separation of Responsibilities

Learned how handlers communicate with the database while keeping business logic separate from SQL queries.

---

## Duplicate Email Validation

Prevented multiple accounts from using the same email.

Covered:

- Email Lookup
- QueryRow()
- sql.ErrNoRows
- UNIQUE Constraints
- Conflict Responses

Learned how application-level validation works together with database constraints.

---

## Login System

Implemented secure user authentication.

Implemented:

```
POST /login
```

Covered:

- User Lookup
- Password Verification
- bcrypt.CompareHashAndPassword()
- Authentication Flow
- Unauthorized Responses

Learned how authentication systems verify passwords without ever decrypting them.

---

## JWT Theory

Studied the complete theory behind token-based authentication.

Covered:

- JSON Web Tokens (JWT)
- Tokens
- Header
- Payload
- Signature
- Authorization Header
- Bearer Token
- Token Expiration
- JWT Security

Understood why applications use JWT instead of sending passwords with every request.

---

# 💻 Concepts Learned

- Authentication
- Authorization
- Public Routes
- Protected Routes
- User Registration
- User Login
- Request Validation
- HTTP Status Codes
- Password Hashing
- bcrypt
- Password Verification
- UUID
- User Models
- Request Models
- Database Models
- SQL INSERT
- QueryRow()
- sql.ErrNoRows
- UNIQUE Constraints
- JWT
- Header
- Payload
- Signature
- Bearer Token
- Authorization Header

---

# 🧠 Important Concepts Learned

- Authentication verifies user identity.
- Authorization controls access to resources.
- Passwords should never be stored in plaintext.
- Hashes are one-way and cannot be reversed.
- bcrypt automatically salts passwords before hashing.
- Login verifies passwords instead of decrypting them.
- Every API should validate incoming requests.
- Request models and database models serve different purposes.
- Database constraints improve data integrity but should be complemented by application-level validation.
- JWT allows clients to authenticate without repeatedly sending passwords.
- JWT payloads are encoded, not encrypted.
- JWT signatures prevent token tampering.

---

# ⚠️ Common Mistakes I Learned

- Storing plaintext passwords.
- Encrypting passwords instead of hashing them.
- Trusting client input without validation.
- Returning different login errors for invalid email and invalid password.
- Exposing internal database errors directly to clients.
- Storing sensitive information inside JWT payloads.
- Assuming JWT payloads are encrypted.
- Reusing request models as database models.
- Ignoring duplicate email validation.
- Forgetting to hash passwords before saving them.

---

# 🎯 Interview Notes

## Difference Between Authentication and Authorization

Authentication verifies who the user is.

Authorization determines what that authenticated user is allowed to access.

---

## Why Use bcrypt Instead of Storing Passwords?

bcrypt converts passwords into one-way hashes, making it practically impossible to recover the original password if the database is compromised.

---

## Why Hash Passwords Instead of Encrypting Them?

Passwords only need verification, not recovery.

Hashing provides one-way security, whereas encryption is reversible.

---

## Why Use JWT?

JWT allows authenticated users to access protected APIs without sending their passwords with every request.

---

## What Are the Three Parts of a JWT?

A JWT consists of:

- Header
- Payload
- Signature

---

## Why Is JWT Payload Not Suitable for Sensitive Data?

JWT payloads are Base64 encoded, not encrypted, meaning anyone possessing the token can decode and read the payload.

---

## Why Use UUIDs for User IDs?

UUIDs provide globally unique identifiers that are difficult to predict and suitable for distributed systems.

---

# 💡 Biggest Takeaways

Today I learned that authentication is much more than creating login and registration endpoints. Modern backend applications must securely validate users, hash passwords using bcrypt, prevent duplicate accounts, protect user credentials, and prepare APIs for secure token-based authentication.

I also learned the complete theory behind JWTs, including how tokens prove authentication, why signatures prevent tampering, and why passwords should only be transmitted during login. These concepts transformed the Student API from a simple CRUD application into the foundation of a secure backend system.

---

# 📈 Progress

Completed:

- ✅ Authentication vs Authorization
- ✅ Users Table
- ✅ User Models
- ✅ User Registration
- ✅ Input Validation
- ✅ Password Hashing (bcrypt)
- ✅ Database Integration
- ✅ Duplicate Email Validation
- ✅ Login System
- ✅ JWT Theory

---

# 🔥 Looking Ahead

Next Steps:

- JWT Generation
- JWT Authentication Middleware
- Protected Routes
- Authorization Middleware
- Role-Based Access Control (RBAC)
- Access Tokens
- Refresh Tokens
- Logout
- Password Reset
- PostgreSQL
- Docker
- Redis
- Gin Framework

---

# 💭 Reflection

Day 17 marked one of the biggest transitions in my Go backend journey.

Unlike previous days, today's focus was not on building additional CRUD functionality but on understanding how real backend applications establish trust between users and servers. I learned how authentication systems securely store user credentials using bcrypt, validate registration requests, prevent duplicate accounts, verify user identities during login, and prepare applications for token-based authentication.

Perhaps the biggest lesson of the day was realizing that security is not a single feature but a collection of carefully designed engineering practices. Understanding authentication, password hashing, validation, and JWT theory fundamentally changed how I think about backend development.

This day established the security foundation that every production backend application depends upon and prepared the project for implementing JWT-based authentication and protected APIs in the next stage of the journey. 🚀💙
