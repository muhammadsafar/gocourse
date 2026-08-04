package main

import (
	"flag"
	"fmt"
	"os"
)

func main90() {

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}
	// fmt.Println("Arguments:", os.Args[1])

	// define flags
	var name string
	var age int
	var ismale bool

	flag.StringVar(&name, "name", "John", "name of the user")
	flag.IntVar(&age, "age", 16, "age of the user")
	flag.BoolVar(&ismale, "ismale", true, "is the user male")

	//parse the flags
	flag.Parse()

	fmt.Println("Name:", name) //use "" for name multi words
	fmt.Println("Age:", age)
	fmt.Println("Is Male:", ismale)

}
