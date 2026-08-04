package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main70() {

	str := "Hello Go!"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "Coding"
	result := str1 + " " + str2

	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[:7]) //slicing get index dari awal sampai ke-7 , [2:] = dari index ke2 smapai akhir, [2:5] = dari index 2 sampai ke 5

	num := 16
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))
	fmt.Println(len(str3))

	// string splitting

	fruits := "apple, orange, banana, lemon"
	parts := strings.Split(fruits, ",")

	fmt.Println("parts of fruit >>", parts[0])

	countries := []string{"Indonesia", "Malaysia", "Brunei", "Saudi arabia"}

	joined := strings.Join(countries, ", ")

	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Coding"))

	replaced := strings.Replace(str, "Go", "World", 1) // 1 = berapa kali ganti dalam isi string
	fmt.Println(replaced)

	strwspace := " Hello Everyone"
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo", 3))
	fmt.Println(strings.Count("Hello", "e"))
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str5 := "Hello, 123 Go! 11"
	re := regexp.MustCompile(`\d+`)
	matcher := re.FindAllString(str5, -1)

	fmt.Println(matcher)

	str6 := "Hello, こんにちは"
	fmt.Println(utf8.RuneCountInString(str6))

	//STRING BUILDER

	var strBuilder strings.Builder

	strBuilder.WriteString("Hi, ")
	strBuilder.WriteString("my ")
	strBuilder.WriteString("name ")
	strBuilder.WriteString("is ")
	strBuilder.WriteString("Abdullah ")

	//convert builder to strings
	res := strBuilder.String()
	fmt.Println(res)

	//using writeRune

	strBuilder.WriteRune(' ')
	strBuilder.WriteString("How are you")

	res = strBuilder.String()
	fmt.Println(res)

	//builder juga bisa direset
	strBuilder.Reset()
	strBuilder.WriteString("Starting fresh...")

	res = strBuilder.String()
	fmt.Println(res)

}
