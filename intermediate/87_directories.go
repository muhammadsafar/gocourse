package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

func main87() {

	err := os.Mkdir("subdir", 0755)
	checkError(err)

	checkError(os.Mkdir("subdir1", 0755))

	defer os.RemoveAll("subdir")

	os.WriteFile("subdir1/file", []byte("Hello, World!"), 0755)

	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	res, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range res {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))

	res, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range res {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("../../.."))

	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	//filepath .walk and filepath .walkdir are more advanced ways to traverse directories

	pathfile := "subdir/parent/child"
	fmt.Println("working directory..")

	filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", path, "error:", err)
			return err
		}

		fmt.Println(path)

		return nil
	})

	checkError(err)

	// checkError(os.RemoveAll("subdir"))
	checkError(os.Remove("subdir1"))

	checkError(err)
}
