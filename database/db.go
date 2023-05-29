package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

func ConnectDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	connectionDSN := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", HOST, PORT, USER, DB_NAME, PASS)
	fmt.Println("connection name is\t\t", connectionDSN)
	db, err := gorm.Open(postgres.Open(connectionDSN), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
