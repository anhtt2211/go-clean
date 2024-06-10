package http

import (
	"My-Clean/internal/application/use-cases"
	"My-Clean/internal/presentation/http/handler"
	"My-Clean/internal/presentation/http/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(userUseCase *use_cases.UserUseCase, authUseCase *use_cases.AuthUseCase) *mux.Router {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// Subrouters for public and private routes
	privateRoute := r.PathPrefix("/api").Subrouter()
	publicRoute := r.PathPrefix("/api").Subrouter()

	// Apply middleware for private routes
	privateRoute.Use(middleware.AuthMiddleware)

	// Auth routes
	authHandler := handler.NewAuthHandler(authUseCase)
	auth := publicRoute.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	// User routes
	userHandler := handler.NewUserHandler(userUseCase)
	users := privateRoute.PathPrefix("/users").Subrouter()
	users.HandleFunc("/me", userHandler.GetUsers).Methods("GET")
	users.HandleFunc("/", userHandler.GetUsers).Methods("GET")
	users.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	users.HandleFunc("/", userHandler.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PUT")
	users.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
