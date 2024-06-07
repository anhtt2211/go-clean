package main

import (
	use_cases2 "My-Clean/internal/application/use-cases"
	"My-Clean/internal/presentation/http/middleware"
	"fmt"
	"net/http"

	"My-Clean/internal/infrastructure/persistence"
	"My-Clean/internal/infrastructure/persistence/repositories"
	"My-Clean/internal/presentation/http/handler"
	"github.com/gorilla/mux"
)

func main() {
	persistence.Connect()
	db := persistence.DB

	fmt.Println("Starting server on port 8000")

	userRepo := repositories.NewGORMUserRepository(db)

	userUseCase := use_cases2.NewUserUseCase(userRepo)
	authUseCase := use_cases2.NewAuthUseCase(userRepo)

	userHandler := handler.NewUserHandler(userUseCase)
	authHandler := handler.NewAuthHandler(authUseCase)

	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	privateRoute := r.PathPrefix("/api").Subrouter()
	publicRoute := r.PathPrefix("/api").Subrouter()

	privateRoute.Use(middleware.AuthMiddleware)

	auth := publicRoute.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	users := privateRoute.PathPrefix("/users").Subrouter()
	users.HandleFunc("/", userHandler.GetUsers).Methods("GET")
	users.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	users.HandleFunc("/", userHandler.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PUT")
	users.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")

	http.ListenAndServe("127.0.0.1:8000", r)
}
