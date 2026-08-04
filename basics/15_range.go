package main

import "fmt"

func main4Range() {

	msgTest := "Happy coding!"

	for i, v := range msgTest {
		fmt.Printf("char -%v is %c \n", i, v)
	}

}
