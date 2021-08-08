package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

const serverAddress = "localhost:8081"
const serverMessage = "test"
const timeout = 30

func main() {
	fmt.Println("starting client")
	client(serverAddress, serverMessage)
}

func client(address string, message string) {

	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("error trying to resolve address: " + err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("error trying to dial address: " + err.Error())
		return
	}
	defer conn.Close()

	_, err = io.Copy(conn, strings.NewReader(message))
	if err != nil {
		fmt.Println("error trying to send message: " + err.Error())
		return
	}

	buffer := make([]byte, len(message))
	deadline := time.Now().Add(timeout * time.Second)

	err = conn.SetReadDeadline(deadline)
	if err != nil {
		fmt.Println("error trying to set answer packet read timeout: " + err.Error())
		return
	}

	_, _, err = conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println("error trying to read answer packet: " + err.Error())
		return
	}

	fmt.Printf("%s received from server", string(buffer))
}
