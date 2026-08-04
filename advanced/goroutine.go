package main

import (
	"fmt"
	"time"
)

//Goroutines are just funstion that leave the main thread and run concurrently with the main thread. They are lightweight and managed by the Go runtime. We can create a goroutine by using the go keyword followed by a function call. The function will run in a separate goroutine, allowing the main thread to continue executing other code without waiting for the goroutine to finish.

//Goroutines do not stop the program flow and are non blocking

func main() {

	// var wg sync.WaitGroup

	var err error

	// wg.Add(1)
	fmt.Println("Starting main function")
	// go sayHello(&wg)
	go sayHello()
	fmt.Println("After sayhello Main function is doing other work")
	// wg.Wait()

	go func() {
		err = doWork()
		fmt.Println("err sudah diisi:", err)
	}()

	// err = doWork()
	go printNumbers()
	go printLetters()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("work completed successfully")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Main function is done")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	println("Hello from goroutine")
	// defer wg.Done()
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	fmt.Println("doWork started")

	time.Sleep(1 * time.Second)

	fmt.Println("doWork finished")

	return fmt.Errorf("Something went wrong")
}
