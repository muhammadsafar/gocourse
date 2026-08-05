package main

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Goroutine is working...")
// 		time.Sleep(5 * time.Second)

// 		done <- struct{}{} // Mengirim sinyal selesai ke channel, simple ping struct kosong 0 byte
// 		// channel anything bisa bool, int, dll even false yg penting return val
// 	}()

// 	<-done // Menunggu sinyal selesai dari channel
// 	fmt.Println("Main function is done.")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		fmt.Println("Sending...")
// 		ch <- 9 //blocking until the value is received by the receiver
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch //blocking until the value is sent to the channel by the sender
// 	fmt.Println("Received value from channel:", value)
// }

// ==============SYNCHRONIZING MULTIPLE GOROTINES===============
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(1 * time.Second)
// 			done <- 1 //val bebas
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done //wait each go finish
// 	}

// 	fmt.Println("All Gorotines are finished")
// }

//============sync data exchange============

func mainChannelsync() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) //channel closed before send a value to the channel

	for value := range data {
		fmt.Println("Recv value : ", value, ":", time.Now())
	}
}
