package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp","localhost:8080")

	if err != nil {
		fmt.Println("Something went wrong during creating Server")
	}

	defer listener.Close()
	fmt.Println("Server is Running on 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Something went wrong during Connecting")
		}

		go handleIncomingConnection(conn)
	}
}

func handleIncomingConnection(conn net.Conn) {
	defer conn.Close()
	
	reader := bufio.NewReader(conn)

	line,err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	if line ==""  {
		fmt.Println("EMPTY")
	}

	fmt.Println("reader::::",line)

}