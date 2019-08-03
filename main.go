package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Launching server ...")
	listener, error := net.Listen("tcp", ":8080")
	defer listener.Close()
	simpleHandle(error)

	fileToWrite, error := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fileToWrite.Close()
	simpleHandle(error)

	connection, error := listener.Accept()
	defer connection.Close()
	simpleHandle(error)

	for {
		message, error := bufio.NewReader(connection).ReadString('\n')
		simpleHandle(error)

		fmt.Println("Message received:", message)
		fileToWrite.WriteString(time.Now().String() + " received: " + message + "\n")
		fileToWrite.Sync()
		answer := strings.ToUpper(message)
		connection.Write([]byte(answer))
	}
}

func simpleHandle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
