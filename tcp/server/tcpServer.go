package main

import (
	"log"
	"net"
	"os"
)

func recvMessage(client net.Conn) error {
	var message []byte
	message = make([]byte, 1024)

	// 因为发送方可能发送大于1024字节的数据，所以要采用循环读取
	for {
		msgLen, err := client.Read(message)
		if err != nil {
			break
		}
		if msgLen > 0 {
			log.Println(string(message[0:msgLen]))
			// 给客户端回复相同的信息
			_, err := client.Write(message[0:msgLen])
			if err != nil {
				break
			}
		}
	}

	return nil
}

func main() {
	server, err := net.Listen("tcp", "localhost:9700")
	if err != nil {
		log.Fatal("start TCP server failed!\n")
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
