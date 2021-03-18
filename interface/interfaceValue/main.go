package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// รู้หรือไม่ Interface มี value และ type อยู่ในตัวของมันเอง
	var w io.Writer // มันคือ Interface
	fmt.Printf("%T , %v\n", w, w)

	w = os.Stdout
	fmt.Printf("%T , %v\n", w, w)

	w = &bytes.Buffer{}
	// w.Write([]byte("hello"))
	fmt.Printf("%T , %v\n", w, w)

	w = nil
	fmt.Printf("%T , %v\n", w, w)

	var byteBuff *bytes.Buffer
	fmt.Println(byteBuff == nil)
	fmt.Println(w == nil)
	w = byteBuff
	fmt.Println(w == nil) // เป็น false เพราะว่า type มันไม่เป็น nil ถึงแม้ว่า value จะเป็น nil ก็ตาม
	fmt.Printf("%T , %v\n", w, w)

	var byteBuff2 *bytes.Buffer
	printBuff(byteBuff2) // panic ทันที
}

func printBuff(w io.Writer) {
	if w != nil {
		fmt.Println(w.Write([]byte("Hello")))
	}
}
