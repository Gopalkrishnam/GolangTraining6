package main

import (
	"fmt"
	"time"
)

func Test() {
	fmt.Println("We are in Test method...")
}

func main() {
	fmt.Println("Goroutine...")

	go Test()

	go Test()

	go Test()

	go fmt.Println("Hello")

	time.Sleep(10 * time.Millisecond)

}
