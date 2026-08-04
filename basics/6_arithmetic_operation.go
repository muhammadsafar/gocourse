package main

import "fmt"

func mainq2() {

	var a, b int = 10, 3

	var res int

	res = a + b
	println("Addition:", res)

	res = a - b
	println("Subtraction:", res)

	res = a * b
	println("Multiplication:", res)

	res = a / b
	println("Division:", res)

	res = a % b
	println("Modulus:", res)

	const p float32 = 3.14
	fmt.Println("phi >>", p)
}
