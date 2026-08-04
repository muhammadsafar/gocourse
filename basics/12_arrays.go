package main

import "fmt"

func main1Array() {
	// var arrayName [size]elementType;

	var numbers [5]int
	numbers[2] = 20
	fmt.Println(numbers)
	numbers[1] = 20
	numbers[3] = 30
	numbers[4] = 40
	numbers[0] = 90

	fmt.Println(numbers)

	fruits := [4]string{"durian", "apple", "grape", "guava"}
	fmt.Printf("buah array adalah %s", fruits)

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := &arr1 //pakai pointer artinya refer ke memory aslinya, bukan copy/buat array baru, jadi saat ubah value maka aslinya ikut berubah

	arr2[0] = 10

	fmt.Println("\narray 1 >>", arr1)
	fmt.Println("array 2 >>", arr2)

	for _, val := range arr2 {
		fmt.Printf("value adalah %v\n", val)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index ke %v adalah %v\n", i, numbers[i])
	}

	a, b := somFunc()
	fmt.Printf("value a is %v, value b is %v ", a, b)

	fmt.Print("\ncomparing arrays========\n")
	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println("array2 1:3", array2[1:3])

	fmt.Print("ärray 1 is eq to array 2 ", array1 == array2)

	fmt.Print("\n==========MATRIX==========\n")
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{11, 12, 13},
		{21, 22, 23},
	}

	fmt.Print(matrix)

}

//explicit return
func somFunc() (int, int) { // atau predefined(a int, b int)
	a := 120    // 	a = 120
	b := 20     // 	b = 20
	return a, b // predefined return
}
