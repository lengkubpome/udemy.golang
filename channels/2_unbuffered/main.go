package main

import (
	"fmt"
)

func sender(ch chan int, done chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete.")
	done <- "Done from sender"
	fmt.Println("Sender done")
}

func receiver(ch chan int, done chan string) {

	for v := range ch {
		fmt.Println("\tReceive value : ", v)
	}

	done <- "Done from receiver"
	close(done)
}

func square(ch chan int, chSquare chan int, done chan string) {
	for v := range ch {
		chSquare <- v * v
	}
	close(chSquare)
	done <- "Done from square"
}

func main() {
	ch := make(chan int)
	chSquare := make(chan int)
	done := make(chan string)

	go sender(ch, done)
	go square(ch, chSquare, done)
	go receiver(chSquare, done)

	doneStatus := <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)

	for v := range done {
		fmt.Println("Done status : ", v)
	}

	fmt.Println("Main exit")

}
