package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Timer expired...")
	})

	time.Sleep(30 * time.Second)
}
