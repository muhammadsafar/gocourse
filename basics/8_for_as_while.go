package main

import "fmt"

func main_foraswhile() {

	i := 1

	for i <= 5 {
		println("Iteration number:", i)
		i++
	}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}

		fmt.Println("Odd number:", num)
		num++
	}

}
