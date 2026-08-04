package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math error : square root of negative number")
	}

	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data..")
	}

	return nil
}

func main68() {

	// res, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(res)

	// res1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// fmt.Println(res1)

	// data := []byte{}
	// if err := process(data); err != nil {

	// err2 := process(data)
	// if err2 != nil {

	// 	fmt.Println("Error : ", err2)
	// 	return
	// }

	// fmt.Println("Data Processed Successfully")

	err3 := eprocess()
	if err3 != nil {
		fmt.Println("err3>>", err3)
		// return
	}

	err4 := readData()
	if err4 != nil {
		fmt.Println("err4>>", err4)
		return
	}

	fmt.Println("Data read succesfully")

}

type myError struct {
	message string
}

// struct myErrorjadi error karena implement Error() bawaan go , asal return string
func (m *myError) Error() string {
	return fmt.Sprintf("Error : %s", m.message)
}

func eprocess() error {
	return &myError{"Custome error message"}
}

func readData() error {

	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("confg error")
}
