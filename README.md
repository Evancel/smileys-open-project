# Social App Backend API

A Go-based REST API for a social networking application designed for foreigners to connect through shared interests like Coworking, Photography, Food, and Languages.

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Email**: SMTP

## Features

- ✅ User Registration with email validation
- ✅ User Login with JWT authentication
- ✅ Password Reset via email
- ✅ Interest Groups (Coworking, Photography, Food, Languages)
- ✅ CORS middleware for cross-origin requests
- ✅ Request logging middleware
- ✅ Secure password hashing with bcrypt

## Project Structure

```
windsurf-project/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   └── server.go            # HTTP server setup and routing
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   ├── db.go                # Database connection
│   │   └── migrations.go        # Database migrations
│   ├── handlers/
│   │   └── auth_handler.go      # Authentication HTTP handlers
│   ├── middleware/
│   │   ├── auth.go              # JWT authentication middleware
│   │   ├── cors.go              # CORS middleware
│   │   └── logging.go           # Request logging middleware
│   ├── models/
│   │   └── user.go              # Data models
│   ├── repository/
│   │   └── user_repository.go   # Database operations
│   └── service/
│       ├── auth_service.go      # Authentication business logic
│       └── email_service.go     # Email sending service
├── pkg/
│   ├── response/
│   │   └── response.go          # HTTP response helpers
│   └── validator/
│       └── validator.go         # Input validation helpers
├── .env.example                 # Environment variables template
├── go.mod                       # Go module dependencies
└── README.md                    # This file
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- SMTP server credentials (optional, for email functionality)

## Installation

1. **Clone the repository**
   ```bash
   cd windsurf-project
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```bash
   # Create database
   createdb socialapp
   
   # Or using psql
   psql -U postgres
   CREATE DATABASE socialapp;
   ```

4. **Configure environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

   The server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

#### Health Check
```http
GET /health
```

#### User Registration
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe",
  "interests": [1, 2, 3]
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "email": "user@example.com",
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

#### User Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

#### Request Password Reset
```http
POST /api/auth/password-reset/request
Content-Type: application/json

{
  "email": "user@example.com"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "If the email exists, a password reset link has been sent"
  }
}
```

#### Confirm Password Reset
```http
POST /api/auth/password-reset/confirm
Content-Type: application/json

{
  "token": "abc123def456...",
  "new_password": "newsecurepassword123"
}
```

### Protected Endpoints (Require Authentication)

#### Get User Profile
```http
GET /api/auth/profile
Authorization: Bearer <your-jwt-token>
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5432/socialapp?sslmode=disable` |
| `JWT_SECRET` | Secret key for JWT signing | `your-secret-key-change-in-production` |
| `SMTP_HOST` | SMTP server host | `smtp.gmail.com` |
| `SMTP_PORT` | SMTP server port | `587` |
| `SMTP_USER` | SMTP username/email | - |
| `SMTP_PASSWORD` | SMTP password | - |
| `FRONTEND_URL` | Frontend application URL | `http://localhost:3000` |
| `PORT` | Server port | `8080` |
| `ENVIRONMENT` | Environment (development/production) | `development` |

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    bio TEXT,
    avatar_url VARCHAR(500),
    is_verified BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Password Reset Tokens Table
```sql
CREATE TABLE password_reset_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Interest Groups Table
```sql
CREATE TABLE interest_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### User Interests Table
```sql
CREATE TABLE user_interests (
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    interest_id INTEGER NOT NULL REFERENCES interest_groups(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, interest_id)
);
```

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

### Docker Support (Optional)
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

## Security Best Practices

- ✅ Passwords are hashed using bcrypt
- ✅ JWT tokens expire after 24 hours
- ✅ Password reset tokens expire after 1 hour
- ✅ Email enumeration protection on password reset
- ✅ CORS enabled for cross-origin requests
- ⚠️ Change `JWT_SECRET` in production
- ⚠️ Use HTTPS in production
- ⚠️ Set up proper database backups

## Integration with Frontend

### React Native (Expo)
```javascript
const API_URL = 'http://localhost:8080/api';

// Register user
const register = async (userData) => {
  const response = await fetch(`${API_URL}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(userData),
  });
  return response.json();
};

// Login
const login = async (email, password) => {
  const response = await fetch(`${API_URL}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });
  return response.json();
};
```

### React.js
```javascript
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For issues and questions, please open an issue on GitHub.
