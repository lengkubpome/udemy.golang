package hr

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	// person      Person
	Person
	Designation string
}
