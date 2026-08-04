package main

import (
	"fmt"
	"math"
)

// methode dalam interface harus di implement semua oleh method, jika tidak maka tidak termasuk implement
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.diameter()
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main65() {

	r := rect{height: 2, width: 6}
	c := circle{radius: 5}

	//r2 := rect2{height: 4, width: 8}

	measure(r)
	measure(c)
	//measure(r2)

	myPrinter("Musa", 24.5, true, 89)

	printType("Musa")
	printType(120)
	printType(true)
	printType(12.5)

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is boolean")
	default:
		fmt.Println("Type is unknown")
	}
}

func myPrinter(i ...interface{}) {
	for _, val := range i {
		fmt.Println(val)
	}
}

// type rect2 struct {
// 	width, height float64
// }

// func (r2 rect2) area() float64 {
// 	return r2.height * r2.width
// }

// func (r2 rect2) perim() float64 {
// 	return 2 * (r2.height + r2.width)
// }
