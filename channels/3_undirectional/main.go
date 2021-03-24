package main

import "fmt"

// send only
func sender(ch chan<- int) {
	ch <- 88
	close(ch) // close ได้เฉพาะ send เท่านั้น
}

// receive only
func receiver(ch <-chan int, done chan bool) {
	fmt.Println("Receive ", <-ch)
	done <- true

}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sender(ch)
	go receiver(ch, done)
	<-done
	fmt.Println("Done")
}
