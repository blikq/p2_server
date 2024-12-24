package server

import (
	"fmt"
	"log" // import the log package
	"net"
)

func Start() {
	// Start the server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		// panic(err)
		fmt.Println("error while listening: ", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error

			fmt.Println("error tcp: ", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// handle connection
	fmt.Println("connected to: ", conn.RemoteAddr())
	for {
		var buffer [1024]byte // 1KB buffer to read data
		_, err := conn.Read(buffer[:])
		if err != nil {
			log.Printf("err while reading from conn: %v, exiting ...", err)
			return
		}
	}
}
