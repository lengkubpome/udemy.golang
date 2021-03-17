package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 3.
// func intervene() func() {
// 	fmt.Println("1")
// 	return func() {
// 		fmt.Println("2")
// 	}
// }

// 4.
// func loadingImage() {
// 	defer stopWatch("loadingImage")()
// 	time.Sleep(3 * time.Millisecond)
// 	fmt.Println("loadingImage : done")
// }

// func stopWatch(name string) func() {
// 	start := time.Now()
// 	return func() {
// 		fmt.Printf("function %s : tooks time %vs\n", name, time.Now().Sub(start).Seconds())
// 	}
// }

// 5.
// func add(a, b int) (c int) {
// 	defer func() {
// 		fmt.Printf("c is %d\n", c)
// 	}()
// 	return a + b
// }

func main() {
	// 1.
	// defer fmt.Println("a")
	// fmt.Println("b")
	// fmt.Println("c")

	// 2.
	// x := 1
	// addX := func(a int) int {
	// 	x = x + a
	// 	return x
	// }
	// defer fmt.Println(addX(4))
	// fmt.Println(x)

	// 3.
	// defer intervene()()
	// fmt.Println("****")

	// 4.
	// loadingImage()

	// 5.
	// fmt.Println(add(2, 3))

	// 6.
	resp, err := http.Get("https://golang.org")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
