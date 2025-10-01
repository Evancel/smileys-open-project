# 🚀 START HERE - Social App Backend API

Welcome! This is your complete Go backend API for a social networking application.

## 📋 Quick Navigation

| Document | Purpose |
|----------|---------|
| **[START_HERE.md](START_HERE.md)** | 👈 You are here - Overview and quick links |
| **[QUICK_START.md](QUICK_START.md)** | ⚡ Get up and running in 5 minutes |
| **[README.md](README.md)** | 📖 Complete project documentation |
| **[API_DOCUMENTATION.md](API_DOCUMENTATION.md)** | 🔌 API endpoints reference |
| **[ARCHITECTURE.md](ARCHITECTURE.md)** | 🏗️ System architecture and design |
| **[PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)** | 📊 High-level project overview |
| **[DEPLOYMENT.md](DEPLOYMENT.md)** | 🌐 Production deployment guide |
| **[CONTRIBUTING.md](CONTRIBUTING.md)** | 🤝 How to contribute |

---

## ✨ What You Have

A **production-ready** Go REST API with:

### Core Features ✅
- ✅ **User Registration** - Email validation, password hashing, interest selection
- ✅ **User Login** - JWT authentication with 24-hour tokens
- ✅ **Password Reset** - Email-based two-step process
- ✅ **Protected Routes** - JWT middleware for secure endpoints
- ✅ **Interest Groups** - Pre-populated (Coworking, Photography, Food, Languages)

### Tech Stack 🛠️
- **Backend**: Go 1.21+
- **Database**: PostgreSQL with auto-migrations
- **Auth**: JWT tokens + bcrypt password hashing
- **Email**: SMTP support (optional)
- **API**: RESTful with JSON responses

### Architecture 🏗️
```
Clean Architecture with:
├── Handlers (HTTP layer)
├── Services (Business logic)
├── Repositories (Data access)
├── Middleware (CORS, Auth, Logging)
└── Models (Data structures)
```

---

## 🎯 Choose Your Path

### Path 1: Quick Start (5 minutes)
**Goal**: Get the API running locally

1. **Ensure PostgreSQL is running**
2. **Create database**: `createdb socialapp`
3. **Run the API**: `go run cmd/api/main.go`
4. **Test it**: `curl http://localhost:8080/health`

👉 **[Full Quick Start Guide](QUICK_START.md)**

### Path 2: Learn the API (10 minutes)
**Goal**: Understand available endpoints

1. **Read API docs**: [API_DOCUMENTATION.md](API_DOCUMENTATION.md)
2. **Test endpoints**: Run `.\test_api.ps1` (PowerShell)
3. **Try with Postman**: Import endpoints from docs

### Path 3: Understand Architecture (20 minutes)
**Goal**: Learn how it's built

1. **Read architecture**: [ARCHITECTURE.md](ARCHITECTURE.md)
2. **Explore code structure**: See [Project Structure](#project-structure)
3. **Review design patterns**: Clean architecture, repository pattern

### Path 4: Deploy to Production (30 minutes)
**Goal**: Deploy to a server

1. **Choose deployment method**: Docker, VPS, or Cloud
2. **Follow deployment guide**: [DEPLOYMENT.md](DEPLOYMENT.md)
3. **Configure SSL**: Use Let's Encrypt
4. **Set up monitoring**: Health checks and logs

---

## 📁 Project Structure

```
windsurf-project/
│
├── 📄 Documentation
│   ├── START_HERE.md          ← You are here
│   ├── QUICK_START.md         ← Setup guide
│   ├── README.md              ← Full documentation
│   ├── API_DOCUMENTATION.md   ← API reference
│   ├── ARCHITECTURE.md        ← System design
│   ├── PROJECT_SUMMARY.md     ← Overview
│   ├── DEPLOYMENT.md          ← Production guide
│   └── CONTRIBUTING.md        ← Contribution guide
│
├── 🔧 Configuration
│   ├── .env                   ← Environment variables
│   ├── .env.example           ← Environment template
│   ├── go.mod                 ← Go dependencies
│   ├── Dockerfile             ← Docker image
│   ├── docker-compose.yml     ← Docker services
│   ├── Makefile               ← Build commands
│   └── .gitignore             ← Git ignore rules
│
├── 💻 Source Code
│   ├── cmd/api/main.go        ← Application entry point
│   │
│   ├── internal/              ← Private application code
│   │   ├── api/               ← HTTP server setup
│   │   ├── config/            ← Configuration
│   │   ├── database/          ← DB connection & migrations
│   │   ├── handlers/          ← HTTP handlers
│   │   ├── middleware/        ← CORS, Auth, Logging
│   │   ├── models/            ← Data models
│   │   ├── repository/        ← Database operations
│   │   └── service/           ← Business logic
│   │
│   └── pkg/                   ← Public reusable packages
│       ├── response/          ← HTTP response helpers
│       └── validator/         ← Input validation
│
└── 🧪 Testing
    └── test_api.ps1           ← API test script
```

---

## 🔌 API Endpoints at a Glance

### Public Endpoints
```
GET  /health                              Health check
POST /api/auth/register                   Register new user
POST /api/auth/login                      User login
POST /api/auth/password-reset/request     Request password reset
POST /api/auth/password-reset/confirm     Confirm password reset
```

### Protected Endpoints (Require JWT)
```
GET  /api/auth/profile                    Get user profile
```

👉 **[Full API Documentation](API_DOCUMENTATION.md)**

---

## 🚀 Quick Commands

### Development
```bash
# Run the application
go run cmd/api/main.go

# Or use Makefile
make run

# Run tests
go test ./...

# Build binary
go build -o bin/api cmd/api/main.go
```

### Docker
```bash
# Start all services (API + PostgreSQL)
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down
```

### Testing
```powershell
# Windows PowerShell
.\test_api.ps1

# Or manually
curl http://localhost:8080/health
```

---

## 🎓 Learning Resources

### Understanding the Code

1. **Start with main.go** (`cmd/api/main.go`)
   - See how the application initializes
   - Understand dependency injection

2. **Explore the API layer** (`internal/api/server.go`)
   - See route definitions
   - Understand middleware chain

3. **Review a complete flow** (`internal/handlers/auth_handler.go`)
   - Handler → Service → Repository → Database
   - See error handling patterns

4. **Study the architecture** ([ARCHITECTURE.md](ARCHITECTURE.md))
   - Request flow diagrams
   - Layer responsibilities
   - Security architecture

### Key Concepts

- **Clean Architecture**: Separation of concerns across layers
- **Repository Pattern**: Abstraction over data access
- **Dependency Injection**: Services receive dependencies
- **Middleware Chain**: CORS → Logging → Auth → Handler
- **JWT Authentication**: Stateless token-based auth

---

## 🔐 Security Features

- ✅ **Password Hashing**: bcrypt with cost factor 10
- ✅ **JWT Tokens**: 24-hour expiration
- ✅ **Reset Tokens**: 1-hour expiration, one-time use
- ✅ **Email Protection**: No user enumeration
- ✅ **Input Validation**: All endpoints validated
- ✅ **SQL Injection**: Parameterized queries

---

## 🌐 Frontend Integration

### React Native (Expo)
```javascript
const API_URL = 'http://localhost:8080/api';

const register = async (userData) => {
  const response = await fetch(`${API_URL}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(userData),
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

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
```

---

## 📊 Database Schema

### Tables Created Automatically
- **users** - User accounts and profiles
- **password_reset_tokens** - Temporary reset tokens
- **interest_groups** - Categories (4 pre-populated)
- **user_interests** - User-to-interest relationships

All migrations run automatically on startup!

---

## 🐛 Troubleshooting

### Common Issues

**"Cannot connect to database"**
```bash
# Check PostgreSQL is running
pg_isready

# Create database if missing
createdb socialapp
```

**"Port 8080 already in use"**
```bash
# Change port in .env
PORT=3000
```

**"Module not found"**
```bash
# Install dependencies
go mod download
go mod tidy
```

👉 **[Full Troubleshooting Guide](QUICK_START.md#common-issues)**

---

## 📈 Next Steps

### Immediate (Now)
1. ✅ Get the API running locally
2. ✅ Test endpoints with test script
3. ✅ Read API documentation

### Short-term (This Week)
1. 📱 Integrate with React Native/React.js frontend
2. 📧 Configure SMTP for email functionality
3. 🧪 Add unit tests for your use cases

### Long-term (This Month)
1. 🚀 Deploy to production
2. 📊 Add monitoring and logging
3. ⚡ Implement additional features

---

## 🎯 Feature Roadmap

### Implemented ✅
- User registration and authentication
- Password reset functionality
- Interest groups system
- JWT-based security
- CORS and logging middleware

### Recommended Additions 📋
- [ ] Email verification on registration
- [ ] User profile update endpoint
- [ ] Social features (posts, comments)
- [ ] Real-time messaging
- [ ] Image upload for avatars
- [ ] Search and filter users
- [ ] Rate limiting
- [ ] Refresh tokens

---

## 💡 Tips for Success

1. **Start Simple**: Get it running locally first
2. **Read the Docs**: Everything is documented
3. **Test Early**: Use the test script frequently
4. **Secure Production**: Change JWT_SECRET, use HTTPS
5. **Monitor Logs**: Watch for errors and performance issues
6. **Backup Database**: Set up automated backups
7. **Version Control**: Commit changes regularly

---

## 📞 Getting Help

### Documentation
- Check the relevant .md file for your question
- Search existing issues on GitHub

### Common Questions
- **Setup issues?** → [QUICK_START.md](QUICK_START.md)
- **API usage?** → [API_DOCUMENTATION.md](API_DOCUMENTATION.md)
- **Architecture questions?** → [ARCHITECTURE.md](ARCHITECTURE.md)
- **Deployment help?** → [DEPLOYMENT.md](DEPLOYMENT.md)

### Contributing
- Want to contribute? → [CONTRIBUTING.md](CONTRIBUTING.md)
- Found a bug? Open an issue
- Have a feature idea? Open a discussion

---

## 🎉 You're Ready!

Your Go backend API is **production-ready** and waiting for you to:

1. **Run it**: `go run cmd/api/main.go`
2. **Test it**: `.\test_api.ps1`
3. **Build with it**: Connect your frontend
4. **Deploy it**: Follow the deployment guide

**Happy coding!** 🚀

---

**Project Status**: ✅ Production Ready  
**Last Updated**: October 2, 2025  
**Go Version**: 1.21+  
**License**: MIT
