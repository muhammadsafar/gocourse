package main

import "fmt"

func main5Function() {

	// func name(param list) return type {
	//block code
	// return
	// }

	res := adds(120, 34)
	fmt.Println("func add >>", res)

	var ll = Lingkaran{
		r: 7,
	}
	luasLingkaran := ll.luas()
	fmt.Println("method luas lingkaran>>", luasLingkaran)

	great := func() {
		fmt.Println("Happy coding!!!")
	}

	great()

	operation := adds

	result := operation(12, 56)

	fmt.Println(result)

	result2 := applyOperation(5, 3, adds)
	fmt.Println("5 + 3 = ", result2)

	//return and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Print("6 * 2 = ", multiplyBy2(6))
}
func adds(a, b int) int {
	return a + b
}

type Lingkaran struct {
	r int
}

func (l Lingkaran) luas() int {
	return l.r * l.r
}

// function that take a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
