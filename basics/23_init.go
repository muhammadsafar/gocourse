package main

import "fmt"

//init running saat pertama kali sebelum main, tidak terima param dan tidak return value
/**
Kapan SEBAIKNYA init() digunakan?

Gunakan init() hanya untuk hal-hal penting seperti:

Load environment variable

Register database driver

Register router

Inisialisasi konfigurasi

*/

func init() {
	fmt.Println("Initializing package1..")
}

func init() {
	fmt.Println("Initializing package2..")
}

func init() {
	fmt.Println("Initializing package3..")
}

func main1() {

	fmt.Println("Inside the main function..")
}
