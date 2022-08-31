package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3001")

	if err != nil {
		fmt.Println("Error, dialing... ", err.Error())
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please let us know, how do we call you ?")
	clientName, _ := inputReader.ReadString('\n')

	trimmedClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What is your message to the server ? Type Q to quit")

		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		_, err := conn.Write([]byte(trimmedClient + " syas: " + trimmedInput + "\n"))

		if err != nil {
			fmt.Println("Unable to communicating to the server... ", err.Error())
		}
	}
}
