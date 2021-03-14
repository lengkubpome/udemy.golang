package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 24}
	// fmt.Printf("Hi I'm %s. %d years old \n", filicity.Name, filicity.Age)
	// fmt.Printf("Hi I'm %s. %d years old \n", oliver.Name, oliver.Age)
	const greetPerson = `Hi I'm {{.Name}}. {{.Age}} years old {{"\n"}}`

	// greetTemplate, err := template.New("greetingFromPerson").Parse(greetPerson)
	// if err != nil {
	// 	return
	// }

	// ใช้ template.Must ช่วยทำให้โค้ดดูสั้นลง
	greetTemplate := template.Must(template.New("greetingFromPerson").Parse(greetPerson))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)
}
