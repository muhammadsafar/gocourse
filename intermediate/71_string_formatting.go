package main

import "fmt"

func main71() {

	num := 424
	fmt.Printf("%05d\n", num)

	msg := "Hello"
	fmt.Printf("[%10s]\n", msg)
	fmt.Printf("[%-10s]\n", msg)

	msg2 := "Hello \nworld!"
	msg3 := "Hello \nworld!"

	fmt.Println(msg2)
	fmt.Println(msg3)

	//sqlQuery := `SELECT * FROM users WHERE age > 30`

}
