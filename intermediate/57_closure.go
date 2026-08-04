package main

import "fmt"

func main57() {

	/**
	  closure adalah fungsi yang memanggil / membawa var dari scope luar
	*/

	seq := adder()

	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	seq2 := adder()
	fmt.Println(seq2())

	subtractor := func() func(int) int {
		countdown := 99

		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println("Subtractor >> ")
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(12)) // -12
	fmt.Println(subtractor(10)) // -10
	fmt.Println(subtractor(2))  // -2

}

func adder() func() int {
	i := 0
	fmt.Println("preveious value of i :", i)
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
