package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
    conn, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to database")
        panic(err)
    }// Log SQL queries for debugging purposes

    DB = conn
    return DB
}
