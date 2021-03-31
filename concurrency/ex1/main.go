package main

import (
	"fmt"
	"sync"
)

// var balance = 0
var chBalance = make(chan int)
var chDisposit = make(chan int)

func bankSystem() {
	balance := 0
	for {
		if len(chDisposit) == 0 {

			select {
			case x := <-chDisposit:
				balance += x
			case chBalance <- balance:
				// do nothing
			}
		} else {
			select {
			case x := <-chDisposit:
				balance += x
			}
		}
	}
}

func deposit(x int, w *sync.WaitGroup) {
	// balance += x
	chDisposit <- x
	w.Done()

}

func finalBalance() int {
	// return balance
	return <-chBalance
}

func main() {
	go bankSystem()
	fmt.Println("Balance before disposit : ", finalBalance())
	w := &sync.WaitGroup{}
	for i := 0; i < 300000; i++ {
		w.Add(1)
		go deposit(100, w) // เกิดการเขียนข้อมูลที่ยังไม่อัพเดท balance
	}
	fmt.Println("Balance before waitgroup : ", finalBalance())
	w.Wait()
	fmt.Println("Balance finalBalance : ", finalBalance())
	fmt.Println("Balance finalBalance : ", finalBalance())
}
