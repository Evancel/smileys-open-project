# Deployment Guide

## Table of Contents
1. [Local Development](#local-development)
2. [Docker Deployment](#docker-deployment)
3. [Production Deployment](#production-deployment)
4. [Environment Configuration](#environment-configuration)
5. [Database Setup](#database-setup)
6. [Monitoring & Maintenance](#monitoring--maintenance)

---

## Local Development

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Git

### Setup Steps

1. **Clone and navigate to project**
   ```bash
   cd windsurf-project
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```bash
   createdb socialapp
   ```

4. **Configure environment**
   ```bash
   # .env file is already created
   # Update DATABASE_URL with your PostgreSQL password
   ```

5. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

6. **Test the API**
   ```powershell
   # Windows PowerShell
   .\test_api.ps1
   
   # Or manually
   curl http://localhost:8080/health
   ```

---

## Docker Deployment

### Using Docker Compose (Recommended)

**Start all services:**
```bash
docker-compose up -d
```

This will start:
- PostgreSQL database (port 5432)
- Go API server (port 8080)

**View logs:**
```bash
docker-compose logs -f api
```

**Stop services:**
```bash
docker-compose down
```

**Rebuild after code changes:**
```bash
docker-compose up -d --build
```

### Using Docker Only

**Build image:**
```bash
docker build -t socialapp-api .
```

**Run container:**
```bash
docker run -d \
  --name socialapp-api \
  -p 8080:8080 \
  -e DATABASE_URL="postgres://user:pass@host:5432/socialapp?sslmode=disable" \
  -e JWT_SECRET="your-secret-key" \
  socialapp-api
```

---

## Production Deployment

### Option 1: Traditional VPS (DigitalOcean, AWS EC2, etc.)

#### 1. Server Setup

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Go
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install PostgreSQL
sudo apt install postgresql postgresql-contrib -y

# Install nginx (reverse proxy)
sudo apt install nginx -y
```

#### 2. Database Setup

```bash
sudo -u postgres psql
CREATE DATABASE socialapp;
CREATE USER socialapp_user WITH PASSWORD 'strong_password';
GRANT ALL PRIVILEGES ON DATABASE socialapp TO socialapp_user;
\q
```

#### 3. Deploy Application

```bash
# Clone repository
git clone <your-repo-url>
cd windsurf-project

# Build application
go build -o api cmd/api/main.go

# Create systemd service
sudo nano /etc/systemd/system/socialapp.service
```

**Service file content:**
```ini
[Unit]
Description=Social App API
After=network.target postgresql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/home/ubuntu/windsurf-project
ExecStart=/home/ubuntu/windsurf-project/api
Restart=on-failure
RestartSec=5s
Environment="DATABASE_URL=postgres://socialapp_user:strong_password@localhost:5432/socialapp?sslmode=disable"
Environment="JWT_SECRET=your-production-secret-key"
Environment="PORT=8080"
Environment="ENVIRONMENT=production"

[Install]
WantedBy=multi-user.target
```

**Enable and start service:**
```bash
sudo systemctl daemon-reload
sudo systemctl enable socialapp
sudo systemctl start socialapp
sudo systemctl status socialapp
```

#### 4. Configure Nginx

```bash
sudo nano /etc/nginx/sites-available/socialapp
```

**Nginx configuration:**
```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}
```

**Enable site:**
```bash
sudo ln -s /etc/nginx/sites-available/socialapp /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

#### 5. SSL Certificate (Let's Encrypt)

```bash
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d api.yourdomain.com
```

### Option 2: Docker on VPS

```bash
# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Clone and deploy
git clone <your-repo-url>
cd windsurf-project
docker-compose up -d
```

### Option 3: Cloud Platforms

#### Heroku

1. **Create Heroku app**
   ```bash
   heroku create your-app-name
   ```

2. **Add PostgreSQL**
   ```bash
   heroku addons:create heroku-postgresql:mini
   ```

3. **Set environment variables**
   ```bash
   heroku config:set JWT_SECRET=your-secret-key
   heroku config:set ENVIRONMENT=production
   ```

4. **Create Procfile**
   ```
   web: ./api
   ```

5. **Deploy**
   ```bash
   git push heroku main
   ```

#### AWS Elastic Beanstalk

1. Install EB CLI
2. Initialize: `eb init`
3. Create environment: `eb create production`
4. Deploy: `eb deploy`

#### Google Cloud Run

```bash
# Build and push image
gcloud builds submit --tag gcr.io/PROJECT_ID/socialapp-api

# Deploy
gcloud run deploy socialapp-api \
  --image gcr.io/PROJECT_ID/socialapp-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

---

## Environment Configuration

### Production Environment Variables

```env
# Database
DATABASE_URL=postgres://user:password@host:5432/socialapp?sslmode=require

# Security
JWT_SECRET=<generate-strong-random-key>

# SMTP (Production)
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USER=apikey
SMTP_PASSWORD=<your-sendgrid-api-key>

# Frontend
FRONTEND_URL=https://yourdomain.com

# Server
PORT=8080
ENVIRONMENT=production
```

### Generating Secure JWT Secret

```bash
# Linux/Mac
openssl rand -base64 32

# PowerShell
[Convert]::ToBase64String((1..32 | ForEach-Object { Get-Random -Minimum 0 -Maximum 256 }))
```

---

## Database Setup

### Backup Database

```bash
# Create backup
pg_dump -U postgres socialapp > backup_$(date +%Y%m%d).sql

# Restore backup
psql -U postgres socialapp < backup_20251002.sql
```

### Database Migrations

Migrations run automatically on application startup. To manually run:

```bash
# Connect to database
psql -U postgres socialapp

# Check tables
\dt

# View users
SELECT * FROM users;
```

---

## Monitoring & Maintenance

### Health Checks

```bash
# Check API health
curl https://api.yourdomain.com/health

# Expected response
{"status":"ok"}
```

### Logs

**Systemd service logs:**
```bash
sudo journalctl -u socialapp -f
```

**Docker logs:**
```bash
docker-compose logs -f api
```

### Performance Monitoring

**Database connections:**
```sql
SELECT count(*) FROM pg_stat_activity WHERE datname = 'socialapp';
```

**API metrics (recommended tools):**
- Prometheus + Grafana
- Datadog
- New Relic

### Backup Strategy

1. **Automated daily backups**
   ```bash
   # Add to crontab
   0 2 * * * pg_dump -U postgres socialapp > /backups/socialapp_$(date +\%Y\%m\%d).sql
   ```

2. **Retention policy**
   - Keep daily backups for 7 days
   - Keep weekly backups for 4 weeks
   - Keep monthly backups for 12 months

### Security Checklist

- [ ] Change default JWT_SECRET
- [ ] Use HTTPS/TLS in production
- [ ] Configure CORS for specific domains
- [ ] Enable database SSL (sslmode=require)
- [ ] Set up firewall rules
- [ ] Implement rate limiting
- [ ] Regular security updates
- [ ] Monitor failed login attempts
- [ ] Set up intrusion detection

### Scaling Considerations

**Horizontal Scaling:**
- Deploy multiple API instances behind load balancer
- Use managed PostgreSQL (AWS RDS, Google Cloud SQL)
- Implement Redis for session caching

**Vertical Scaling:**
- Increase server resources (CPU, RAM)
- Optimize database queries
- Add database indexes

---

## Troubleshooting

### Common Issues

**Issue: Cannot connect to database**
```bash
# Check PostgreSQL is running
sudo systemctl status postgresql

# Check connection
psql -U postgres -h localhost
```

**Issue: Port already in use**
```bash
# Find process using port 8080
sudo lsof -i :8080

# Kill process
sudo kill -9 <PID>
```

**Issue: Permission denied**
```bash
# Fix file permissions
chmod +x api
chown www-data:www-data api
```

---

## Rollback Procedure

1. **Stop current version**
   ```bash
   sudo systemctl stop socialapp
   ```

2. **Restore previous version**
   ```bash
   git checkout <previous-commit>
   go build -o api cmd/api/main.go
   ```

3. **Restore database if needed**
   ```bash
   psql -U postgres socialapp < backup.sql
   ```

4. **Restart service**
   ```bash
   sudo systemctl start socialapp
   ```

---

## Support & Resources

- **Documentation**: See README.md and API_DOCUMENTATION.md
- **Architecture**: See ARCHITECTURE.md
- **Quick Start**: See QUICK_START.md

---

**Last Updated:** October 2, 2025
