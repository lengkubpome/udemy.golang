package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		log.Println("New client is connected : ", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			return
		}
		countingDownHandler(conn)

	}

}

func countingDownHandler(conn net.Conn) {
	defer func() {
		io.WriteString(conn, "Your connection will be closed by server")
		conn.Close()
	}()

	io.WriteString(conn, "Enter number : ")
	count := 5
	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n")
		time.Sleep(time.Second)
		count--
		if count < 0 {
			io.WriteString(conn, "Enter number : ")
			break
		}
	}
}
