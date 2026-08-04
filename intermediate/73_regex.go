package main

import (
	"fmt"
	"regexp"
)

func main73() {

	fmt.Println("he said, \n\"I am great\"")
	fmt.Println(`he said, "I am great"`)

	//compile a regex pattern to match emailaddress
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z{2,}]`)

	email := "muhammad@gmail.com"
	email2 := "invalid_email"

	//match
	fmt.Println("Email : ", re.MatchString(email))
	fmt.Println("Email 2 : ", re.MatchString(email2))

	//capturing group
	//compile a regex pattern to capture date components

	re2 := regexp.MustCompile(`(\d){4}-(\d{2}-(\d{2}))`)
	date := "2025-09-12"
	submatch := re2.FindStringSubmatch(date)
	fmt.Println(submatch)
	fmt.Println("date >>" + submatch[0])
	fmt.Println("date >>" + submatch[1])
	fmt.Println("date >>" + submatch[2])
	fmt.Println("date >>" + submatch[3])

	//target string

	str := "Hello Golang"

	re = regexp.MustCompile(`[aiueo]`)
	result := re.ReplaceAllString(str, "*")

	fmt.Println("replace >>" + result)

	//i - case insensitive
	//m - multi line model
	//s - dot mathes all

	re = regexp.MustCompile(`(?i)go`)

	//Test string
	text := "Golang is going great"

	//match
	fmt.Println("Match >>", re.MatchString(text))

}
