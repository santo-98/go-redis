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

	s.Set("someKey", 1)
	key, value := s.Get("someKey")
	fmt.Println("Key and Value: ", key, value)
}
