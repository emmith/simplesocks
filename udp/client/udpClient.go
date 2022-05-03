package main

import (
	"log"
	"net"
	"os"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:9800")
	client, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("UDP Client is dailing failed!")
		os.Exit(1)
	}

	defer client.Close()

	client.Write([]byte("i am udp client"))
}
