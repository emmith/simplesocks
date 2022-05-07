package main

import (
	"log"
	"net"
)

func main() {
	client, err := net.Dial("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("TCP Client is dailing failed!")
	}

	_, err = client.Write([]byte("i am tcp client"))
	if err != nil {
		log.Fatal(err)
	}

	msg := make([]byte, 1024)
	_, err = client.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(msg))

	client.Close()
}
