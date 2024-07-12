package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"user_auth/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Structure of each proverb(text field)
type Proverb struct {
	gorm.Model
	Text string `json:"text"`
}

// helper struct to decode JSON data correctly
type ProverbsData struct {
	Proverbs []string `json:"Proverbs"`
}

func InsertProverb() {
	// Open JSON file
	jsonFile, err := os.Open("proverbs_only.json")
	if err != nil {
		log.Fatalf("File could not be found: %v", err)
	}
	defer jsonFile.Close()
	fmt.Println("File found successfully!")

	// Decode JSON file
	var proverbsData ProverbsData
	err = json.NewDecoder(jsonFile).Decode(&proverbsData)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	fmt.Println("Successfully decoded JSON data!")

	// Print each proverb (for testing)
	// for index, proverb := range proverbsData.Proverbs {
	// 	fmt.Printf("Proverb %d: %s\n", index+1, proverb)
	// }

	// Establish PostgreSQL connection
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	//Refactor this to call on storage
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode),
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to Database!.")

	// Auto migrate Proverb table
	err = db.AutoMigrate(&Proverb{})
	if err != nil {
		log.Fatalf("Error migrating Proverb table: %v", err)
	}
	fmt.Println("Proverb table migrated successfully.")

	// Prepare Proverb records for batch insertion
	var proverbRecords []Proverb
	for _, text := range proverbsData.Proverbs {
		proverbRecords = append(proverbRecords, Proverb{Text: text})
	}

	// Batch insert data into database
	result := db.Create(&proverbRecords)
	if result.Error != nil {
		log.Fatalf("Error inserting proverbs into database: %v", result.Error)
	}
	fmt.Println("Proverbs successfully inserted into database!")
}
