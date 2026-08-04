package main

import (
	"errors"
	"fmt"
)

func main6MultiReturn() {

	// func functionName (param1 type, param2 type.... ) (returntype1, returtype2 ..){
	// 	//code block
	// 	return v1, v2 ...
	// }

	bagi, persen := divide(12, 4)
	fmt.Printf("12 / 4 > div : %v, remind %v \n", bagi, persen)

	res, err := compare(5, 5)
	if err != nil {
		fmt.Printf("Error found : %s", err)
	} else {
		fmt.Println(res)
	}

}

func divide(a, b int) (int, int) {

	quotient := a / b
	reminder := a % b

	return quotient, reminder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare which is greater")
	}
}
