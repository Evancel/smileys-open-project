package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"windsurf-project/internal/config"
	"windsurf-project/internal/handlers"
	"windsurf-project/internal/middleware"
	"windsurf-project/internal/repository"
	"windsurf-project/internal/service"
)

type Server struct {
	config  *config.Config
	db      *sql.DB
	router  *mux.Router
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	s := &Server{
		config: cfg,
		db:     db,
		router: mux.NewRouter(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// Initialize repositories
	userRepo := repository.NewUserRepository(s.db)

	// Initialize services
	authService := service.NewAuthService(userRepo, s.config.JWTSecret)
	emailService := service.NewEmailService(s.config)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, emailService)

	// API router with middleware
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(middleware.Logging)
	api.Use(middleware.CORS)

	// Health check
	s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Public routes
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	api.HandleFunc("/auth/password-reset/request", authHandler.RequestPasswordReset).Methods("POST")
	api.HandleFunc("/auth/password-reset/confirm", authHandler.ResetPassword).Methods("POST")

	// Protected routes (require authentication)
	protected := api.PathPrefix("/auth").Subrouter()
	protected.Use(middleware.Auth(authService))
	protected.HandleFunc("/profile", authHandler.GetProfile).Methods("GET")

	// Handle OPTIONS for CORS preflight
	s.router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
