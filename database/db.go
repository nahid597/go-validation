package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Set default values if environment variables are not set
	dbUser := getEnvOrDefault("DB_USER", "gouser")
	dbPass := getEnvOrDefault("DB_PASS", "gopassword")
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := getEnvOrDefault("DB_PORT", "3306")
	dbName := getEnvOrDefault("DB_NAME", "go_validator")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Printf("Attempting to connect to database at %s:%s...\n", dbHost, dbPort)

	// Retry connection with backoff
	var database *gorm.DB
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		fmt.Printf("Database connection attempt %d failed: %v\n", i+1, err)
		if i < maxRetries-1 {
			fmt.Printf("Retrying in %d seconds...\n", (i+1)*2)
			time.Sleep(time.Duration((i+1)*2) * time.Second)
		}
	}

	if err != nil {
		log.Printf("Failed to connect to database after %d attempts: %v", maxRetries, err)
		log.Printf("Database config: user=%s, host=%s, port=%s, database=%s", dbUser, dbHost, dbPort, dbName)
		log.Fatal("Failed to connect to database!")
	}

	DB = database
	fmt.Println("Database connection established successfully!")
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
