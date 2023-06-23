package main

import (
	"fmt"
	"net"
	"os"
)

const (
	server = ":3456"
)

func main() {
	fmt.Println("tcp client..")
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Received error while connecting...")
		os.Exit(0)
	}
	defer conn.Close()
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Receiver error while reading...")
		os.Exit(0)
	}
	fmt.Println("Read ", n, "bytes...")
	fmt.Println("data is ", string(data))

	n, err = conn.Write([]byte("Thanks for accepting connection..."))
	if err != nil {
		fmt.Println("Received error while writing...")
		os.Exit(0)
	}
	fmt.Println("Wrote", n, "bytes")
}
