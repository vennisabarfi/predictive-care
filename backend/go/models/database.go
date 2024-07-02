package models

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// load environment variables
func env_Variable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}

var DB *gorm.DB

func ConnectDB() {
	db_url := env_Variable("DB_URL")
	dsn := os.Getenv(db_url)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error accessing database", err)
	}

}

type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	username  string  `json:"username" gorm:"unique"`
	Email     *string `json:"email" gorm:"unique"`
	Password  string  `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
