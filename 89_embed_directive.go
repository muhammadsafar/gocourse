package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed example.txt
var content string

//go:embed intermediate
var intermediateDir embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	cont, err := intermediateDir.ReadFile("intermediate/output.txt")
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return
	}
	fmt.Println("Content of embedded file from intermediate directory:", string(cont))

	err = fs.WalkDir(intermediateDir, "intermediate", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Visited:", path)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking embedded directory:", err)
	}

}
