package main

import (
	"fmt"
	"os"
)

func checkError2(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main88() {

	temFileName := "temporaryFile"

	tmpFile, err := os.CreateTemp("", temFileName)
	checkError2(err)

	fmt.Println("Temporary file is created..", tmpFile.Name())

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError2(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created: ", tempDir)

}
