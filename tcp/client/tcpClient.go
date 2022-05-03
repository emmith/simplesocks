package main

import (
	"log"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("TCP Client is dailing failed!")
		os.Exit(1)
	}

	client.Write([]byte("i am tcp client"))

	client.Close()
}
