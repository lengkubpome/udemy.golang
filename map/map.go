package main

import "fmt"

func main() {
	items := map[string]int{
		"pen":    3,
		"pencil": 5,
		"a":      0,
		"b":      0,
		"c":      0,
	}

	orders := make(map[string]int)
	orders["pen"] = 16
	orders["pencil"] = 14

	x := items
	x["ruler"] = 7
	x = orders
	fmt.Println("items", items)
	fmt.Println("x", x)
	fmt.Println("items[\"pencil\"]", items["pencil"])
	fmt.Println("orders\n\n", orders)

	v, ok := items["pencil"]
	if !ok {
		fmt.Println("not Exit", v)
	} else {
		fmt.Println("Exit", v)
	}

	for k, v := range items {
		fmt.Println(k, v)
	}
}
