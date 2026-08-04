package main

import (
	"fmt"
)

const Pi = 3.14
const API_KEY = "12345-ABCDE"

func main6() {
	fmt.Println("Happy", Pi, " KEY:", API_KEY)

	const day int = 7

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
		Sunday    = 7
	)

	fmt.Println("Days in a week:", day)
	fmt.Println("Wednesday is day number:", Wednesday)
}
