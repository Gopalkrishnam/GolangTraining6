// Queue
// write -> Queue -> read

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channles")
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("In goroutine")
		fmt.Println("Before read")
		v := <-ch
		fmt.Println("After read")
		fmt.Println("Data read from the channel...", v)
	}()
	fmt.Println("Before write")
	ch <- 10
	fmt.Println("After write")
	wg.Wait()
}
