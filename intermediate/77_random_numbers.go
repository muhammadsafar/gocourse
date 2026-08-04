package main

import (
	"fmt"
	"math/rand"
)

func main77() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println(val.Intn(101))

	fmt.Print(rand.Intn(101))

	fmt.Print(rand.Float64()) //0.0 and 1.0

	for {
		fmt.Println("Welcome to the golang game!")
		fmt.Println("1. Roll the disc")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice (1 or 2)")

		var ch int

		_, err := fmt.Scan(&ch)
		if err != nil || ch != 1 && ch != 2 {
			fmt.Println("Invalid choice,please enter 1 or 2")
			continue
		}

		if ch == 2 {
			fmt.Println("Thanks for playing..")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show the results
		fmt.Printf("you rolled a %d and %d \n", die1, die2)
		fmt.Println("Total: ", die1+die2)

		//ask of the user want to roll

		fmt.Print("Do you want to roll again? (y/n):")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)

		if err != nil || rollAgain != "y" && rollAgain != "n" {
			fmt.Println("Invalid input, assumin no.")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing...")
			break
		}

	}
}
