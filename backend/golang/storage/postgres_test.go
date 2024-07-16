package storage

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// var gormDB *gorm.DB

// set up mock database for testing
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {

	testdb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Unexpected error when setting up mock database connection %v", err)
	}

	// var dialector *gorm.DB

	// connect to mock db
	// Create the GORM dialector using the mock database connection
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testdb,
		PreferSimpleProtocol: true,
	})

	if err != nil {
		log.Fatalf("Error running GORM dialector %v", err)
	}

	//connect to database using dialector
	db, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to dialector with GORM %v", err)
	}

	return db, mock //use to test database operations

}
