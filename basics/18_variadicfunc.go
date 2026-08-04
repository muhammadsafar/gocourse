package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main7() {

	// 	func funcname (param1 type1, param2 type2 param3 ... type3) return type{
	// block code
	// return
	// }

	numbers := []int{1, 23, 45, 56, 67, 89, 90}

	listnum, total := sum(numbers...)
	fmt.Printf("Total list num %v >> %v", listnum, total)

}

func sum(arrnum ...int) (string, int) {

	// ubah []int menjadi []string
	strSlice := make([]string, len(arrnum))
	for i, v := range arrnum {
		strSlice[i] = strconv.Itoa(v)
	}

	// join slice string
	listnum := strings.Join(strSlice, ",")

	// hitung total
	total := 0
	for _, v := range arrnum {
		total += v
	}
	return listnum, total
}
