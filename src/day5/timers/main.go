package main

import (
	"fmt"
	"mavenir/timer"
	"time"
)

func TimerCallBack(timerID string, interval int) {
	fmt.Println("timerId:", timerID, "interval:", interval)
}

func main() {
	fmt.Println("Timer implementation....")

	timer.Init()
	timer.StartTimer(10, TimerCallBack)
	timer.StartTimer(5, TimerCallBack)
	timer.StartTimer(20, TimerCallBack)
	timer.StartTimer(7, TimerCallBack)

	timerId, _ := timer.StartTimer(2, TimerCallBack)

	timer.CancelTimer(timerId)

	time.Sleep(1 * time.Minute)
}
