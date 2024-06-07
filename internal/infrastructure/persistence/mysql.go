package persistence

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "go_user:password@tcp(127.0.0.1:3306)/go_crud_auth"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Database connection established")
}
