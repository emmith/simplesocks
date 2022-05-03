package main

import (
	"io"
	"log"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1024)

	// 因为发送方可能发送大于1024字节的数据，所以要采用循环读取
	for {
		len, err := client.Read(message)
		if err == io.EOF {
			break
		}
		if len > 0 {
			log.Println(string(message[0:len]))
		}
	}

	return nil
}

func main() {
	server, err := net.Listen("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("start UDP server failed!\n")
		os.Exit(1)
	}
	defer server.Close()

	log.Println("TCP server is running...")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error\n")
			continue
		}

		log.Println("the tcp client is connectted...")
		go recvMessage(client)
	}
}
