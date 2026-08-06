package main

import "fmt"

// chan<-  artinya cuma bisa send only
// <-chan receive only

func mainchannel_directtion() {

	ch := make(chan int)
	producer(ch)
	consumer(ch)
}

// send only
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// receive only
func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Recv : ", val)
	}
}
