package main

import (
	"fmt"
	"time"
)

// 2023/06/21 14:08:00
func main() {
	fmt.Println("time package function/methods")

	currTime := time.Now()

	fmt.Println("Current year is ", currTime.Year())
	fmt.Println("Current Month is ", currTime.Month())
	fmt.Println("Day", currTime.Day())
	fmt.Println("Hour, minute, second", currTime.Hour(), currTime.Minute(), currTime.Second())

	s := fmt.Sprintf("%d/%d/%d %d:%d:%d", currTime.Year(), currTime.Month(), currTime.Day(), currTime.Hour(), currTime.Minute(), currTime.Second())

	fmt.Println(s)

	timeFormat := "2006/01/02 15:04:05"

	timeStr := currTime.Format(timeFormat)

	fmt.Println(timeStr)

	t, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		fmt.Println("Received error ", err)
	} else {
		fmt.Println("Parsed time is", t)
	}

	

}
