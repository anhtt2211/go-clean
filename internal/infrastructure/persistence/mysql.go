package persistence

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:userpassword@tcp(127.0.0.1:3306)/?parseTime=true" // Remove database name from DSN
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto create the database if it doesn't exist
	err = DB.Exec("CREATE DATABASE IF NOT EXISTS go_crud_auth").Error
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// Connect to the database
	dsnWithDB := "user:userpassword@tcp(127.0.0.1:3306)/go_crud_auth?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database connection established")
}
