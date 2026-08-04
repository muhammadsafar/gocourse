package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person3 struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	//City    string   `xml:"city,omitempty"`
	// Email   string   `xml:"email"`
	Email string `xml:"-"`

	Address Address3 `xml:"address"`
}

// kalau "-" artinya betul2 diabaikan atai tidak tampil
// kalau omitempty : tampilkan jika data tidak nil 0 atau ""

type Address3 struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main96() {

	person := Person3{
		Name: "Abdullah",
		Age:  21,
		// City:  "PPU",
		Email: "mail.com",
		Address: Address3{
			City:  "Balikpapan",
			State: "D12",
		},
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling data into XML", err)
	}

	fmt.Println("xml >>\n", string(xmlData))

	//indent
	xmlData, err = xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML", err)
	}

	fmt.Println("xml indent>>\n", string(xmlData))

	//unmarshalling
	xmlRaw := `<person><name>John</name><age>23</age><address><city>ppu</city><state>12SS</state></address></person>`

	var personxml Person3

	err = xml.Unmarshal([]byte(xmlRaw), &personxml)

	if err != nil {
		log.Fatalln("Error Unmarshalling Person", err)
	}

	fmt.Println("Object Person >>", personxml)
	fmt.Println("Object Person age + 30>>", personxml.Age+30)

	fmt.Println(personxml)
	fmt.Println("local string : ", personxml.XMLName.Local)
	fmt.Println("namespace : ", personxml.XMLName.Space)

	book := Book{
		ISBN:       "9863-275248-7550-97234",
		Title:      "Titanic",
		Author:     "Leo",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	dataXml, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling data", err)
	}

	fmt.Println("data>>", string(dataXml))

}

// attribute
// <book isbn="9863-275248-7550-97234" title="Titanic" author="Leo"></book>
type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}
