package main

import (
	"encoding/base64"
	"fmt"
)

func main81() {

	data := []byte("Hello, World! base 64 encoding in Go.")

	//encoding
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("encode>>", encoded)

	//decoding
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}
	// []byte -> string
	fmt.Println("decode>>", string(decoded))

	// url safe encoding
	dataUnsafe := []byte("He~llo, Wo+rl/d! base 64 encoding in Go.")
	urlSafeEncoded := base64.URLEncoding.EncodeToString(dataUnsafe)
	fmt.Println("URL safe encode>>", urlSafeEncoded)

	//dec
	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("Error decoding base64 URL safe:", err)
		return
	}
	fmt.Println("URL safe decode>>", string(urlSafeDecoded))
}
