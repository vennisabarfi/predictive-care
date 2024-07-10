package handlers

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// get current directory
func Getwd() (dir string, err error) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not retrieve current directory", err)
	}
	fmt.Println(mydir)
	return
}

type Entry struct {
	gorm.Model
	Proverbs string `json:"proverbs"`
}

func InsertProverb() {
	file, err := os.Open("proverbs_data.json")
	if err != nil {
		println("File could not be found!")
		panic(err)
	}
	defer file.Close()

	// Parse json data into slice
}

// insert proverb data into gorm
// read file and then read json data
// write into a list and then batch insert
// func InsertProverb() {
// 	filename := "proverbs_cleaned.json"
// 	file, err := os.Open(filename) //read file
// 	if err != nil {
// 		log.Fatal("Could not locate file!", err)
// 	}
// 	println("File found successfully!", filename)
// 	defer file.Close() //ensure file is closed

// 	//read file data into a slice of bytes
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatal("Error reading file", err)
// 	}
// 	println("File read successfully!")
// 	fmt.Printf("read %d bytes: %q\n", count, data[:count])

// }
