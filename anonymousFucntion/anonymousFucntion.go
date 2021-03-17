package main

import "fmt"

func createF(list []int) []func() {
	result := []func(){}
	for _, v := range list {
		// captureV := v // เรียกผ่าน address ของค่า เพราะค่าไม่เรียกผ่าน address ระบบมันจะไปเรียกแต่ตำแหน่งสุดท้ายของ Loop
		v := v // เรียกแบบนี้ก็ได้
		fmt.Println("inside createF", v)
		result = append(result, func() {
			fmt.Println(v)
		})
	}
	return result
}

func main() {
	// fList := []func(){}
	fList := createF([]int{108, 12, 39, 4, 5})
	for _, v := range fList {
		v()
	}
}
