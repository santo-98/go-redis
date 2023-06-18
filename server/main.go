package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type Data struct {
	Key    interface{}
	Value  interface{}
	Action string
}

var dataStore = make(map[interface{}]interface{})

func main() {
	host := flag.String("host", "localhost", "provide value for host")
	port := flag.String("port", "3000", "provide value for port")
	flag.Parse()

	fmt.Println("Starting ", *host, " : ", *port)
	listener, err := net.Listen("tcp", *host+":"+*port)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	defer listener.Close()

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
	switch action := data.Action; action {
	case "set":
		setData(data, conn)
	case "get":
		getData(data, conn)
	}
}

func setData(data Data, conn net.Conn) {
	dataStore[data.Key] = data.Value
	conn.Write([]byte("Data has Set!"))
	log.Println(data.Key, data.Value)
	conn.Close()
}

func getData(data Data, conn net.Conn) {
	enc := encoder(Data{
		Key:   data.Key,
		Value: data.Value,
	})

	conn.Write(enc)
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

func encoder(data Data) []byte {
	gob.Register(Data{})
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(data)
	if err != nil {
		log.Fatal("Encode error:", err)
	}

	return buf.Bytes()
}
