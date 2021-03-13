package main

import (
	"fmt"
	"reflect"
)

func main() {
	fruits := [5]string{
		"apple",
		"banan",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits), fruits)

	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots\t", reflect.TypeOf(twoSlots))
	fmt.Println("Three slots\t", reflect.TypeOf(threeSlots))

	indexFruits := [...]string{
		4:  "apple",
		0:  "banan",
		2:  "papaya",
		34: "orange",
		1:  "mango",
	}
	fmt.Println(len(indexFruits), indexFruits)

	ptrFruits := &fruits
	ptrFruits[0] = "watermelon"
	fmt.Println(fruits)
	fmt.Println(ptrFruits)

	a := [2][2]int{{1, 2}, {4, 6}}
	fmt.Println(a)
}
