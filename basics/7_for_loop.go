package main

import "fmt"

func mainq1() {

	//simple for loop
	for i := 1; i <= 5; i++ {
		println("Iteration number:", i)
	}

	//iteration over collection
	numbers := []int{10, 20, 30, 40, 50}

	for i, num := range numbers {
		println("Index:", i, "Value:", num)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Odd number:", i)
		if i == 5 {
			break
		}
	}

	rows := 5
	for i := 1; i <= rows; i++ {
		//inner loop
		for j := 1; j <= rows; j++ {
			fmt.Print(" ")
		}

		//inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()

	}

	for i := range 10 {
		fmt.Println(10 - i)
	}
}
