package main

import "fmt"

func main9() {

	//panic(intrface{})

	process2(10)
	process2(-1)

}

func process2(i int) {

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if i < 0 {
		fmt.Println("Before panic")
		panic("input must be a non negative number")
		//fmt.Println("After panic")

		//defer fmt.Println("Defer 3")
	}

	fmt.Println("processing input : ", i)
}
