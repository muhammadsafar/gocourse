package main

import "fmt"

func main10() {

	process3()
	fmt.Println("Return from process")
}

func process3() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r) //if no panic, recover will return nil
		}
	}()

	fmt.Println("start process")
	panic("Something went wrong")
	fmt.Println("End process")
}
