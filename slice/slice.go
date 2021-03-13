package main

import "fmt"

func main() {
	// slice from array
	fruits := [5]string{"apple", "banana", "papaya", "orange", "mango"}
	myFavor := fruits[1:4]
	fmt.Println(len(myFavor), cap(myFavor), myFavor)

	yourFavor := myFavor
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor)

	yourFavor[0] = "guava"
	fmt.Println(len(myFavor), cap(myFavor), myFavor)
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor)
	fmt.Println("\n")

	// using build-in function -> make
	s := make([]int, 5, 10)
	fmt.Printf("len : %d, cap: %d \n", len(s), cap(s))

	// using build-in function -> append
	x := [10]int{}
	a := x[0:5]
	b := x[5:7]
	for i := range a {
		a[i] = i * i
	}
	b[0] = 98
	b[1] = 99

	a = append(a, b...)
	a[len(a)-1] = 25
	a = append(a, 71, 72, 73, 74)
	fmt.Println("x :\t", x)
	fmt.Println("a :\t", len(a), cap(a), a)
	fmt.Println("b :\t", len(b), cap(b), b)

	sliceX := append([]byte("helll "), "world"...)
	fmt.Printf("%s", sliceX)

}
