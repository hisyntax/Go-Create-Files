package main 

import (
	"os"
)

func main() {
	// Create the file
	file, err := os.Create("testing.txt")
	if err != nil {
		//Handle the error

		return
	}
	//Close the file
	defer file.Close()
	
	// Write data into the created file
	file.WriteString("data")
}
