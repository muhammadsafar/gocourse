package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	//sender di luar goroutine, receiver di dalam goroutine
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch)
		fmt.Println("Sending value to channel 3 secs later")
	}()
	ch <- 1

	// sender di dalam goroutine, receiver di luar goroutine
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Sending value to channel 2 secs later")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("Sending value to channel 3 secs later")
	// }()
	// receiver := <-ch
	// fmt.Println(receiver)

	fmt.Println("End of program")
}
