package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func running(name string, track chan<- struct{}) {
	defer func() {
		fmt.Printf("### %s stopped running ###\n", name)
	}()

	rand.Seed(time.Now().Unix())
	// <- chan Time
	chFinish := time.After(time.Duration(10*time.Second + time.Duration(rand.Intn(100))*time.Millisecond))

	for isRacing() {
		select {
		case <-chFinish:
			track <- struct{}{}
		default:
			// do nothing
		}
	}
}

func isRacing() bool {
	select {
	case <-finishRacing:
		return false
	default:
		return true
	}

}

var finishRacing = make(chan struct{})

func main() {
	defer func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Done")
	}()

	// จัดการ cancel โปรแกรมด้วยวิธีกด Crl+c
	chInterrupt := make(chan os.Signal, 1)
	signal.Notify(chInterrupt, os.Interrupt, syscall.SIGBUS, syscall.SIGINT)

	track1 := make(chan struct{})
	track2 := make(chan struct{})
	track3 := make(chan struct{})
	abort := make(chan struct{})

	go running("h1", track1)
	go running("h2", track2)
	go running("h3", track3)

	// รับข้อมูล Enter จากคีย์บอร์ด
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("About to abort.")
		abort <- struct{}{}
	}()

	lapseTricker := time.NewTicker(1 * time.Second)
	lapseChan := lapseTricker.C

	for done := false; !done; {
		select {
		case <-track1:
			fmt.Println("Winner is Horse 1")
			done = true
		case <-track2:
			fmt.Println("Winner is Horse 2")
			done = true
		case <-track3:
			fmt.Println("Winner is Horse 3")
			done = true
		case v := <-lapseChan:
			fmt.Println("Horses are running : ", v.Second())
		case <-abort:
			close(finishRacing)
			done = true
			fmt.Println("User pressing Enter.")
		case v := <-chInterrupt:
			close(finishRacing)
			done = true
			fmt.Printf("User interrupt by : %v\n", v)
		}
	}

}
