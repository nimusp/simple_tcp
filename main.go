package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Launching server ...")
	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		simpleHandle(error)
	}

	connection, error := listener.Accept()
	if error != nil {
		simpleHandle(error)
	}

	for {
		message, error := bufio.NewReader(connection).ReadString('\n')
		if error != nil {
			simpleHandle(error)
		}

		fmt.Println("Message received:", message)
		answer := strings.ToUpper(message)
		connection.Write([]byte(answer + "/n"))
	}
}

func simpleHandle(err error) {
	fmt.Println(err)
	os.Exit(-1)
}
