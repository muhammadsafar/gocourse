package main

import (
	"fmt"
	"net/http"
)

func main22() {
	fmt.Println("Hello, Standar Library!")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
