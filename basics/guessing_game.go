package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mainguess() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(10) + 1
	fmt.Println("Welcome to  the Guessing Game!")
	fmt.Println("I have chosen a number between 1 adn 100")
	fmt.Println("Can you uess what it is?")

	var guess int

	for {

		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		}
	}
}
