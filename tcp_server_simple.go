package main

import (
	"fmt"
	"net"
)

func doServerWork(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading... ", err.Error())
			return //terminate program
		}
		fmt.Printf("recieved data: %v", string(buf))
	}
}

func main() {
	fmt.Println("Starting server....")

	listener, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		fmt.Println("Error listening... ", err.Error())
		return
	}

	for { // listen and accept connections from clients
		conn, err := listener.Accept()
		if err != nil {
			return // terminate program
		}
		go doServerWork(conn)
	}

}
