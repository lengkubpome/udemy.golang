//  Chat Server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan string

var chMessage = make(chan string)
var chEnteringClient = make(chan client)
var chLeavingClient = make(chan client)

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err) // หยุดระบบ
	}

	go boardcaseter()

	for {
		conn, err := listenner.Accept() // รอ client เชื่อมต่อ
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	clientName := conn.RemoteAddr()
	chOutGoingMsg := make(chan string)
	chEnteringClient <- chOutGoingMsg
	fmt.Println("Client connected : ", clientName)

	go clientWriter(conn, chOutGoingMsg)
	scanner := bufio.NewScanner(conn) // return false ค่าเมื่อก่อน error

	for scanner.Scan() {
		msg := scanner.Text()
		chMessage <- fmt.Sprintf("%s => %s\n", clientName, msg)

	}

	chLeavingClient <- chOutGoingMsg
	fmt.Printf("%s disconnected.\n", clientName)
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func boardcaseter() {
	clients := make(map[client]bool) // เช็คจำนวน client ที่กำลังออนไลน์อยู่
	for {

		select {
		case msg := <-chMessage:
			for k := range clients {
				k <- msg
			}

		case client := <-chEnteringClient:
			clients[client] = true

		case client := <-chLeavingClient:
			clients[client] = false
		}
	}
}
