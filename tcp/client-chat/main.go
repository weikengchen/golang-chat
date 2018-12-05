package main

import (
	"bufio"
	"fmt"
	"github.com/jonfk/golang-chat/tcp/common"
	"io"
	"log"
	"net"
	"os"
	"time"
	"strconv"
)

const (
	CONN_HOST = "52.53.139.26"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	fmt.Print("Enter the size of the packet to receive: ")
	// Read the data size
	reader := bufio.NewReader(os.Stdin)
	datasize_str, err := reader.ReadString('\n')
	datasize_str = datasize_str[:len(datasize_str) - 1]
	if err != nil {
		log.Fatal(err)
	}

	datasize, err := strconv.Atoi(datasize_str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the number of connections to make: ")
	// Read the number of connections to make
	reader = bufio.NewReader(os.Stdin)
	testnum_str, err := reader.ReadString('\n')
	testnum_str = testnum_str[:len(testnum_str) - 1]
	if err != nil {
		log.Fatal(err)
	}

	testnum, err := strconv.Atoi(testnum_str)
	if err != nil {
		log.Fatal(err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	for i := 0; i < testnum; i++ {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Fatal(err)
		}

		writeInput(conn, datasize)

		for {
			msg, err := common.ReadMsg(conn)
			// Receiving EOF means the data has been sent
			if err == io.EOF {
				conn.Close()

				break
			}

			if err != nil {
				log.Fatal(err)
			}
			_ = msg
		}

		conn.Close()
	}
	end := time.Now()
	fmt.Printf("Time taken: %s\n", end.Sub(start))
}

func writeInput(conn *net.TCPConn, datasize int) {
	datasize_str := strconv.Itoa(datasize)

	err := common.WriteMsg(conn, datasize_str)
	if err != nil {
		log.Fatal(err)
	}
}

