package server

import "fmt"

func broadcastMessage() {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for _, user := range openConnections {
				if user.Name != msg.UserName {
					hst := fmt.Sprintf("\n[%s][%s]: %s", msg.time, msg.UserName, msg.msg)
					history = append(history, hst)
					fmt.Fprintf(user.Conn, "\n[%s][%s]:%s\n", msg.time, msg.UserName, msg.msg)
				}
				fmt.Fprintf(user.Conn, "[%s][%s]:", msg.time, user.Name)
			}
			mutex.Unlock()
		case cli := <-entering:
			mutex.Lock()
			for _, user := range openConnections {
				if user.Name != cli.UserName {
					fmt.Fprintf(user.Conn, "\n%s %s", cli.UserName, cli.msg)
				} else {
					for _, previous := range history {
						fmt.Fprintf(user.Conn, previous)
					}
				}
			}
			mutex.Unlock()

		case cli := <-leaving:
			mutex.Lock()
			for _, user := range openConnections {
				if user.Name != cli.UserName {
					fmt.Fprintf(user.Conn, "\n%s %s:", cli.UserName, cli.msg)
				}
			}
			mutex.Unlock()
		}
	}
}
