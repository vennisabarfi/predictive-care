package handlers

import (
	"fmt"
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
func insertProverb() {

	// plan, _ := ioutil.ReadFile()
}
