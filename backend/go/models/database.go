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

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
}

func MigrateDatabase() error {
	// AutoMigrate all models here
	if err := DB.AutoMigrate(&User{}); err != nil {
		return err
	}
	// Add more AutoMigrate calls for other models if needed

	return nil
}

//Models defined below

// User model
type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	username  string  `json:"username" gorm:"unique"`
	Email     *string `json:"email" gorm:"unique"`
	Password  string  `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
