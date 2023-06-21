package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time zone")

	l, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println(err)
	} else {
		tm := time.Date(2023, 6, 21, 14, 28, 0, 0, l)
		fmt.Println(tm)
	}
}
