package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	employeeInfo person //embedded struct
	// person //anonymous field
	emplId string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi I'm %s and I'm %d years old", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi I'm %s, my ID is %s , and i earn %.2f.\n", e.employeeInfo.name, e.emplId, e.salary)
}

func main66() {

	emp := Employee{
		employeeInfo: person{
			name: "Nurul",
			age:  29,
		},

		emplId: "E001",
		salary: 2300,
	}

	fmt.Println(emp.employeeInfo.name)
	fmt.Println(emp.employeeInfo.age)
	fmt.Println(emp.emplId)
	fmt.Println(emp.salary)
	emp.introduce()

}
