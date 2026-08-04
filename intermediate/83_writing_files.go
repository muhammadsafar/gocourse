package main

import (
	"fmt"
	"os"
)

func main83() {
	// Write your code here

	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Error creating files.", err)

		return
	}
	defer file.Close()

	//write some text to the file

	data := []byte("Hello, Gophers!")
	_, err = file.Write(data)

	if err != nil {
		fmt.Println("Error writing to file.", err)
		return
	}

	fmt.Println("File written successfully.")

	file, err = os.Create("WriteString.txt")
	if err != nil {
		fmt.Println("Error creating file.", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString("Hello, Gophers! Writing with WriteString method.")
	if err != nil {
		fmt.Println("Error writing to file.", err)
		return
	}
	fmt.Println("File written successfully using WriteString method.")

}
