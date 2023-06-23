package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Timer has expired...")
	})

	timer1 := time.NewTimer(2 * time.Second)

	timer2 := time.NewTimer(3 * time.Second)

	timer3 := time.NewTimer(4 * time.Second)
	t1, t2, t3 := false, false, false
	for {
		select {
		case _, ok := <-timer1.C:
			if ok {
				fmt.Println("Timer1 expired")
			}
			t1 = ok
		case _, ok := <-timer2.C:
			if ok {
				fmt.Println("Timer2 expired")
			}
			t2 = ok
		case _, ok := <-timer3.C:
			if ok {
				fmt.Println("Timer3 expired")
			}
			t3 = ok
		}
		fmt.Println(t1, t2, t3)
		if t1 && t2 && t3 {
			break
		}
	}
}
