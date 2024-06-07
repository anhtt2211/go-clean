package server

import (
	"My-Clean/internal/use-cases"
	"log"
	"net/http"

	"My-Clean/internal/infrastructure/persistence"
	"My-Clean/internal/infrastructure/persistence/repositories"
	"My-Clean/internal/presentation/http/handler"
	"github.com/gorilla/mux"
)

func main() {
	persistence.Connect()
	db := persistence.DB

	userRepo := repositories.NewMySQLUserRepository(db)
	userUseCase := use_cases.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
