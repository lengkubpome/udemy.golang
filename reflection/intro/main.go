package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{ a int }

func main() {
	look(MyStruct{a: 3})
}

func look(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.TypeOf(v).Kind()) // struct
	fmt.Println(reflect.ValueOf(v))
}
