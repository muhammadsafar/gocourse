package main

import "fmt"

func main8() {

	process(30)

}

// defer likes async await in node js

func process(i int) {
	defer fmt.Println("defer value of i : ", i)          // 6 =30
	defer fmt.Println("Deferred statement executed 1st") //5
	defer fmt.Println("Deferred statement executed 2nd") //4
	defer fmt.Println("Deferred statement executed 3rd") //3

	i++

	fmt.Println("Normal statement executed") //1
	fmt.Println("value of i : ", i)          //2 = 31
}
