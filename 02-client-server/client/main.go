package main

import (
	"io"
	"net"
	// "net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// TODO: connect to server on localhost port 
	conn, err := net.Dial("tcp", "localhost:4922")
	if err != nil {
		logrus.Error(err)
	}
	//excuse last
	defer conn.Close()
	//send request
	// httppostturl := "https://192.168.88.105:4922"
	// request, err := http.NewRequest("GET", httppostturl)
	mustCopy(os.Stdout, conn)
}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logrus.Error(err)
	}
}
