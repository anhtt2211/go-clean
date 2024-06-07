package persistence

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:userpassword@tcp(127.0.0.1:3306)/go_crud_auth"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Database connection established")
}
