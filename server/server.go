package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

type Data struct {
	Key   interface{}
	Value interface{}
}

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
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	data := decoder(buffer)
	fmt.Println(data.Key, data.Value)
	conn.Write([]byte("Thanks! Got your message:"))

	conn.Close()
}

func decoder(encodedData []byte) Data {
	var buf bytes.Buffer
	buf.Write(encodedData)
	decoder := gob.NewDecoder(&buf)

	var decodedData Data
	err := decoder.Decode(&decodedData)
	if err != nil {
		log.Fatal("Decode error:", err)
	}

	return decodedData
}
