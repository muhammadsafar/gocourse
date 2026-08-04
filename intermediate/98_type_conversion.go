package main

import "fmt"

func main98() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool(true)

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)

	g := "Hello Goo "
	// var h []byte
	h := []byte(g)
	fmt.Println(h)

	i := []byte{255, 120, 72}
	j := string(i)
	fmt.Println(j)

}
