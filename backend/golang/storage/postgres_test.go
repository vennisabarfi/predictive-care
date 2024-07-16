package storage

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

var db *gorm.DB

// set up mock database for testing
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Unexpected error when setting up mock database connection %v", err)
	}

	mock
}
