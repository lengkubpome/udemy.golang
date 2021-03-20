package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byName []*Person

func (p byName) Len() int {
	return len(p)
}

func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Custom Sort
type customSort struct {
	Persons []*Person
	less    func(i, j *Person) bool
}

func (p customSort) Len() int {
	return len(p.Persons)
}

func (p customSort) Less(i, j int) bool {
	return p.less(p.Persons[i], p.Persons[j])
}

func (p customSort) Swap(i, j int) {
	p.Persons[i], p.Persons[j] = p.Persons[j], p.Persons[i]
}

func main() {
	p := []*Person{
		{"B", 20},
		{"A", 20},
		{"I", 21},
		{"J", 23},
		{"C", 29},
		{"D", 24},
		{"H", 29},
		{"E", 22},
		{"A", 22},
		{"D", 21},
	}

	printPerson(p)
	sort.Sort(byName(p))
	printPerson(p)

	// Custom Sort
	sort.Sort(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.name != j.name {
			return i.name < j.name
		}
		if i.age != j.age {
			return i.age < i.age
		}
		return false
	}})
	printPerson(p)

	// Reverse Sort
	sort.Sort(sort.Reverse(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.name != j.name {
			return i.name < j.name
		}
		if i.age != j.age {
			return i.age < i.age
		}
		return false
	}}))
	printPerson(p)
}

func printPerson(p []*Person) {
	fmt.Println("-------------")
	for _, v := range p {
		fmt.Println(*v)
	}
}
