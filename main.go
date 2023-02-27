package main

import (
	"beka/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Listening on the port :8989")
		server.Server("8989")
	}
	if len(os.Args[1:]) == 1 {
		server.Server(os.Args[1])
		fmt.Println("Listening on the port :", os.Args[1:])
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}
