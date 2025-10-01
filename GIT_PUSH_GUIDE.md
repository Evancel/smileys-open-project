# Git Push Guide - Push to GitHub

## Option 1: Install Git and Push via Command Line

### Step 1: Install Git
Download and install Git from: https://git-scm.com/download/win

### Step 2: Open Git Bash (or PowerShell after restart)
Navigate to your project directory:
```bash
cd c:\Users\Evancelj\smileys-open-project\CascadeProjects\windsurf-project
```

### Step 3: Initialize Git Repository
```bash
git init
```

### Step 4: Configure Git (if not already done)
```bash
git config user.name "Evancel"
git config user.email "your-email@example.com"
```

### Step 5: Add All Files
```bash
git add .
```

### Step 6: Create Initial Commit
```bash
git commit -m "Initial commit: Go Social App Backend API with user registration and password reset"
```

### Step 7: Add Remote Repository
```bash
git remote add origin https://github.com/Evancel/smileys-open-project.git
```

### Step 8: Push to GitHub
```bash
# If pushing to main branch
git branch -M main
git push -u origin main

# Or if pushing to a specific branch
git checkout -b backend-api
git push -u origin backend-api
```

---

## Option 2: Using GitHub Desktop (Easiest)

### Step 1: Install GitHub Desktop
Download from: https://desktop.github.com/

### Step 2: Sign in to GitHub
Open GitHub Desktop and sign in with your GitHub account

### Step 3: Add Local Repository
1. Click "File" → "Add Local Repository"
2. Browse to: `c:\Users\Evancelj\smileys-open-project\CascadeProjects\windsurf-project`
3. Click "Add Repository"

### Step 4: Create Initial Commit
1. You'll see all files in the "Changes" tab
2. Add a commit message: "Initial commit: Go Social App Backend API"
3. Click "Commit to main"

### Step 5: Publish to GitHub
1. Click "Publish repository"
2. Select your account: Evancel
3. Repository name: smileys-open-project
4. Uncheck "Keep this code private" if you want it public
5. Click "Publish repository"

---

## Option 3: Upload via GitHub Web Interface

### Step 1: Go to Your Repository
Visit: https://github.com/Evancel/smileys-open-project

### Step 2: Create New Folder (if needed)
1. Click "Add file" → "Create new file"
2. Type: `backend-api/README.md`
3. This creates a new folder

### Step 3: Upload Files
1. Navigate to the folder
2. Click "Add file" → "Upload files"
3. Drag and drop all files from:
   `c:\Users\Evancelj\smileys-open-project\CascadeProjects\windsurf-project`
4. Add commit message: "Add Go backend API"
5. Click "Commit changes"

**Note**: This method may have file size limits and is less efficient for large projects.

---

## Option 4: Using Visual Studio Code

### Step 1: Open Project in VS Code
```bash
code c:\Users\Evancelj\smileys-open-project\CascadeProjects\windsurf-project
```

### Step 2: Initialize Repository
1. Click the Source Control icon (left sidebar)
2. Click "Initialize Repository"

### Step 3: Stage All Files
1. Click the "+" icon next to "Changes" to stage all files
2. Or stage individual files

### Step 4: Commit
1. Enter commit message: "Initial commit: Go Social App Backend API"
2. Click the checkmark icon or press Ctrl+Enter

### Step 5: Add Remote and Push
1. Click "..." menu in Source Control
2. Select "Remote" → "Add Remote"
3. Enter: `https://github.com/Evancel/smileys-open-project.git`
4. Name it: `origin`
5. Click "..." → "Push" → "Push to..."
6. Select `origin/main`

---

## Recommended: Option 1 or 2

**Option 1 (Command Line)** - Best for developers, most control
**Option 2 (GitHub Desktop)** - Easiest for beginners, visual interface

---

## After Pushing

### Verify Upload
Visit: https://github.com/Evancel/smileys-open-project

You should see:
- All source code files
- Documentation files
- README.md displayed on the repository page

### Update Repository Description
1. Go to repository settings
2. Add description: "Go backend API for social networking app with user authentication and password reset"
3. Add topics: `golang`, `rest-api`, `jwt`, `postgresql`, `backend`

### Add Repository Details
Consider adding:
- **Website**: Your deployed API URL (if deployed)
- **Topics**: golang, rest-api, authentication, postgresql, jwt, social-app
- **License**: MIT (already included in project)

---

## Troubleshooting

### "Permission denied (publickey)"
You need to set up SSH keys or use HTTPS with personal access token.

**Solution**: Use HTTPS URL and GitHub Personal Access Token
1. Go to GitHub Settings → Developer settings → Personal access tokens
2. Generate new token with `repo` scope
3. Use token as password when pushing

### "Repository not found"
Make sure the repository exists at: https://github.com/Evancel/smileys-open-project

If not, create it first:
1. Go to https://github.com/new
2. Repository name: `smileys-open-project`
3. Click "Create repository"

### Large files warning
If you get warnings about large files:
- The `main.exe` file (10MB) might trigger warnings
- It's already in `.gitignore`, but if it was committed, remove it:
  ```bash
  git rm --cached main.exe
  git commit -m "Remove binary file"
  ```

---

## Recommended Commit Message

```
Initial commit: Go Social App Backend API

Features:
- User registration with email validation
- User login with JWT authentication
- Password reset via email
- Interest groups system (Coworking, Photography, Food, Languages)
- PostgreSQL database with auto-migrations
- CORS, logging, and authentication middleware
- Comprehensive documentation
- Docker support

Tech Stack: Go 1.21, PostgreSQL, JWT, bcrypt
```

---

## Next Steps After Pushing

1. ✅ Verify all files are uploaded
2. ✅ Check README.md displays correctly
3. ✅ Add repository description and topics
4. ✅ Consider adding a LICENSE file (MIT recommended)
5. ✅ Set up GitHub Actions for CI/CD (optional)
6. ✅ Add branch protection rules (optional)

---

## Questions?

If you encounter any issues:
1. Check Git installation: `git --version`
2. Check Git configuration: `git config --list`
3. Verify remote URL: `git remote -v`
4. Check GitHub repository exists and you have access

---

**Choose the method that works best for you and follow the steps above!**
