package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person2 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age"`
}

func main95() {

	p1 := Person2{
		FirstName: "Musa",
		LastName:  "",
		Age:       21,
	}

	jd1, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Error masrhalling struct", err)
	}

	fmt.Println(string(jd1))

}
