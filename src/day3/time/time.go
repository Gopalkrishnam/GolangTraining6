package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time...")
	fmt.Println("Reading current time")
	t := time.Now()
	fmt.Println("Current time is", t)

	//Attributes year, month day, hour, minute, second

	t1 := time.Date(2023, 6, 21, 13, 55, 0, 0, &time.Location{})
	fmt.Println("Constructed time is ", t1)

	unixTime := t1.Unix()
	fmt.Println(unixTime)

	t2 := time.Unix(unixTime, 0)
	fmt.Println("Convereted unix time", t2)
}
