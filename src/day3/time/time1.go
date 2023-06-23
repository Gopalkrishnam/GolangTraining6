package main

import (
	"fmt"
	"time"
)

const (
	Expiry = 45
)

func main() {
	fmt.Println("Time operations")

	currentTime := time.Now()
	fmt.Println(currentTime)
	unixTime := currentTime.Unix()

	expiryTime := unixTime + 45*60

	newTime := time.Unix(expiryTime, 0)
	fmt.Println(newTime)

	goTime := currentTime.Add(45 * time.Minute)
	fmt.Println(goTime)

	

}
