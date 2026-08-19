package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool
type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func mainsort() {
	numbers := []int{5, 3, 4, 1, 2}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers : ", numbers)

	// stringSlice := []string{"Muhammad", "Abdullah", "Nusaiba", "Baharuddin"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted String : ", stringSlice)

	//custom practice
	people := []Person{
		{"Musa", 30},
		{"Aming", 21},
		{"Tahang", 22},
		{"Aco", 34},
		{"Funding", 35},
	}

	fmt.Println("people : ", people)
	// sort.Sort(ByAge(people))
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(age).Sort(people)
	fmt.Println("Sorted people by age : ", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sorted people by age desc : ", people)
	By(name).Sort(people)
	fmt.Println("Sorted people by name : ", people)
	By(lenName).Sort(people)
	fmt.Println("Sorted people by length name : ", people)

	//==============SORT.SLICE
	stringSlice := []string{"banana", "apple", "watermelon", "cherry"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last char >>", stringSlice)
}

//output
/**
Sorted numbers :  [1 2 3 4 5]
people :  [{Musa 30} {Aming 21} {Tahang 22} {Aco 34} {Funding 35}]
Sorted people by age :  [{Aming 21} {Tahang 22} {Musa 30} {Aco 34} {Funding 35}]
Sorted people by age desc :  [{Funding 35} {Aco 34} {Musa 30} {Tahang 22} {Aming 21}]
Sorted people by name :  [{Aco 34} {Aming 21} {Funding 35} {Musa 30} {Tahang 22}]
Sorted people by length name :  [{Aco 34} {Musa 30} {Aming 21} {Tahang 22} {Funding 35}]
Sorted by last char >> [banana apple watermelon cherry]
*/
