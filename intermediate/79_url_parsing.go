package main

import (
	"fmt"
	"net/url"
)

func main79() {

	//[scheme:][//[userinfo@]host[:port]][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path/to/resource?search=golang#section1"

	parseUrl, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Schema>>", parseUrl.Scheme)
	fmt.Println("Host>>", parseUrl.Host)
	fmt.Println("Port>>", parseUrl.Port())
	fmt.Println("Path>>", parseUrl.Path)
	fmt.Println("RawQuery>>", parseUrl.RawQuery)
	fmt.Println("Fragment>>", parseUrl.Fragment)

	url2 := "https://muhammad.baharuddin.com/path?name=musa&age=20"

	parsed2, err := url.Parse(url2)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("\nSchema>>", parsed2.Scheme)
	fmt.Println("Host>>", parsed2.Host)
	fmt.Println("Port>>", parsed2.Port())
	fmt.Println("Path>>", parsed2.Path)
	fmt.Println("RawQuery>>", parsed2.RawQuery)
	fmt.Println("Fragment>>", parsed2.Fragment)

	qp := parsed2.Query()
	fmt.Println("Name>>", qp.Get("name"))
	fmt.Println("Age>>", qp.Get("age"))

	//building URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/api/v1/microservice",
	}

	query := baseUrl.Query()
	query.Set("version", "1.0")

	baseUrl.RawQuery = query.Encode()

	fmt.Println("\nBuilt URL:", baseUrl.String())

	val := url.Values{}
	val.Add("name", "musa")
	val.Add("age", "30")
	val.Add("place", "ppu")
	val.Add("contry", "id")

	//encode
	encodedQuery := val.Encode()
	fmt.Println("Encoded Query String>>", encodedQuery)

	baseUrl1 := "https://example.com/search"
	fullUrl := fmt.Sprintf("%s?%s", baseUrl1, encodedQuery)
	fmt.Println("Full URL with Query>>", fullUrl)
}
