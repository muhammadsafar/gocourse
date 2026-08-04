package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main86() {

	relativepath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	//join path using filepath join

	joinPath := filepath.Join("home", "documents", "downloads", "file.zip")
	fmt.Println("Join path : ", joinPath)

	normalizedPath := filepath.Clean("/home/user/../user/docs/./file.txt")
	fmt.Println("Normalized Path : ", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Directory: ", dir)
	fmt.Println("File: ", file)
	fmt.Println(filepath.Base("/home/user/docs"))

	fmt.Println("is relative path:", filepath.IsAbs(relativepath))
	fmt.Println("is absolute path:", filepath.IsAbs(absolutePath))

	fmt.Println("extention : ", filepath.Ext(file))

	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println("Error finding relative path:", err)
		return
	}

	fmt.Println("Relative Path:", rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		fmt.Println("Error finding relative path:", err)
		return
	}

	fmt.Println("Relative Path:", rel)

	absPath, err := filepath.Abs(relativepath)
	if err != nil {
		fmt.Println("Error finding absolute path:", err)
		return
	} else {
		fmt.Println("Absolute Path:", absPath)
	}
}
