package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var Peguin = []string{
	"Welcome to TCP-Chat!",
	"	 _nnnn_ ",
	"        dGGGGMMb",
	"       @p~qp~~qMb",
	"       M|@||@) M|",
	"       @,----.JM|",
	"      JS^\\__/  qKL",
	"     dZP        qKRb",
	"    dZP          qKKb",
	"   fZP            SMMb",
	"   HZM            MMMM",
	"   FqM            MMMM",
	" __| \".        |\\dS\"qML",
	" |    `.       | `' \\Zq",
	"_)      \\.___.,|     .'",
	"\\____   )MMMMMP|   .'",
	"     `-'       `--'",
}

func handleConn(conn net.Conn) {
	// To Print Logo
	for _, v := range Peguin {
		fmt.Fprintf(conn, v)
		fmt.Fprintf(conn, "\n")
	}
	conn.Write([]byte("[ENTER YOUR NAME]:"))
	clientName, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("Some problem with your name..."))
		return
	}
	clientName = strings.Trim(clientName, "\r\n")
	if !ValidName(clientName, conn) {
		return
	}

	User := Clients{clientName, conn}
	mutex.Lock()
	openConnections[clientName] = User
	mutex.Unlock()
	mutex.Lock()
	if len(openConnections) > 10 {
		fmt.Fprintf(User.Conn, "Chat is full\n")
		delete(openConnections, clientName)
		conn.Close()
		mutex.Unlock()
		return
	}
	mutex.Unlock()

	joined := Message{
		msg:      "has joined our chat...\n",
		time:     time.Now().Format("2006-01-02 15:04:05"),
		UserName: clientName,
	}

	entering <- joined

	fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), clientName)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := strings.Trim(input.Text(), " ")
		if !isValidtext(text) {
			fmt.Fprintln(conn, "The empty messages are prohibited")
			fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), clientName)
			continue
		}
		messageNew := Message{
			msg:      text,
			time:     time.Now().Format("2006-01-02 15:04:05"),
			UserName: clientName,
		}

		messages <- messageNew

	}
	text := Message{
		msg:      "has left our chat...",
		time:     time.Now().Format("2006-01-02 15:04:05"),
		UserName: clientName,
	}
	leaving <- text
	mutex.Lock()
	delete(openConnections, clientName)
	conn.Close()
	mutex.Unlock()
}
