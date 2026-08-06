package main

import (
	"fmt"
	"time"
)

func main_nonblocking() {

	// ch := make(chan int)

	//==NON BLOCKING RECEIVE OPERATION
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Receive : ", msg)
	// default:
	// 	fmt.Println("No messae available ")
	// }

	//==NON BLOCKING SEND OPERATION
	// select {
	// case ch <- 1:
	// 	fmt.Println("sent >> ", ch)
	// default:
	// 	fmt.Println("Channel isnt ready to receive")
	// }

	//==NON BLOCKING OPERATION IN REAL TIME SYSTEMs

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received : ", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("waiting for...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true

}
