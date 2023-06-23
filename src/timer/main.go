package main

import (
	"fmt"
	"mavenir/timer"
	"time"
)

func TimerCB(timerId string, timerInterval int) {
	fmt.Println("Timer Expired ", timerId, timerInterval)
}

func main() {
	fmt.Println("Demo timers...")
	timer.Init()
	timer.StartTimer(10, TimerCB)
	timer.StartTimer(5, TimerCB)
	timer.StartTimer(3, TimerCB)
	timer.StartTimer(8, TimerCB)
	timer.StartTimer(6, TimerCB)
	timer.StartTimer(21, TimerCB)
	timer.StartTimer(5, TimerCB)
	time.Sleep(30 * time.Second)
}
