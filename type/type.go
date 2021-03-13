package main

import "fmt"

type KG float64
type LB float64

func (kg KG) toLB() LB {
	return LB(kg / 0.453592)
}

func (lb LB) toKG() KG {
	return KG(lb / 2.20462)
}

func main() {
	k := KG(1)
	b := LB(2.20462)

	fmt.Println(k == b.toKG())

	x := 5
	if v := x + 1; (x <= 5) && false {
		fmt.Println("Hi, I'm in if", v)
	} else {
		fmt.Println("Hi, I'm in else", v)
	}
}
