package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp","localhost:8080")

	if err != nil {
		fmt.Println("Something went wrong during Dialing")
	}

	defer conn.Close()

}