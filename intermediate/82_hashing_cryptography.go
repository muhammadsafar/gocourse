package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main82() {

	//HASHING
	pass := "password123"
	hash := sha256.Sum256([]byte(pass))
	hash512 := sha512.Sum512([]byte(pass))

	fmt.Println("pass is:", pass)
	fmt.Println("hash:", hash)
	fmt.Printf("SHA256 Hash hex val is %x \n", hash)
	fmt.Printf("SHA512 Hash hex val is %x \n", hash512)

	//SALTING

	pass2 := "password123"
	salt, err := generateSaltedHash()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	saltStr := base64.StdEncoding.EncodeToString(salt)
	signUpHash := hashPasswordWithSalt(pass2, salt)

	fmt.Println("\npass2 is:", pass2)

	originalPassWithoutSalt := sha256.Sum256([]byte(pass2))
	fmt.Println("\nOriginal pass hash (no salt) hex:", fmt.Sprintf("%x", originalPassWithoutSalt))

	fmt.Println("salt (base64):", saltStr)
	fmt.Println("salted hash (base64):", signUpHash)

	//retrieved the salted string and decoded it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}

	loginHash := hashPasswordWithSalt(pass2, decodedSalt)

	//compare the stored singupHash with loginhash
	if loginHash == signUpHash {
		fmt.Println("Password verified successfully!")
	} else {
		fmt.Println("Invalid password!")
	}

}

func generateSaltedHash() ([]byte, error) {
	salt := make([]byte, 16)

	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		return nil, err
	}

	return salt, nil

}

// func to hash password with salt
func hashPasswordWithSalt(password string, salt []byte) string {

	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
