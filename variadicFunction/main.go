package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	printEachLine("abc", "def", 123, "ghi")

	x := []interface{}{"ABD", "DEF", 123, "GHI"}
	printEachLine(x...)
}
