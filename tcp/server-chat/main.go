package main

import (
	"fmt"
	"github.com/jonfk/golang-chat/tcp/common"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	datasize_str, err := common.ReadMsg(conn)
	if err != nil {
		if err == io.EOF {
			conn.Close()
			return
		}
		log.Println(err)
		return
	}

	datasize, err := strconv.Atoi(datasize_str)
	if err != nil {
		log.Fatal(err)
	}

	msg := strings.Repeat("x", datasize)
	err = common.WriteMsg(conn, msg)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
}
