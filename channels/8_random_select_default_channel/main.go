package main

import (
	"fmt"
	"time"
)

func running(name string, ch <-chan int) {
	for {
		fmt.Println(name, <-ch)
	}

}

func main() {

	//วิธีที่ 1 ใช้ Channel
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// ch3 := make(chan int)
	// ch4 := make(chan int)

	// go running("Ch1", ch1)
	// go running("Ch2", ch2)
	// go running("Ch3", ch3)
	// go running("Ch4", ch4)

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case ch1 <- i:
	// 	case ch2 <- i:
	// 	case ch3 <- i:
	// 	case ch4 <- i:
	// 	default:
	// 		fmt.Println("Default")
	// 	}
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// วิธีที่ 2 ใช้วิธี time ticker
	// tricker1 := time.NewTicker(1 * time.Second)
	// tricker2 := time.NewTicker(2 * time.Second)
	// tricker3 := time.NewTicker(3 * time.Second)
	// tricker4 := time.NewTicker(4 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case v := <-tricker1.C:
	// 		fmt.Println("ch1", v.Nanosecond())
	// 	case v := <-tricker2.C:
	// 		fmt.Println("ch2", v.Nanosecond())
	// 	case v := <-tricker3.C:
	// 		fmt.Println("ch3", v.Nanosecond())
	// 	case v := <-tricker4.C:
	// 		fmt.Println("ch4", v.Nanosecond())
	// 		// default: //ถ้าไม่มี default ระบบจะรอจนกว่าถึงเวลา <-
	// 		// 	fmt.Println("Default")
	// 	}
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// วิธีที่ 3
	after2Second := time.After(2 * time.Second)

	for i := 0; i < 10; i++ {
		select {
		case v := <-after2Second:
			fmt.Println("Got value from after2Second channel", v.Second())
		default:
			fmt.Println("Default")
		}
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")
}
