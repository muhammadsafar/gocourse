package main

import (
	"errors"
	"fmt"
)

func main69() {

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Operation completed successfully!")

}

type customError struct {
	code    int
	message string
	err     error
}

// Error return the error message.Implementing Error() method of error interface
/**
bawaan Go, aslinya Error() adalah

type error interface {
    Error() string
}

butuh return string

*/

/*
setiap bisa jadi error kalau implement atau punya method Error()
*/
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v \n", e.code, e.message, e.err)
}

// function that return a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong!!",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong!",
			err:     err,
		}
	}

	return nil
}

func doSomethingElse() error {
	return errors.New("INTERNAL_SERVER_ERROR")
}
