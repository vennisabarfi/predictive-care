package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

// func NewConnection(config *Config) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
// 	)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return db, err
// 	}
// 	return db, nil
// }

func GetDBConfig() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func ConnectToDB() (*gorm.DB, error) {
	config := GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	fmt.Println("Connected to Database!")
	return db, nil
}
