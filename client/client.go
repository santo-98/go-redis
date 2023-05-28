// package client
package main

import (
	"bufio"
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

// func Connect(host string, port string) {
func main() {
	// fmt.Print("Starting Client with ", host, " : ", port)
	// connection, err := net.Dial("tcp", host+":"+port)
	fmt.Println("Starting Client with ", "localhost:3000")

	connection, err := net.Dial("tcp", "localhost:3000")

	if err != nil {
		fmt.Println("Client error: ", err)
		os.Exit(1)
	}

	enc := encoder("1", 1)
	_, err = connection.Write(enc)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	message, _ := bufio.NewReader(connection).ReadString('\n')

	// Print server relay.
	log.Print("Server message: " + message)
	defer connection.Close()
}

func encoder(key interface{}, value interface{}) []byte {
	packet := Data{
		Key:   key,
		Value: value,
	}

	gob.Register(Data{})
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(packet)
	if err != nil {
		log.Fatal("Encode error:", err)
	}

	return buf.Bytes()
}
