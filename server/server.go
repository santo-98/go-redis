package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Start(host string, port string) {
	fmt.Print("Starting ", host, " : ", port)
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
	buf, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	fmt.Println("Client message:", string(buf[:len(buf)-1]))

	conn.Write([]byte("message: " + string(buf[:len(buf)-1])))

	conn.Close()
}
