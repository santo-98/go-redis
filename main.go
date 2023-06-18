package main

import (
	"fmt"

	"github.com/santo-98/go-redis/client"
)

func main() {
	s := client.Config{
		Host: "localhost",
		Port: "3000",
	}

	s.Set("new", 1)
	key, value := s.Get("ne1w")
	fmt.Println("Key and Value: ", key, value)
}
