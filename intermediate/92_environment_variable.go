package main

import (
	"fmt"
	"os"
	"strings"
)

func main92() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User:", user)
	fmt.Println("Home Directory:", home)

	err := os.Setenv("FRUIT", "APPLE")

	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
	fruit := os.Getenv("FRUIT")
	fmt.Println("FRUIT:", fruit)

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}
	fruit = os.Getenv("FRUIT")
	fmt.Println("FRUIT after unsetting:", fruit)

	str := "a=b=c=d"

	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))

	// for _, e := range os.Environ() {
	// 	kupair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(kupair[0])
	// }

	//a=b=c=d
}
