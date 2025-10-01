# API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
Protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## Endpoints

### 1. Health Check
Check if the API is running.

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "ok"
}
```

---

### 2. User Registration
Register a new user account.

**Endpoint:** `POST /api/auth/register`

**Request Body:**
```json
{
  "email": "john.doe@example.com",
  "username": "johndoe",
  "password": "SecurePass123!",
  "first_name": "John",
  "last_name": "Doe",
  "interests": [1, 2, 3]
}
```

**Field Validations:**
- `email`: Required, valid email format
- `username`: Required, 3-100 characters, alphanumeric with underscore/hyphen
- `password`: Required, minimum 8 characters
- `first_name`: Optional
- `last_name`: Optional
- `interests`: Optional array of interest group IDs

**Success Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG4uZG9lQGV4YW1wbGUuY29tIiwiZXhwIjoxNzA5NDI4ODAwLCJpYXQiOjE3MDkzNDI0MDAsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiam9obmRvZSJ9.abc123...",
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "first_name": "John",
      "last_name": "Doe",
      "is_verified": false,
      "is_active": true,
      "created_at": "2025-10-02T01:43:34Z",
      "updated_at": "2025-10-02T01:43:34Z"
    }
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "user with this email already exists"
}
```

---

### 3. User Login
Authenticate a user and receive a JWT token.

**Endpoint:** `POST /api/auth/login`

**Request Body:**
```json
{
  "email": "john.doe@example.com",
  "password": "SecurePass123!"
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "first_name": "John",
      "last_name": "Doe",
      "is_verified": false,
      "is_active": true,
      "created_at": "2025-10-02T01:43:34Z",
      "updated_at": "2025-10-02T01:43:34Z"
    }
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": "invalid email or password"
}
```

---

### 4. Request Password Reset
Request a password reset token to be sent via email.

**Endpoint:** `POST /api/auth/password-reset/request`

**Request Body:**
```json
{
  "email": "john.doe@example.com"
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "message": "If the email exists, a password reset link has been sent"
  }
}
```

**Notes:**
- Always returns success to prevent email enumeration attacks
- Token expires in 1 hour
- In development mode (no SMTP configured), token is printed to console

---

### 5. Confirm Password Reset
Reset password using the token received via email.

**Endpoint:** `POST /api/auth/password-reset/confirm`

**Request Body:**
```json
{
  "token": "abc123def456789...",
  "new_password": "NewSecurePass123!"
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "message": "Password has been reset successfully"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "invalid or expired token"
}
```

---

### 6. Get User Profile (Protected)
Get the authenticated user's profile information.

**Endpoint:** `GET /api/auth/profile`

**Headers:**
```
Authorization: Bearer <your-jwt-token>
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "user_id": 1,
    "email": "john.doe@example.com",
    "username": "johndoe",
    "exp": 1709428800,
    "iat": 1709342400
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": "missing authorization header"
}
```

---

## Interest Groups

The following interest groups are pre-populated in the database:

| ID | Name | Description |
|----|------|-------------|
| 1 | Coworking | Connect with professionals and digital nomads |
| 2 | Photography | Share and discuss photography |
| 3 | Food | Discover local cuisine and restaurants |
| 4 | Languages | Practice and learn new languages |

---

## Error Codes

| Status Code | Description |
|-------------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Missing or invalid token |
| 404 | Not Found |
| 500 | Internal Server Error |

---

## Example Usage

### cURL Examples

**Register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "interests": [1, 2]
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Get Profile:**
```bash
curl -X GET http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Request Password Reset:**
```bash
curl -X POST http://localhost:8080/api/auth/password-reset/request \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

**Confirm Password Reset:**
```bash
curl -X POST http://localhost:8080/api/auth/password-reset/confirm \
  -H "Content-Type: application/json" \
  -d '{
    "token": "YOUR_RESET_TOKEN",
    "new_password": "newpassword123"
  }'
```

---

## Postman Collection

You can import these endpoints into Postman using the following structure:

1. Create a new collection called "Social App API"
2. Add environment variables:
   - `base_url`: `http://localhost:8080`
   - `token`: (will be set after login)
3. Add the endpoints listed above
4. For protected routes, add `Authorization: Bearer {{token}}` header

---

## Rate Limiting

Currently, there is no rate limiting implemented. Consider adding rate limiting middleware for production use.

---

## CORS

CORS is enabled for all origins (`*`). In production, configure this to allow only your frontend domains.
