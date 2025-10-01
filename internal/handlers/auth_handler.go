package handlers

import (
	"encoding/json"
	"net/http"

	"windsurf-project/internal/models"
	"windsurf-project/internal/service"
	"windsurf-project/pkg/response"
	"windsurf-project/pkg/validator"
)

type AuthHandler struct {
	authService  *service.AuthService
	emailService *service.EmailService
}

func NewAuthHandler(authService *service.AuthService, emailService *service.EmailService) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		emailService: emailService,
	}
}

// Register handles user registration
// POST /api/auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if err := validator.ValidateEmail(req.Email); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validator.ValidateUsername(req.Username); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Register user
	authResp, err := h.authService.Register(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Send welcome email (async, don't block on failure)
	go h.emailService.SendWelcomeEmail(req.Email, req.Username)

	response.Created(w, authResp)
}

// Login handles user login
// POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if err := validator.ValidateEmail(req.Email); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validator.ValidateRequired("password", req.Password); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Login user
	authResp, err := h.authService.Login(&req)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, authResp)
}

// RequestPasswordReset handles password reset requests
// POST /api/auth/password-reset/request
func (h *AuthHandler) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var req models.PasswordResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if err := validator.ValidateEmail(req.Email); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Request password reset
	token, err := h.authService.RequestPasswordReset(req.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to process password reset request")
		return
	}

	// Send password reset email (async)
	if token != "" {
		go h.emailService.SendPasswordResetEmail(req.Email, token)
	}

	// Always return success to prevent email enumeration
	response.Success(w, map[string]string{
		"message": "If the email exists, a password reset link has been sent",
	})
}

// ResetPassword handles password reset confirmation
// POST /api/auth/password-reset/confirm
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req models.PasswordResetConfirm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate input
	if err := validator.ValidateRequired("token", req.Token); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validator.ValidatePassword(req.NewPassword); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Reset password
	if err := h.authService.ResetPassword(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, map[string]string{
		"message": "Password has been reset successfully",
	})
}

// GetProfile returns the current user's profile
// GET /api/auth/profile
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// User claims are set by the auth middleware
	claims := r.Context().Value("user")
	if claims == nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	response.Success(w, claims)
}
