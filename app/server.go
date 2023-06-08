package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error receiving message from client!", err.Error())
		}
		// fmt.Println(msg)
		fmt.Println("connection was successful! Writing message...")
		conn.Write([]byte("+PONG\r\n"))
	}

}
