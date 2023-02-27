package server

import (
	"log"
	"net"
	"sync"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Message struct {
	msg      string
	time     string
	UserName string
}

var (
	openConnections = make(map[string]Clients)
	entering        = make(chan Message)
	leaving         = make(chan Message)
	messages        = make(chan Message)
	history         []string
	mutex           sync.Mutex
)

type Clients struct {
	Name string
	Conn net.Conn
}

func Server(s string) {
	ln, err := net.Listen("tcp", ":"+s)
	logFatal(err)
	defer ln.Close()
	go broadcastMessage()
	for {
		conn, err := ln.Accept()
		logFatal(err)
		go handleConn(conn)
	}
}
