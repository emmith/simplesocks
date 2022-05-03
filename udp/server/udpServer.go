package main

import (
	"log"
	"net"
	"os"
)

var limitChan = make(chan bool, 1000)

func recvMessage(client *net.UDPConn) {
	var message []byte
	message = make([]byte, 1024)

	len, _, err := client.ReadFromUDP(message)
	if err != nil {
		log.Println("read udp message error", err.Error())
	}
	res := string(message[:len])
	log.Println(res)
	<-limitChan
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:9800")
	udpConn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatal("start udp server failed!\n")
		os.Exit(1)
	}

	defer udpConn.Close()

	log.Println("UDP server is running")

	for {
		limitChan <- true
		go recvMessage(udpConn)
	}
}
