package main

import (
	"fmt"
	"unicode/utf8"
)

func main60() {

	message := "Hello, \nGo!"
	message2 := "Hello, \tGo!"
	message3 := "Hello, \rGo!" //Go!lo Go!->replace Hel

	rawMsg := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)

	fmt.Println(rawMsg)

	fmt.Println("length of message variable is : ", len(rawMsg))

	fmt.Println(`message 1st character>>`, message[0]) // ASCII

	greeting := "Hello "
	name := "Abdullah"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value of 65
	str1a := "apple" // A has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97

	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str1a > str1)

	for _, char := range message {
		//fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count : ", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	ach := 'ش'

	fmt.Println(ch)
	fmt.Println(ach)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", ach)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const SALAM = "سلام"
	fmt.Println(SALAM)

	ahello := "السلام عليكم"

	for _, runeVal := range ahello {
		fmt.Printf("%c\n", runeVal)
	}

	//rune dipakai untuk karakter khsuus misal tulisan arab,jepang, emoji dll, karena misal emoji sepanjang 4 bit, jadi

	/**
		Karena tidak semua karakter bisa disimpan dalam 1 byte.
	Contoh: emoji, huruf Mandarin, Arab, Jepang, aksen, dll.

	"😊" sebagai string → bytes: F0 9F 98 8A

	'😊' sebagai rune → kode Unicode: 128522

	Jadi rune membantu kamu memanipulasi karakter, bukan byte.
	*/

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)

}
