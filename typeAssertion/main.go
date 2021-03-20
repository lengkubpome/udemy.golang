package main

import (
	"fmt"
)

// คือการตรวจเช็คชนิดของ Type และ Interface
// x.(T)

func testValue(x interface{}) {
	switch v := x.(type) {

	case string:
		fmt.Println("string value : ", v)
	case int:
		fmt.Println("int value : ", v)
	case float32, float64:
		fmt.Println("float value : ", v)
	default:
		fmt.Println("No match any")
	}

}

func main() {
	testValue("abc")
	testValue(123)
	testValue(true)
	testValue(1234.23)
	testValue(float32(1234.23))
}
