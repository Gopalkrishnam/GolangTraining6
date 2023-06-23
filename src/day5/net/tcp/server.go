package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Demo: Tcp server.....")
	lst, err := net.Listen("tcp", ":3456")
	if err != nil {
		fmt.Println("Received err while listening, hence can't proceed", err)
		os.Exit(0)
	}
	for {
		conn, err := lst.Accept()

		if err != nil {
			fmt.Println("Received error while accepting connection", err)
			continue
		}
		defer conn.Close()
		data := make([]byte, 1024)
		n, err := conn.Write([]byte("Sending ack for connection"))
		if err != nil {
			fmt.Println("Received error while reading...")
			continue
		}
		fmt.Println("Wrote ", n, "bytes..")
		n, err = conn.Read(data)
		if err != nil {
			fmt.Println("Received error while reading...")
			continue
		}
		fmt.Println("Receive ", n, "bytes ")
		fmt.Println("Data is ", string(data))
	}
}
