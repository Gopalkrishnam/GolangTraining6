package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Ticker....")
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer is expired")
		}
	}
}
