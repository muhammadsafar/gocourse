package main

func main5() {

	// 🟩 PascalCase
	// Contoh: MyVariable, GetUserName, CalculateTotal
	// Digunakan untuk:
	// - Struct
	// - Interface
	// - Exported functions (huruf besar di awal artinya bisa diakses dari luar package)
	// Contoh:
	// type UserProfile struct { Name string }
	// func GetUserName(u UserProfile) string { return u.Name }

	// 🟧 camelCase
	// Contoh: myVariable, getUserName, calculateTotal
	// Digunakan untuk:
	// - Variabel lokal
	// - Fungsi internal (tidak diexport)
	// Contoh:
	// func calculateTotal(price int, qty int) int { return price * qty }

	// 🟦 snake_case
	// Contoh: my_variable, get_user_name, calculate_total
	// Digunakan untuk:
	// - Nama tabel atau kolom di database
	// - Nama file
	// Contoh (SQL):
	// CREATE TABLE user_accounts (user_id SERIAL PRIMARY KEY, user_name VARCHAR(50));

	// 🟥 UPPER_SNAKE_CASE (Screaming Snake Case)
	// Contoh: MY_CONSTANT, API_KEY, MAX_CONNECTIONS
	// Digunakan untuk:
	// - Konstanta
	// - Environment variable
	// Contoh:
	// const API_KEY = "12345-ABCDE"
	// const MAX_USERS = 100

	// 🟨 MixedCase
	// Contoh: htmlDocument, javaScript, eBayAPI, iPhoneModel
	// Digunakan untuk:
	// - Nama teknologi atau istilah khusus (brand name, dll)
	// Contoh (JavaScript):
	// let htmlDocument = document;
	// let javaScriptVersion = "ES6";

}
