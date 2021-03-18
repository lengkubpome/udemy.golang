package main

import "fmt"

type Person struct {
	Name    string
	Surname string
}

func (p *Person) FullName() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Id     string
	Office string
	Person
}

func (e *Employee) Detail() string {
	return "ID : " + e.Id + ". Office : " + e.Office + ". Fullname : " + e.Person.FullName()
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	Language []string
}

func (p *Programmer) Detail() string {
	return "Programmer : " + p.Employee.Detail()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detail() string {

	return "Tester : " + t.Employee.Detail()
}

func main() {
	david := Person{
		Name:    "David",
		Surname: "Light",
	}

	empDavid := Employee{
		Person: david,
		Id:     "1234",
		Office: "Thailand",
	}

	progDavid := Programmer{
		Employee: empDavid,
		Language: []string{"Golang", "Java"},
	}

	fmt.Printf("%+v\n", progDavid)

	oliver := Person{
		Name:    "Oliver",
		Surname: "Queen",
	}

	empOliver := Employee{
		Person: oliver,
		Id:     "456",
		Office: "Thailand",
	}

	testerOliver := Tester{
		Employee: empOliver,
		Tools:    []string{"Robot"},
	}

	fmt.Printf("%+v\n", testerOliver)

	fmt.Println(empDavid.IsSameOffice(&empOliver))
	fmt.Println(progDavid.IsSameOffice(&(testerOliver.Employee)))
	fmt.Println(progDavid.Detail())

	davidDetail := progDavid.Detail
	fmt.Println("davidDetail : ", davidDetail())

	isSameOffice := (*Employee).IsSameOffice
	fmt.Println(isSameOffice(&empDavid, &empOliver))
}
