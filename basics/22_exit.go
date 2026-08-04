package main

import (
	"fmt"
	"os"
)

func main11() {

	defer fmt.Println("Deferred statement..")

	fmt.Println("Starting the main function..")

	os.Exit(1)

	//this never be executed
	fmt.Println("End of function")

}
