package main

import (
	"fmt"

	"example.com/hr"
)

func main() {
	// x := struct {
	// 	name string
	// 	age  int
	// }{age: 23, name: "Filicity"}

	x := hr.Person{Name: "Filicity", Age: 23}

	y := x
	y.Name = "Oliver"

	z := &x
	z.Name = "Ola" //(*z).name = "Ola"

	fmt.Println(x)
	// fmt.Println(x.name)
	// fmt.Println(x.age)

	fmt.Println(y)

	// a := &Person{"Khonkaen", 12}
	a := new(hr.Person)
	*a = hr.Person{Name: "Khonkaen", Age: 12}
	fmt.Printf("%+v\n", *a)
	fmt.Printf("%T\n\n", a)

	filicity := hr.Employee{
		// person:      Person{"Filicity", 29},
		// designation: "Programmer",

		// ทำ Embedding
		Person:      hr.Person{Name: "Filicity", Age: 29},
		Designation: "Programmer",
	}
	fmt.Printf("%+v\n", filicity)
	// fmt.Printf(filicity.person.name) // มันยาวไป
	fmt.Printf("%v\n", filicity.Name) // เลยเปลี่ยนมาใช้ Embedding
	fmt.Printf("%v\n", filicity.Age)  // เลยเปลี่ยนมาใช้ Embedding
}
