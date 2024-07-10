package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

type Proverb struct {
	gorm.Model
	ID       uint              `json: "id"`
	Proverbs map[string]string `json:"proverbs"`
}

func InsertProverb() {
	jsonFile, err := os.Open("proverbs_data.json")
	if err != nil {
		println("File could not be found!")
		panic(err)
	}
	println("File found successfully!")
	defer jsonFile.Close()

	// Read the contents of the file into a []byte slice
	fileinfo, err := jsonFile.Stat() //find file size
	if err != nil {
		panic("Could not obtain stat/ file size in bytes")
	}
	data := make([]byte, fileinfo.Size()) //read file to retrieve json data
	count, err := jsonFile.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	println("JSON File successfully read into byte slice!")
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// println("Data:", data)

	// Parse json data into slice
	var proverbs Proverb //variable to hold parsed JSON data
	err = json.Unmarshal(data, &proverbs)
	if err != nil {
		println("Error unmarshaling json data")
		panic(err)
	}
	println("Successfully parsed json data!")

	// Access and print each proverb
	for key, proverb := range proverbs.Proverbs {
		fmt.Printf("Proverb %s: %s\n", key, proverb)
	}

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
