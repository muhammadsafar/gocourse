package main

import (
	"fmt"
	"reflect"
)

type Greater struct {
}

func (g Greater) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	g := Greater{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	var method reflect.Method

	fmt.Println("Type : ", t)

	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("method %d : %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)
	result := m.Call([]reflect.Value{reflect.ValueOf("Musa"), reflect.ValueOf("Abdullah")})
	fmt.Println("Greet result : ", result[0].String())
	//[] string {"Musa"}
	//[] type{type("blabla"),type("blabla")}

}

//=EXMAPLE 2 with STRUCT
// type person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	p := person{Name: "Musa", Age: 21}
// 	v := reflect.ValueOf(p)
// 	v1 := reflect.ValueOf(&p).Elem()

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v \n", i, v.Field(i))
// 	}
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Muhammad")
// 	} else {
// 		fmt.Println("cant set")
// 	}
// 	fmt.Println("Modified person >>", p)
// }

//EXAMPLE 1

// func main() {

// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()
// 	fmt.Println("value >>", x)
// 	fmt.Println("reflect value >>", v)
// 	fmt.Println("type  >>", t)
// 	fmt.Println("kind  >>", t.Kind())
// 	fmt.Println("is int  >>", t.Kind() == reflect.Int)
// 	fmt.Println("is string  >>", t.Kind() == reflect.String)
// 	fmt.Println("is zero  >>", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("v1 ", v1.Int())
// 	fmt.Println("v2 ", v2.Type())
// 	v1.SetInt(21)
// 	fmt.Println("v1 modif", v1.Int())

// 	var itf interface{} = "hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("v3 type : ", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value :", v3.String())
// 	}

// }
