package handlers

import (
	"fmt"
	"log"
	"os"
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

// insert proverb data into gorm
// read file and then read json data
// write into a list and then batch insert
func InsertProverb() {
	filename := "proverbs_cleaned.json"
	file, err := os.Open(filename) //read file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //ensure file is closed

	//read file data into a slice of bytes
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
