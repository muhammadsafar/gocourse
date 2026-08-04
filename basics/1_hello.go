package main

import (
	"fmt"
	"gocourse/library"
	"strings"
)

func mainx() {

	fmt.Println("\n==========================")
	fmt.Println("Hello, World!")

	var an = library.Animal{
		Name:  "Cat",
		Sound: "Meow",
	}

	fmt.Println("\n=========================")
	fmt.Printf("Animal Name: %s, Sound: %s\n", an.Name, an.Sound)

	animaslSound := library.GetAnimalSound(an)
	fmt.Printf("Animal Sound from function: %s\n", animaslSound)

	fmt.Println("\n=========================")
	var arr = []string{"apple", "banana", "cherry"}

	for _, fruit := range arr {
		if strings.Contains(fruit, "a") {
			fmt.Println(fruit)
		}
	}

	fmt.Println("\n=======================")
	var familyNames = library.PrintFamily("Baharuddin", "Ramlah", "Nurul", "Aisyah", "Muhammad")
	fmt.Println("Family Names:", familyNames)

	fmt.Println("\n=======================")
	// Menggunakan method milik struct PersegiPanjang dari package library
	pp := library.PersegiPanjang{
		Panjang: 10,
		Lebar:   5,
	}

	luas := pp.Luas()
	keliling := pp.Keliling()
	fmt.Printf("Persegi Panjang - Panjang: %d, Lebar: %d, Luas: %d, Keliling: %d\n", pp.Panjang, pp.Lebar, luas, keliling)

	fmt.Println("\n=======================")
	sisi := 7
	luasPersegi, kelilingPersegi := library.HitungPersegi(sisi)
	fmt.Printf("Predefined persegi - Sisi: %d, Luas: %d, Keliling: %d\n", sisi, luasPersegi, kelilingPersegi)

	luasPersegi2, kelilingPersegi2 := library.HitungPersegi2(sisi)
	fmt.Printf("explicit return ersegi - Sisi: %d, Luas: %d, Keliling: %d\n", sisi, luasPersegi2, kelilingPersegi2)

	//=====STANDAR LIBRARY=====//
	fmt.Println("\n=======================")
	var str = "Hello, Go Standard Library!"
	fmt.Println("Original String:", str)
	fmt.Println("Uppercase:", strings.ToUpper(str))
	fmt.Println("Lowercase:", strings.ToLower(str))
	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))
	fmt.Println("Replace 'Go' with 'Golang':", strings.ReplaceAll(str, "Go", "Golang"))

}
