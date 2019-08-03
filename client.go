package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer connection.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Text to send:")
		message, _ := reader.ReadString('\n')

		fmt.Fprint(connection, message+"\n")
		answer, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Println("Message from server " + answer)
	}
}
