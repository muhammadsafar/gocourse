package main

import (
	"fmt"
)

func main62() {

	//Printing Functions
	fmt.Print("Hello")
	fmt.Print("Go!")
	fmt.Print(12.45)

	fmt.Println("Happy")
	fmt.Println("Coding Go!")
	fmt.Println(12.345)

	name := "Abdullah"
	age := 3

	fmt.Printf("Name: %s, Age: %d \n", name, age)
	fmt.Printf("Binary %b , Hex %X \n", age, age)

	//Formatting functions
	s := fmt.Sprint("Hello ", "Coding", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World", 123, 345)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name %s, Age %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	//Scannin function
	var nm string
	var ag int

	fmt.Print("Enter your name and age >> ")
	//fmt.Scan(&nm, &ag) // setelah input 1 enter, akan req input ke dua
	// fmt.Scanln(&nm, &ag) // saat input 1 enter langsung finish , tanpa tanya req input ke dua
	fmt.Scanf("%s %d", &nm, &ag) //sama seperti scanln
	fmt.Printf("Name : %s , Age : %d \n", nm, ag)

	//Error Formatting Function

	err := checkAge(17)
	if err != nil {
		fmt.Println("Error : ", err)
	}

}

func checkAge(ag int) error {
	if ag < 18 {
		return fmt.Errorf("age %d is too young to drive", ag)
	}

	return nil
}
