package models

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// loadEnvVariable loads environment variables from .env file
func loadEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}

// SetupDatabase initializes the database connection and automigrates the models
func SetupDatabase() error {
	dbURL := loadEnvVariable("DB_URL")
	dsn := os.Getenv(dbURL)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Define User model
	type User struct {
		ID        uint    `json:"id" gorm:"primary_key"`
		Username  string  `json:"username" gorm:"unique"`
		Email     *string `json:"email" gorm:"unique"`
		Password  string  `json:"password"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	// AutoMigrate User model
	return DB.AutoMigrate(&User{})
}
