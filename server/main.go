package main

import (
	"fmt"
	"net"
	"time"
)

const serverAddress = "localhost:8081"
const maxMessageSize = 1024
const timeout = 30

func main() {
	fmt.Println("starting server")
	server(serverAddress)
}

func server(address string) {

	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		fmt.Println("error trying to listen packet: " + err.Error())
		return
	}
	defer pc.Close()

	buffer := make([]byte, maxMessageSize)

	for {
		n, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			fmt.Println("error trying to read packet: " + err.Error())
			return
		}

		fmt.Printf("%s received from %s \n", string(buffer[:n]), addr.String())

		deadline := time.Now().Add(timeout * time.Second)
		err = pc.SetWriteDeadline(deadline)
		if err != nil {
			fmt.Println("error trying to set answer packet write timeout: " + err.Error())
			return
		}

		_, err = pc.WriteTo(buffer[:n], addr)
		if err != nil {
			fmt.Println("error trying to write answer packet: " + err.Error())
			return
		}

		fmt.Println("message processed")
	}
}
