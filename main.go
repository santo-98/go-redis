package main

import (
	"github.com/santo-98/go-redis/server"
)

func main() {
	server.Start("localhost", "3000")
	// client.Connect("localhost", "3000")
}
