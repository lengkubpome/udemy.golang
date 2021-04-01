package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      = sync.RWMutex{}
	balance = 100
)

// In Case milisecond
// 50619100 --> no mutex
// 506817600 --> use mutex
// 50211300 --> use RWmutex

func readX(wg *sync.WaitGroup) int {
	mu.RLock()
	defer func() {
		mu.RUnlock()
		wg.Done()
	}()
	time.Sleep(50 * time.Millisecond)
	return balance

}
func main() {

	start := time.Now()
	defer func() {
		fmt.Println("Took : ", time.Since(start).Nanoseconds())

	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readX(&wg)
	}
	wg.Wait()
}
