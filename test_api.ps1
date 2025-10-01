# PowerShell script to test the Social App API

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Social App API Test Script" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$BASE_URL = "http://localhost:8080"
$API_URL = "$BASE_URL/api"

# Test 1: Health Check
Write-Host "Test 1: Health Check" -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BASE_URL/health" -Method Get
    Write-Host "✓ Health check passed" -ForegroundColor Green
    Write-Host "  Response: $($response | ConvertTo-Json)" -ForegroundColor Gray
} catch {
    Write-Host "✗ Health check failed: $_" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Test 2: User Registration
Write-Host "Test 2: User Registration" -ForegroundColor Yellow
$registerData = @{
    email = "test@example.com"
    username = "testuser"
    password = "password123"
    first_name = "Test"
    last_name = "User"
    interests = @(1, 2)
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_URL/auth/register" -Method Post -Body $registerData -ContentType "application/json"
    Write-Host "✓ Registration successful" -ForegroundColor Green
    $token = $response.data.token
    Write-Host "  User ID: $($response.data.user.id)" -ForegroundColor Gray
    Write-Host "  Email: $($response.data.user.email)" -ForegroundColor Gray
    Write-Host "  Token: $($token.Substring(0, 20))..." -ForegroundColor Gray
} catch {
    Write-Host "✗ Registration failed: $_" -ForegroundColor Red
    if ($_.Exception.Response) {
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $reader.BaseStream.Position = 0
        $responseBody = $reader.ReadToEnd()
        Write-Host "  Error details: $responseBody" -ForegroundColor Red
    }
}
Write-Host ""

# Test 3: User Login
Write-Host "Test 3: User Login" -ForegroundColor Yellow
$loginData = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_URL/auth/login" -Method Post -Body $loginData -ContentType "application/json"
    Write-Host "✓ Login successful" -ForegroundColor Green
    $token = $response.data.token
    Write-Host "  Token received: $($token.Substring(0, 20))..." -ForegroundColor Gray
} catch {
    Write-Host "✗ Login failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 4: Get Profile (Protected Route)
Write-Host "Test 4: Get Profile (Protected)" -ForegroundColor Yellow
try {
    $headers = @{
        "Authorization" = "Bearer $token"
    }
    $response = Invoke-RestMethod -Uri "$API_URL/auth/profile" -Method Get -Headers $headers
    Write-Host "✓ Profile retrieved successfully" -ForegroundColor Green
    Write-Host "  User ID: $($response.data.user_id)" -ForegroundColor Gray
    Write-Host "  Email: $($response.data.email)" -ForegroundColor Gray
} catch {
    Write-Host "✗ Profile retrieval failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 5: Password Reset Request
Write-Host "Test 5: Password Reset Request" -ForegroundColor Yellow
$resetRequestData = @{
    email = "test@example.com"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_URL/auth/password-reset/request" -Method Post -Body $resetRequestData -ContentType "application/json"
    Write-Host "✓ Password reset request sent" -ForegroundColor Green
    Write-Host "  Message: $($response.data.message)" -ForegroundColor Gray
    Write-Host "  (Check console logs for reset token in development mode)" -ForegroundColor Gray
} catch {
    Write-Host "✗ Password reset request failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 6: Invalid Login
Write-Host "Test 6: Invalid Login (Should Fail)" -ForegroundColor Yellow
$invalidLoginData = @{
    email = "test@example.com"
    password = "wrongpassword"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_URL/auth/login" -Method Post -Body $invalidLoginData -ContentType "application/json"
    Write-Host "✗ Invalid login should have failed but didn't" -ForegroundColor Red
} catch {
    Write-Host "✓ Invalid login correctly rejected" -ForegroundColor Green
}
Write-Host ""

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  All Tests Completed!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
