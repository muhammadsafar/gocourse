package main

import (
	"fmt"
)

func main() {

	p := Person{
		firstName: "Muhammad",
		lastName:  "Baharuddin",
		age:       30,
		gender:    "Male",
		address: Address{
			city:    "Balikpapan",
			country: "Indonesia",
			zip:     76143,
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "7097475-3982",
			cell: "08518751423",
		},
	}

	p2 := Person{
		firstName: "Nurul",
		lastName:  "Huda",
		gender:    "Female",
	}

	p2.address.city = "PPU"
	p2.address.zip = 762412

	fmt.Println("Abang >> ", p)
	fmt.Println("Adek >> ", p2)

	//Anonymous
	u1 := struct {
		usename string
		email   string
	}{
		usename: "abdullah",
		email:   "abdullah@gmail.com",
	}

	fmt.Println("Anak >>", u1)

	ms := p.fullName()
	fmt.Println("p fullname >>", ms)

	p2.incrementAgeByOne()
	fmt.Println("Adek >>", p2)

}

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
	zip     int
}

// Value receiver: cukup karena hanya membaca data,
// tidak mengubah isi struct.
func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// Pointer receiver: digunakan karena ingin mengubah
// isi struct (field age).
func (p *Person) incrementAgeByOne() {
	p.age++
}
