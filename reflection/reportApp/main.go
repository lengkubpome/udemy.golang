package main

import (
	"fmt"
	"reflect"

	"example.com/reportApp/report"
)

type Person struct {
	name string `report:"ชื่อ,uppercase"`
	age  int    `report:"อายุ"`
}

type Employee struct {
	name string `report:"ชื่อ"`
	age  int    `report:"อายุ"`
}

type Person2 struct {
	Name string `report:"ชื่อ,uppercase"`
	Age  int    `report:"อายุ"`
}

func main() {
	//	Case1
	// // Person( name : David, age : 23 )
	// fmt.Println(report.Text(Employee{name: "Emily", age: 24}))
	// //	Employee( name : Emily, age : 24 )
	// fmt.Println(report.Text(struct {
	// 	name string `report:",uppercase"`
	// 	age  int    `report:"อายุ"`
	// }{name: "Emily3", age: 24}))
	// // anonymous( name : Emily3, age : 24 )

	// Case 2
	person := Person2{Name: "David", Age: 23}
	fmt.Println("Before : ", person) // {David 23}
	fmt.Println(report.Text(person))
	v := reflect.ValueOf(&person) // <- use address
	v.Elem().Field(0).SetString("DAVID")
	v.Elem().Field(1).SetInt(v.Elem().Field(1).Int() * 2)
	fmt.Println("After : ", person) // {DAVID 46}
}
