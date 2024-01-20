package main

import (
	"io"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "localhost:4922")
	if err != nil {
		logrus.Error(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleConn(conn)
	}

}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
