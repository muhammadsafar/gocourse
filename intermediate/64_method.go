package main

import "fmt"

type Shape struct {
	Rectangel //anonymous field
}

type Rectangel struct {
	p float64
	l float64
}

//predefined multiple return
func (r Rectangel) hitungRectangel() (luas float64, keliling float64) {
	luas = r.p * r.l
	keliling = 2 * (r.p + r.l)
	return
}

//method with pointer receiver
func (r *Rectangel) Scale(factor float64) {
	r.p *= factor
	r.l *= factor
}

func main64() {

	var pp = Rectangel{
		p: 10,
		l: 2,
	}
	luas, kel := pp.hitungRectangel()
	fmt.Printf("Luas Persegi Panjang : %.0f, Keliling : %.0f \n", luas, kel)

	//modify with pointer
	pp.Scale(2)
	luas, kel = pp.hitungRectangel()
	fmt.Printf("*Setelah pointer (edit by memory address) Luas Persegi Panjang: %.0f, Keliling : %.0f", luas, kel)

	num := MyInt(-5)
	fmt.Printf("Number is Positive ? >> %v\n", num.isPositive())

	fmt.Println(num.welcomeMsg())

	s := Shape{
		Rectangel: Rectangel{
			p: 5,
			l: 3,
		},
	}

	fmt.Println(s.hitungRectangel())
}

type MyInt int

//method on a user define type
func (m MyInt) isPositive() bool {
	return m > 0
}

//not create instance karena tidak butuh modify value
func (MyInt) welcomeMsg() string {
	return "Welcome to myInt without instance"
}
