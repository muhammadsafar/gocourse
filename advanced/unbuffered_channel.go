package main

import (
	"fmt"
	"time"
)

//unbuffered channel adalah channel yang tidak memiliki kapasitas tertentu, sehingga setiap pengiriman value ke channel akan diblokir sampai ada receiver yang menerima value tersebut. Jika tidak ada receiver yang menerima value dari channel, maka pengiriman value ke channel akan menunggu sampai ada receiver yang siap untuk menerima value tersebut.

func mainunbuffer() {

	ch := make(chan int)

	//sender di luar goroutine, receiver di dalam goroutine
	go func() {
		// time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
		fmt.Println("received value from channel")
	}()
	ch <- 1
	time.Sleep(2 * time.Second)

	// sender di dalam goroutine, receiver di luar goroutine
	// go func() {
	// 	ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Sending value to channel 2 secs later")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Sending value to channel 3 secs later")
	// }()
	// receiver := <-ch
	// fmt.Println(receiver)

	fmt.Println("End of program")
}
