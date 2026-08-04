package main

import (
	"fmt"
)

// UNBUFFERED CHANNELS SAMPLE

func main2() {
	// Create a channel of type int
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "World"

		for _, e := range "abcde" {
			greeting <- "Alphabetic: " + string(e)
		}

		// close(greeting) // Close the channel after sending all values
	}()

	//pakai close channel
	// for receivedGreeting := range greeting {
	// 	fmt.Println(receivedGreeting)
	// }

	// go func() {
	// 	receivedGreeting := <-greeting
	// 	fmt.Println(receivedGreeting)
	// 	receivedGreeting = <-greeting
	// 	fmt.Println(receivedGreeting)

	// 	for range 5 {
	// 		rcvr := <-greeting
	// 		fmt.Println(rcvr)
	// 	}

	// }()

	receivedGreeting := <-greeting
	fmt.Println(receivedGreeting)
	receivedGreeting = <-greeting
	fmt.Println(receivedGreeting)
	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	// time.Sleep(1 * time.Second) // Wait for goroutines to finish, --> kalau receiver dan sender tidak dalam goroutine yang sama, maka perlu waktu untuk baca receiver
	fmt.Println("End of program.")
}
