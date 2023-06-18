package client

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
	Key    interface{}
	Value  interface{}
	Action string
}

type Config struct {
	Host string
	Port string
}

func (config Config) Set(key interface{}, value interface{}) {
	fmt.Println("Starting Client with ", config.Host, " : ", config.Port)
	connection, err := net.Dial("tcp", config.Host+":"+config.Port)

	if err != nil {
		fmt.Println("Client error: ", err)
		os.Exit(1)
	}

	enc := encoder(Data{
		Key:    key,
		Value:  value,
		Action: "set",
	})
	_, err = connection.Write(enc)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	bufio.NewReader(connection).ReadString('\n')

	defer connection.Close()
}

func (config Config) Get(key interface{}) (interface{}, interface{}) {
	fmt.Println("Starting Client with ", config.Host, " : ", config.Port)

	connection, err := net.Dial("tcp", config.Host+":"+config.Port)

	if err != nil {
		fmt.Println("Client error: ", err)
		os.Exit(1)
	}

	enc := encoder(Data{
		Key:    key,
		Action: "get",
	})
	_, err = connection.Write(enc)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	_, err = connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	data := decoder(buffer)

	defer connection.Close()
	return data.Key, data.Value
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
