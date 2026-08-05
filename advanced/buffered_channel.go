package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int, 2) // buffered channel dengan kapasitas 2

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 	}()
// 	fmt.Println("value : ", <-ch) // akan menunggu sampai ada value yang dikirim ke channel
// 	// fmt.Println("value : ", <-ch) // akan menunggu sampai ada value yang dikirim ke channel
// 	fmt.Println("End of program")
// }

//buffered channel adalah channel yang memiliki kapasitas tertentu, sehingga dapat menampung beberapa value sebelum harus menunggu receiver untuk menerima value tersebut. Jika kapasitas channel penuh, maka pengiriman value ke channel akan diblokir sampai ada receiver yang menerima value dari channel.

func mainbuffered() {
	// =========BLOCKING on send only IF THE BUFFER IS FULL=========
	// make(chan Type, capacity) // buffered channel dengan kapasitas tertentu
	ch := make(chan int, 2) // buffered channel dengan kapasitas 2
	ch <- 1
	ch <- 2
	go func() {
		fmt.Println("waiting for 2 seconds")
		time.Sleep(2 * time.Second)
		fmt.Println("Received : ", <-ch)
	}()
	fmt.Println("Blocking started, channel is full")
	ch <- 3 // akan menunggu sampai ada receiver yang menerima value dari channel
	fmt.Println("Blocking ends")
	// fmt.Println("Received : ", <-ch)
	// fmt.Println("Received : ", <-ch)

	fmt.Println("End of program")
}
