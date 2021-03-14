package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 24}
	// fmt.Printf("Hi I'm %s. %d years old \n", filicity.Name, filicity.Age)
	// fmt.Printf("Hi I'm %s. %d years old \n", oliver.Name, oliver.Age)
	const greetPerson = `Hi I'm {{.Name | upperCase}}. {{.Age}} years old {{"\n"}}`
	const greetPeople = `{{range .}}Hi I'm {{.Name}}. {{.Age}} years old {{"\n"}}{{end}}`

	people := []Person{filicity, oliver}

	// สร้าง Function map
	maps := make(template.FuncMap)
	maps["upperCase"] = upperCase

	// greetTemplate, err := template.New("greetingFromPerson").Parse(greetPerson)
	// if err != nil {
	// 	return
	// }

	// ใช้ template.Must ช่วยทำให้โค้ดดูสั้นลง
	greetTemplate := template.
		Must(template.New("greetingFromPerson").
			// Funcs(template.FuncMap{"upperCase": upperCase}).
			Funcs(maps).
			Parse(greetPerson))
	greetPeopleTemplate := template.Must(template.New("greetPeopleTemplate").Parse(greetPeople))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)
	fmt.Println("========================================")
	greetPeopleTemplate.Execute(os.Stdout, people)
}
