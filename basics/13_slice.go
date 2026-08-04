package main

import (
	"fmt"
	"slices"
)

func main2Slice() {

	// var slicename [] typeElement

	// var numbers []int
	// var num1 = []int{1, 2, 3}

	// num2 := []int{3, 4, 5}

	slice := make([]int, 5)
	slice[0] = 1

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4] //get elemen 1-4

	fmt.Println("slice by make() >>", slice)
	fmt.Println(slice1)

	//append
	slice1 = append(slice1, 6, 7)
	fmt.Println("append 6 7 ", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("slice copy>>", sliceCopy)

	//var nilSlice = []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Printf("slices is equal")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("slice twoD >>", twoD)

	//slice[low :high]
	slice2 := slice1[2:4]

	fmt.Println("slice1>>", slice1)
	fmt.Println("slices 2 copt slice1 got elemen 2:4>>", slice2)
	fmt.Println("the capacity of slice2 is >>", cap(slice2))
	fmt.Println("the le of slice2 is >>", len(slice2))

}
