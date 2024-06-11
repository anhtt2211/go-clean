package main

import (
	"fmt"
	"log"
	"net/http"

	"My-Clean/internal/application/use-cases"
	"My-Clean/internal/infrastructure/persistence"
	"My-Clean/internal/infrastructure/persistence/migrations"
	"My-Clean/internal/infrastructure/persistence/repositories"
	rest "My-Clean/internal/presentation/http"
)

func main() {
	persistence.Connect()
	db := persistence.DB

	// Automatically migrate the schema
	if err := migrations.AutoMigrate(); err != nil {
		log.Fatalf("Error migrating database schema: %v", err)
	}

	fmt.Println("Starting server on port 8000")

	userRepo := repositories.NewGORMUserRepository(db)
	taskRepo := repositories.NewGORMTaskRepository(db)

	userUseCase := use_cases.NewUserUseCase(userRepo)
	authUseCase := use_cases.NewAuthUseCase(userRepo)
	taskUseCase := use_cases.NewTaskUseCase(taskRepo)

	r := rest.NewRouter(userUseCase, authUseCase, taskUseCase)

	if err := http.ListenAndServe("127.0.0.1:8000", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
