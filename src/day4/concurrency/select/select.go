// Goroutine 1 Ch1
// Goroutine 2 Ch2  Goroutine 4
// Goroutine 3 Ch3

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("Select")
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	wg.Add(4)
	go func() {
		defer wg.Done()
		defer close(ch1)
		ch1 <- 10
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)

		ch2 <- 20
	}()

	go func() {
		defer wg.Done()
		defer close(ch3)

		ch3 <- 30
	}()

	go func() {
		defer wg.Done()
		c1, c2, c3 := true, true, true
		for {
			select {
			case v, ok := <-ch1:
				c1 = ok
				if ok {
					fmt.Println(v)
				}
			case v, ok := <-ch2:
				c2 = ok
				if ok {
					fmt.Println(v)
				}
			case v, ok := <-ch3:
				c3 = ok
				if ok {
					fmt.Println(v)
				}
			}
			if c1 == false && c2 == false && c3 == false {
				break
			}
		}
	}()
	wg.Wait()
}
