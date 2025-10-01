# Architecture Documentation

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Frontend Clients                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ React Native │  │   React.js   │  │   Mobile App │      │
│  │   (Expo)     │  │   (Web)      │  │              │      │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘      │
└─────────┼──────────────────┼──────────────────┼─────────────┘
          │                  │                  │
          └──────────────────┼──────────────────┘
                             │ HTTP/REST API
                             │ (JSON)
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                    Go Backend API Server                     │
│                     (Port 8080)                              │
│                                                              │
│  ┌────────────────────────────────────────────────────┐    │
│  │              Middleware Layer                       │    │
│  │  ┌─────────┐  ┌─────────┐  ┌──────────────┐      │    │
│  │  │  CORS   │→ │ Logging │→ │ Auth (JWT)   │      │    │
│  │  └─────────┘  └─────────┘  └──────────────┘      │    │
│  └────────────────────────────────────────────────────┘    │
│                          ▼                                   │
│  ┌────────────────────────────────────────────────────┐    │
│  │              Handler Layer                          │    │
│  │  ┌──────────────────────────────────────┐         │    │
│  │  │      AuthHandler                      │         │    │
│  │  │  - Register()                         │         │    │
│  │  │  - Login()                            │         │    │
│  │  │  - RequestPasswordReset()             │         │    │
│  │  │  - ResetPassword()                    │         │    │
│  │  │  - GetProfile()                       │         │    │
│  │  └──────────────────────────────────────┘         │    │
│  └────────────────────────────────────────────────────┘    │
│                          ▼                                   │
│  ┌────────────────────────────────────────────────────┐    │
│  │              Service Layer                          │    │
│  │  ┌──────────────────┐  ┌──────────────────┐       │    │
│  │  │  AuthService     │  │  EmailService    │       │    │
│  │  │  - Register      │  │  - SendReset     │       │    │
│  │  │  - Login         │  │  - SendWelcome   │       │    │
│  │  │  - ResetPassword │  └──────────────────┘       │    │
│  │  │  - ValidateToken │                             │    │
│  │  └──────────────────┘                             │    │
│  └────────────────────────────────────────────────────┘    │
│                          ▼                                   │
│  ┌────────────────────────────────────────────────────┐    │
│  │            Repository Layer                         │    │
│  │  ┌──────────────────────────────────────┐         │    │
│  │  │      UserRepository                   │         │    │
│  │  │  - Create()                           │         │    │
│  │  │  - GetByEmail()                       │         │    │
│  │  │  - GetByID()                          │         │    │
│  │  │  - UpdatePassword()                   │         │    │
│  │  │  - CreatePasswordResetToken()         │         │    │
│  │  │  - GetPasswordResetToken()            │         │    │
│  │  └──────────────────────────────────────┘         │    │
│  └────────────────────────────────────────────────────┘    │
└──────────────────────────┬──────────────────────────────────┘
                           │ SQL Queries
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    PostgreSQL Database                       │
│  ┌──────────────┐  ┌─────────────────────┐                 │
│  │    users     │  │ password_reset_     │                 │
│  │              │  │      tokens         │                 │
│  └──────────────┘  └─────────────────────┘                 │
│  ┌──────────────┐  ┌─────────────────────┐                 │
│  │  interest_   │  │  user_interests     │                 │
│  │   groups     │  │                     │                 │
│  └──────────────┘  └─────────────────────┘                 │
└─────────────────────────────────────────────────────────────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  SMTP Server │
                    │  (Optional)  │
                    └──────────────┘
```

---

## Request Flow

### User Registration Flow

```
Client                Handler              Service              Repository           Database
  │                     │                    │                      │                   │
  │  POST /register     │                    │                      │                   │
  ├────────────────────>│                    │                      │                   │
  │                     │                    │                      │                   │
  │                     │  Validate Input    │                      │                   │
  │                     ├───────────────────>│                      │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Check if exists     │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  SELECT * FROM    │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │                    │  Hash Password       │                   │
  │                     │                    │  (bcrypt)            │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Create User         │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  INSERT INTO      │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │                    │  Generate JWT        │                   │
  │                     │                    │                      │                   │
  │                     │<───────────────────┤                      │                   │
  │                     │                    │                      │                   │
  │                     │  Send Welcome      │                      │                   │
  │                     │  Email (async)     │                      │                   │
  │                     │                    │                      │                   │
  │  201 Created        │                    │                      │                   │
  │  {token, user}      │                    │                      │                   │
  │<────────────────────┤                    │                      │                   │
```

### Password Reset Flow

```
Client                Handler              Service              Repository           Database
  │                     │                    │                      │                   │
  │  POST /password-    │                    │                      │                   │
  │  reset/request      │                    │                      │                   │
  ├────────────────────>│                    │                      │                   │
  │                     │                    │                      │                   │
  │                     │  Request Reset     │                      │                   │
  │                     ├───────────────────>│                      │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Get User            │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  SELECT * FROM    │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │                    │  Generate Token      │                   │
  │                     │                    │  (random 32 bytes)   │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Save Token          │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  INSERT INTO      │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │  Send Email        │                      │                   │
  │                     │  (async)           │                      │                   │
  │                     │                    │                      │                   │
  │  200 OK             │                    │                      │                   │
  │<────────────────────┤                    │                      │                   │
  │                     │                    │                      │                   │
  │                     │                    │                      │                   │
  │  POST /password-    │                    │                      │                   │
  │  reset/confirm      │                    │                      │                   │
  ├────────────────────>│                    │                      │                   │
  │                     │                    │                      │                   │
  │                     │  Reset Password    │                      │                   │
  │                     ├───────────────────>│                      │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Validate Token      │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  SELECT * FROM    │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │                    │  Hash New Password   │                   │
  │                     │                    │  (bcrypt)            │                   │
  │                     │                    │                      │                   │
  │                     │                    │  Update Password     │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  UPDATE users     │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │                    │                      │                   │
  │                     │                    │  Mark Token Used     │                   │
  │                     │                    ├─────────────────────>│                   │
  │                     │                    │                      │  UPDATE tokens    │
  │                     │                    │                      ├──────────────────>│
  │                     │                    │                      │<──────────────────┤
  │                     │                    │<─────────────────────┤                   │
  │                     │<───────────────────┤                      │                   │
  │  200 OK             │                    │                      │                   │
  │<────────────────────┤                    │                      │                   │
```

---

## Layer Responsibilities

### 1. Handler Layer (`internal/handlers/`)
**Responsibility:** HTTP request/response handling
- Parse HTTP requests
- Validate request format
- Call service layer
- Format HTTP responses
- Handle HTTP errors

**Example:** `auth_handler.go`

### 2. Service Layer (`internal/service/`)
**Responsibility:** Business logic
- Implement business rules
- Coordinate between repositories
- Handle transactions
- Generate tokens
- Send emails

**Example:** `auth_service.go`, `email_service.go`

### 3. Repository Layer (`internal/repository/`)
**Responsibility:** Data access
- Execute database queries
- Map database rows to models
- Handle database errors
- No business logic

**Example:** `user_repository.go`

### 4. Middleware Layer (`internal/middleware/`)
**Responsibility:** Cross-cutting concerns
- CORS handling
- Request logging
- Authentication
- Rate limiting (future)

**Example:** `cors.go`, `logging.go`, `auth.go`

### 5. Model Layer (`internal/models/`)
**Responsibility:** Data structures
- Define data models
- Request/Response DTOs
- Validation tags

**Example:** `user.go`

---

## Data Flow

### Authentication Flow

```
1. User submits credentials
   ↓
2. Handler validates request format
   ↓
3. Service validates business rules
   ↓
4. Repository queries database
   ↓
5. Service compares password hash
   ↓
6. Service generates JWT token
   ↓
7. Handler returns token to client
   ↓
8. Client includes token in subsequent requests
   ↓
9. Auth middleware validates token
   ↓
10. Request proceeds to protected handler
```

---

## Security Architecture

### Password Security
```
Plain Password → bcrypt.GenerateFromPassword() → Hash (stored in DB)
                 (cost factor: 10)

Login: Plain Password → bcrypt.CompareHashAndPassword() → Match/No Match
```

### JWT Token Structure
```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "user_id": 1,
    "email": "user@example.com",
    "username": "johndoe",
    "exp": 1709428800,
    "iat": 1709342400
  },
  "signature": "HMACSHA256(base64UrlEncode(header) + '.' + base64UrlEncode(payload), JWT_SECRET)"
}
```

### Password Reset Token
```
Random 32 bytes → hex.EncodeToString() → 64-character token
                                         ↓
                                    Stored in DB with:
                                    - user_id
                                    - expires_at (1 hour)
                                    - used (boolean)
```

---

## Database Schema Relationships

```
┌─────────────────────────────────────────────────────────────┐
│                         users                                │
│  ┌────────────────────────────────────────────────────┐     │
│  │ id (PK)                                             │     │
│  │ email (UNIQUE)                                      │     │
│  │ username (UNIQUE)                                   │     │
│  │ password_hash                                       │     │
│  │ first_name, last_name, bio, avatar_url             │     │
│  │ is_verified, is_active                             │     │
│  │ created_at, updated_at                             │     │
│  └────────────────────────────────────────────────────┘     │
└──────────────┬──────────────────────────┬────────────────────┘
               │                          │
               │ 1:N                      │ N:M
               │                          │
               ▼                          ▼
┌──────────────────────────┐   ┌─────────────────────────┐
│ password_reset_tokens    │   │   user_interests        │
│ ┌──────────────────────┐ │   │ ┌─────────────────────┐ │
│ │ id (PK)              │ │   │ │ user_id (FK)        │ │
│ │ user_id (FK)         │ │   │ │ interest_id (FK)    │ │
│ │ token (UNIQUE)       │ │   │ │ joined_at           │ │
│ │ expires_at           │ │   │ └─────────────────────┘ │
│ │ used                 │ │   └────────────┬────────────┘
│ │ created_at           │ │                │ N:M
│ └──────────────────────┘ │                │
└──────────────────────────┘                ▼
                              ┌─────────────────────────┐
                              │   interest_groups       │
                              │ ┌─────────────────────┐ │
                              │ │ id (PK)             │ │
                              │ │ name (UNIQUE)       │ │
                              │ │ description         │ │
                              │ │ icon_url            │ │
                              │ │ created_at          │ │
                              │ └─────────────────────┘ │
                              └─────────────────────────┘
```

---

## Configuration Management

```
Environment Variables (.env)
         ↓
config.Load() reads and validates
         ↓
Config struct created
         ↓
Passed to services via dependency injection
         ↓
Used throughout application
```

---

## Error Handling Strategy

```
Database Error
    ↓
Repository wraps error with context
    ↓
Service handles error, may retry or transform
    ↓
Handler converts to appropriate HTTP status
    ↓
Response helper formats JSON error response
    ↓
Client receives structured error
```

**Example Error Response:**
```json
{
  "success": false,
  "error": "user with this email already exists"
}
```

---

## Deployment Architecture (Recommended)

```
┌─────────────────────────────────────────────────────────────┐
│                      Load Balancer                           │
│                      (nginx/HAProxy)                         │
└────────────┬────────────────────────────┬────────────────────┘
             │                            │
             ▼                            ▼
┌─────────────────────┐      ┌─────────────────────┐
│  Go API Instance 1  │      │  Go API Instance 2  │
│  (Docker Container) │      │  (Docker Container) │
└──────────┬──────────┘      └──────────┬──────────┘
           │                            │
           └────────────┬───────────────┘
                        ▼
           ┌─────────────────────────┐
           │  PostgreSQL Database    │
           │  (with replication)     │
           └─────────────────────────┘
                        │
                        ▼
           ┌─────────────────────────┐
           │    Redis Cache          │
           │    (optional)           │
           └─────────────────────────┘
```

---

## Performance Considerations

### Database Connection Pool
```go
db.SetMaxOpenConns(25)  // Maximum open connections
db.SetMaxIdleConns(5)   // Maximum idle connections
```

### Async Operations
- Welcome emails sent asynchronously (goroutine)
- Password reset emails sent asynchronously
- Prevents blocking HTTP responses

### JWT vs Session
- JWT chosen for stateless authentication
- No server-side session storage needed
- Scales horizontally easily

---

## Testing Strategy (Recommended)

```
Unit Tests
  ├── Service Layer Tests (business logic)
  ├── Repository Layer Tests (database operations)
  └── Validator Tests (input validation)

Integration Tests
  ├── API Endpoint Tests (full request/response)
  └── Database Migration Tests

End-to-End Tests
  └── Complete user flows
```

---

## Monitoring & Logging

### Current Logging
- Request/response logging via middleware
- Timestamps and duration
- HTTP method, path, and status code

### Recommended Additions
- Structured logging (JSON format)
- Log levels (DEBUG, INFO, WARN, ERROR)
- Centralized logging (ELK stack, Datadog)
- Application metrics (Prometheus)
- Health check endpoints

---

This architecture provides a solid foundation for a scalable, maintainable social networking application backend.
