// package client
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

// func Connect(host string, port string) {
func main() {
	// fmt.Print("Starting Client with ", host, " : ", port)
	// connection, err := net.Dial("tcp", host+":"+port)
	fmt.Println("Starting Client with ", "localhost:3000")
	connection, err := net.Dial("tcp", "localhost:3000")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println(runtime.NumGoroutine())
	_, err = connection.Write([]byte("Hello from client"))

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	message, _ := bufio.NewReader(connection).ReadString('\n')

	// Print server relay.
	log.Print("Server message: " + message)
	defer connection.Close()
}
