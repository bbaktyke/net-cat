package server

import (
	"fmt"
	"net"
)

func ValidName(username string, connection net.Conn) bool {
	if username == "" || len(username) == 0 {
		fmt.Fprintf(connection, "The username is the necessary condition to enter the chat\n")
		connection.Close()
		return false
	}

	for _, simbol := range username {
		if simbol < 32 || simbol > 127 {
			fmt.Fprintln(connection, "Incorrect input\n")
			connection.Close()
			// fmt.Fprintf(connection, "[%s][%s]:", time, username)
			return false
		}
	}

	for names, _ := range openConnections {
		if username == names {
			fmt.Fprintln(connection, "Username is already taken,  Please try to connect to server and use another name\n")
			connection.Close()
			return false
		}
	}
	return true
}
