package main

import (
	"fmt"
	"sync"
)

var (
	mu      = sync.Mutex{}
	balance = 0
)

func deposit(x int, w *sync.WaitGroup) {
	// critical section
	// Lock
	mu.Lock()
	defer func() {
		mu.Unlock()
		w.Done()
	}()

	balance += x
	addBonus(1)

}

func AddBonus(x int) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()
	addBonus(x)
}

func addBonus(x int) {
	balance += x

}

func finalBalance() int {
	return balance
}

func main() {

	w := &sync.WaitGroup{}
	for i := 0; i < 200000; i++ {
		w.Add(1)
		go deposit(100, w) // เกิดการเขียนข้อมูลที่ยังไม่อัพเดท balance
	}

	w.Wait()
	fmt.Println("Balance finalBalance : ", finalBalance())
}
