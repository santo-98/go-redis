package server

import (
	"fmt"
	"net"
	"os"
)

func Start(host string, port string) {
	fmt.Println("Starting ", host, " : ", port)
	listener, err := net.Listen("tcp", host+":"+port)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening", host, " : ", port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		go processConnection(connection)
	}
}

func processConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	conn.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))

	conn.Close()
}
