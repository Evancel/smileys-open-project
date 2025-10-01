# Quick Start Guide

## Prerequisites Check

Before starting, ensure you have:
- ✅ Go 1.21+ installed (`go version`)
- ✅ PostgreSQL installed and running
- ✅ Git (optional)

## Step-by-Step Setup

### 1. Database Setup

**Option A: Using psql command line**
```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE socialapp;

# Exit psql
\q
```

**Option B: Using createdb**
```bash
createdb -U postgres socialapp
```

### 2. Configure Environment

The `.env` file is already created. Update these values if needed:

```env
DATABASE_URL=postgres://postgres:YOUR_PASSWORD@localhost:5432/socialapp?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
```

**Important:** Replace `YOUR_PASSWORD` with your PostgreSQL password.

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run the Application

```bash
go run cmd/api/main.go
```

You should see:
```
Starting server on port 8080...
```

### 5. Test the API

**Test health endpoint:**
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{"status":"ok"}
```

**Register a new user:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User",
    "interests": [1, 2]
  }'
```

## Common Issues

### Issue: "connection refused" to database

**Solution:** Make sure PostgreSQL is running
```bash
# Windows (if installed as service)
net start postgresql-x64-14

# Or check if it's running
pg_isready
```

### Issue: "database does not exist"

**Solution:** Create the database
```bash
createdb -U postgres socialapp
```

### Issue: "password authentication failed"

**Solution:** Update `DATABASE_URL` in `.env` with correct PostgreSQL password

### Issue: Port 8080 already in use

**Solution:** Change port in `.env`
```env
PORT=3000
```

## Development Workflow

### Running in development mode
```bash
go run cmd/api/main.go
```

### Building for production
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

### Running tests
```bash
go test ./...
```

## Next Steps

1. ✅ API is running on `http://localhost:8080`
2. 📖 Read [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) for endpoint details
3. 🔧 Configure SMTP settings in `.env` for email functionality
4. 🚀 Start building your frontend with React Native or React.js

## Testing with Frontend

### React Native (Expo)
```javascript
const API_URL = 'http://localhost:8080/api';
// Or use your machine's IP for mobile testing
// const API_URL = 'http://192.168.1.100:8080/api';
```

### React.js
```javascript
const API_URL = 'http://localhost:8080/api';
```

## Database Migrations

Migrations run automatically on startup. Tables created:
- ✅ `users` - User accounts
- ✅ `password_reset_tokens` - Password reset tokens
- ✅ `interest_groups` - Interest categories (Coworking, Photography, Food, Languages)
- ✅ `user_interests` - User-to-interest relationships

## Support

If you encounter any issues:
1. Check the console logs for error messages
2. Verify PostgreSQL is running and accessible
3. Ensure all environment variables are set correctly
4. Check [README.md](./README.md) for detailed documentation
