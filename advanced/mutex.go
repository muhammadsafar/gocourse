package main

import (
	"fmt"
	"sync"
)

func mainmutex() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGo := 5
	wg.Add(numGo)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGo {
		go increment()
	}

	wg.Wait()
	fmt.Println("Final Total value >>", counter)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := &counter{}

// 	numGoroutine := 10
// 	for i := range numGoroutine {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Printf("Grooroutine %d\n", i)
// 			for range 1000 {
// 				counter.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final counter value : %d \n", counter.getValue())

// }
