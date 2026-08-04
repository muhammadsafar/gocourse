package main

import (
	"bufio"
	"fmt"
	"os"
)

func main84() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()

	fmt.Println("File opened successfully")

	//read th contents of the file

	// data := make([]byte, 1024) //buffer to read file contents
	// _, err = file.Read(data)

	// if err != nil {
	// 	fmt.Println("Error reading file", err)
	// 	return
	// }

	// fmt.Println("file contents:\n", string(data))

	//=====using scanner==========
	scanner := bufio.NewScanner(file)

	//read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file line by line:", err)
		return
	}

}
