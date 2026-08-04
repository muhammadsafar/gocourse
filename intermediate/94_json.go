package main

import (
	"encoding/json"
	"fmt"
)

type Persons struct {
	FirstName string   `json:"first_name"`
	Age       int      `json:"age,omitempty"`
	Email     string   `json:"email,omitempty"`
	Address   Addresss `json:"address"`
}

type Addresss struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main94() {

	person := Persons{
		FirstName: "Musa",
	}

	//marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling json.", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	per1 := Persons{
		FirstName: "Aisha",
		Age:       25,
		Email:     "aisyah@gmail.com",
		Address: Addresss{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	jsonData2, err := json.Marshal(per1)
	if err != nil {
		fmt.Println("Error marshalling json.", err)
		return
	}
	fmt.Println("JSON Data 2:", string(jsonData2))

	jsonData1 := `{"first_name":"Aisha","age":25,"email":"aisyah@gmail.com","address":{"city":"Jakarta","country":"Indonesia"}}`

	var employeeObj Employees

	//unmarshal json -> object
	err = json.Unmarshal([]byte(jsonData1), &employeeObj)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
		return
	}

	fmt.Println("employee obj>>", employeeObj)
	fmt.Println("Musa age increased by 5 years>>", employeeObj.Age+5)
	fmt.Println("musas city>>", employeeObj.Address.City)

	listOfCities := []Addresss{
		{City: "Penajam Paser Utara", Country: "ID"},
		{City: "Kaboul", Country: "IR"},
		{City: "Tokyo", Country: "JP"},
		{City: "Barcelona", Country: "SP"},
	}

	fmt.Println(listOfCities)
	jsonList, err := json.Marshal(listOfCities)
	if err != nil {
		fmt.Println("Error marshalling json list", err)
		return
	}

	fmt.Println("Json List cities >>", string(jsonList))

	//jsonData2 := `{"name":"Abdullah","age":20,"address":{"city":"Berau","country":"ID"}}`

	var data map[string]interface{}

	json.Unmarshal([]byte(jsonData2), data)

}

type Employees struct {
	FullName string   `json:"full_name"`
	EmplID   string   `json:"emp_id"`
	Age      int      `json:"age"`
	Address  Addresss `json:"address"`
}
