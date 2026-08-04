package main

import (
	"fmt"
	"strconv"
)

func main78() {

	numStr := "12345asa"

	//using Atoi
	num, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println("Error parsing", err)
	}
	fmt.Println("String to nu + 10 ", num+10)

	floatStr := "9.18"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing float", err)
	}
	fmt.Println("String to float >> ", floatNum)

	//using ParseInt
	n, err := strconv.ParseInt(numStr, 10, 32) //decimal base 10 , bit size 32
	if err != nil {
		fmt.Println("Error parsing int", err)
	}

	fmt.Println(n)

	binaryStr := "1101"
	binaryNum, err := strconv.ParseInt(binaryStr, 2, 64) //base 2 01
	if err != nil {
		fmt.Println("Error parsing binary", err)
	}
	fmt.Println("Parsing binary to decimal >> ", binaryNum)

	hexStr := "1a3f"
	hexNum, err := strconv.ParseInt(hexStr, 16, 64) //base 16
	if err != nil {
		fmt.Println("Error parsing hex", err)
	}
	fmt.Println("Parsing hex to decimal >> ", hexNum)

}
