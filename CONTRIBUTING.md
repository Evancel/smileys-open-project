# Contributing to Social App Backend

Thank you for your interest in contributing! This document provides guidelines for contributing to the project.

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Respect differing viewpoints and experiences

## Getting Started

1. **Fork the repository**
2. **Clone your fork**
   ```bash
   git clone https://github.com/your-username/windsurf-project.git
   cd windsurf-project
   ```
3. **Set up development environment** (see QUICK_START.md)
4. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Workflow

### 1. Code Style

Follow Go best practices and conventions:

```go
// Good: Clear function names, proper error handling
func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
    if err := validateInput(req); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    // ... implementation
}

// Bad: Unclear names, ignored errors
func (s *AuthService) Reg(r *RegReq) *AuthResp {
    validateInput(r) // error ignored
    // ... implementation
}
```

### 2. Project Structure

Place new code in the appropriate layer:
- **Handlers** (`internal/handlers/`) - HTTP request/response
- **Services** (`internal/service/`) - Business logic
- **Repositories** (`internal/repository/`) - Database operations
- **Models** (`internal/models/`) - Data structures
- **Middleware** (`internal/middleware/`) - Cross-cutting concerns

### 3. Naming Conventions

- **Files**: `snake_case.go` (e.g., `user_repository.go`)
- **Types**: `PascalCase` (e.g., `UserRepository`)
- **Functions**: `PascalCase` for exported, `camelCase` for private
- **Variables**: `camelCase`
- **Constants**: `PascalCase` or `SCREAMING_SNAKE_CASE`

### 4. Error Handling

Always handle errors properly:

```go
// Good
user, err := s.userRepo.GetByEmail(email)
if err != nil {
    return nil, fmt.Errorf("failed to get user: %w", err)
}

// Bad
user, _ := s.userRepo.GetByEmail(email)
```

### 5. Comments

- Add comments for exported functions
- Explain complex logic
- Use godoc format

```go
// Register creates a new user account with the provided information.
// It validates the input, hashes the password, and generates a JWT token.
// Returns an error if the email already exists or validation fails.
func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
    // Implementation
}
```

## Testing

### Writing Tests

Create test files with `_test.go` suffix:

```go
// user_repository_test.go
package repository

import (
    "testing"
)

func TestUserRepository_Create(t *testing.T) {
    // Arrange
    repo := NewUserRepository(testDB)
    user := &models.User{
        Email:    "test@example.com",
        Username: "testuser",
    }

    // Act
    err := repo.Create(user)

    // Assert
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }
    if user.ID == 0 {
        t.Error("expected user ID to be set")
    }
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./internal/service/

# Verbose output
go test -v ./...
```

## Pull Request Process

### 1. Before Submitting

- [ ] Code follows project style guidelines
- [ ] All tests pass (`go test ./...`)
- [ ] Code builds without errors (`go build cmd/api/main.go`)
- [ ] Added tests for new features
- [ ] Updated documentation if needed
- [ ] Commit messages are clear and descriptive

### 2. Commit Messages

Use clear, descriptive commit messages:

```
Good:
- "Add user profile update endpoint"
- "Fix password reset token expiration bug"
- "Refactor authentication middleware for better error handling"

Bad:
- "Update"
- "Fix bug"
- "Changes"
```

### 3. Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
How has this been tested?

## Checklist
- [ ] Tests pass
- [ ] Code follows style guidelines
- [ ] Documentation updated
```

### 4. Review Process

1. Submit PR with clear description
2. Wait for maintainer review
3. Address feedback
4. Once approved, PR will be merged

## Feature Requests

### Proposing New Features

1. **Check existing issues** to avoid duplicates
2. **Open an issue** with:
   - Clear description of the feature
   - Use case and benefits
   - Proposed implementation (optional)
3. **Wait for discussion** before implementing
4. **Implement after approval**

### Example Feature Request

```markdown
**Feature**: Add user profile image upload

**Use Case**: Users want to upload profile pictures

**Proposed Solution**:
- Add avatar upload endpoint
- Store images in S3/cloud storage
- Update user model with avatar_url

**Benefits**:
- Better user experience
- More personalized profiles
```

## Bug Reports

### Reporting Bugs

Include:
1. **Description** - What happened?
2. **Expected behavior** - What should happen?
3. **Steps to reproduce**
4. **Environment** - OS, Go version, etc.
5. **Logs/Screenshots** - If applicable

### Example Bug Report

```markdown
**Bug**: Password reset token not expiring

**Expected**: Token should expire after 1 hour

**Steps to Reproduce**:
1. Request password reset
2. Wait 2 hours
3. Use token - still works

**Environment**:
- OS: Windows 11
- Go: 1.21
- Database: PostgreSQL 15

**Logs**:
```
[error log here]
```
```

## Code Review Guidelines

### For Reviewers

- Be constructive and respectful
- Explain reasoning for requested changes
- Approve if changes look good
- Test the changes locally if possible

### For Contributors

- Respond to feedback promptly
- Ask questions if feedback is unclear
- Don't take criticism personally
- Thank reviewers for their time

## Documentation

### Updating Documentation

When adding features, update:
- **README.md** - If it affects setup or usage
- **API_DOCUMENTATION.md** - For new endpoints
- **ARCHITECTURE.md** - For architectural changes
- **Code comments** - For complex logic

### Documentation Style

- Use clear, concise language
- Include code examples
- Keep formatting consistent
- Update table of contents if needed

## Database Changes

### Adding Migrations

Add new migrations to `internal/database/migrations.go`:

```go
migrations := []string{
    // ... existing migrations
    `CREATE TABLE new_table (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`,
}
```

### Migration Guidelines

- Never modify existing migrations
- Always add new migrations at the end
- Test migrations on clean database
- Document breaking changes

## Security

### Reporting Security Issues

**DO NOT** open public issues for security vulnerabilities.

Instead:
1. Email security concerns privately
2. Include detailed description
3. Wait for response before disclosure

### Security Best Practices

- Never commit secrets or credentials
- Use environment variables for sensitive data
- Validate all user input
- Use parameterized queries
- Keep dependencies updated

## Performance

### Performance Considerations

- Avoid N+1 queries
- Use database indexes appropriately
- Profile code for bottlenecks
- Use goroutines for async operations
- Implement caching where appropriate

### Example: Optimizing Queries

```go
// Bad: N+1 query problem
for _, user := range users {
    interests, _ := repo.GetUserInterests(user.ID)
    // Process interests
}

// Good: Single query with JOIN
usersWithInterests, _ := repo.GetUsersWithInterests()
```

## Dependencies

### Adding Dependencies

1. **Evaluate necessity** - Is it really needed?
2. **Check license** - Compatible with project?
3. **Check maintenance** - Actively maintained?
4. **Add to go.mod**
   ```bash
   go get github.com/package/name
   ```
5. **Update documentation** if it affects setup

### Updating Dependencies

```bash
# Update all dependencies
go get -u ./...

# Update specific dependency
go get -u github.com/package/name

# Tidy up
go mod tidy
```

## Release Process

### Version Numbering

Follow Semantic Versioning (SemVer):
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes

### Creating a Release

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create git tag
   ```bash
   git tag -a v1.2.0 -m "Release version 1.2.0"
   git push origin v1.2.0
   ```
4. Create GitHub release with notes

## Questions?

- Open an issue for questions
- Check existing documentation
- Ask in discussions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing! 🎉
