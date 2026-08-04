package main

import (
	_ "embed"
	"fmt"
)

//go:embed output.txt
var content string

func main89() {

	fmt.Println("Embedded content:", content)

}
