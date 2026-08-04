package main

import "fmt"

func main2() {

	//swicth expression{
	// case value:
	// code to be executed

	//case value 2:
	//code to be executed

	//default:
	//code else
	//}

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("An apple")

	case "banana":
		fmt.Println("the monkey's fruits")

	default:
		fmt.Println("Unknown fruit")
	}

	// var day string

	// fmt.Print("Enter day = ")
	// fmt.Scan(&day)

	day := "monday"

	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Weekday")

	case "saturday", "sunday":
		fmt.Println("yeeyy Weekend ")
	}

	num := 15
	switch {
	case num > 15:
		fmt.Println("> 15")
	case num > 10 && num < 15:
		fmt.Println("> 0 < 15")
	default:
		fmt.Println("15")
	}

	//check type
	checkType(10)
	checkType(12.4)
	checkType("Abdullah")
	checkType(false)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Its an integer")
	case int32:
		fmt.Println("Its an integer")

	case float64:
		fmt.Println("Its a float")

	case string:
		fmt.Println("Its String")

	default:
		fmt.Println("Unknown type")
	}
}
