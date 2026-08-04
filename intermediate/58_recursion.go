package main

import "fmt"

func main58() {

	fmt.Println("factorial >>")
	fmt.Println(factorial(5))
	fmt.Println(factorial(3))

	fmt.Println("sumOfDigits >>")
	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

}

func factorial(n int) int {
	// base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	//recursive case : factorial of n is n + fatorial (n-1)
	return n * factorial(n-1)
	// n * (n-1) * (n-2) + factorial (n-3) ... factorial
}

func sumOfDigits(n int) int {
	//base case
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)

}
