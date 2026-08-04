package main

import "fmt"

func main59() {

	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr) //0xc00000a098
	fmt.Println(*ptr)

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
