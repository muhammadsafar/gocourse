package main

import (
	"fmt"
	"maps"
)

func main3Map() {

	// var mapva map[keyType] valueType -> baru deklarasi, belum inisialisasi

	// mapVariable = make(map[keyType] valueType)

	//using a Map Literasi
	// mapVar = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int) //-> inisial agar bisa tambahkan data

	myMap["age"] = 12
	myMap["score"] = 97
	myMap["key1"] = 12

	fmt.Println("map >>", myMap)
	fmt.Println("map score>>", myMap["score"])

	fmt.Println("undefine key>>", myMap["nokey"])
	myMap["score"] = 94
	fmt.Println("map score update>>", myMap["score"])

	//delete key on map
	delete(myMap, "age")
	fmt.Println("map update after delete>>", myMap)

	//clear
	// clear(myMap)
	// fmt.Println("map clear>>", myMap)

	v, isAssociate := myMap["key1"]
	fmt.Println(v)
	fmt.Println("is associate>>", isAssociate)

	//test map2

	myMap2 := map[string]int{"a": 1, "b": 20}
	myMap3 := map[string]int{"a": 1, "b": 20}
	fmt.Println("myMap2>>", myMap2)

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("my map 3 and my map 2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Printf("key- %v  >> value is %v\n", k, v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("the map is init to nil")
	}

	myMap4 = make(map[string]string) // <-- wajib

	myMap4["key1"] = "value1"
	fmt.Println("length map >>", len(myMap4))

	//map dalam map

	myMap5 := make(map[string]map[string]string)
	myMap5["myMap4"] = myMap4

	fmt.Println("map5 >>", myMap5)
}
