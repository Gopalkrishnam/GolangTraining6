// produce -> consumer

package main

import (
	"fmt"
	"sync"
)

var (
	channel = make(chan int, 10)
	wg      = sync.WaitGroup{}
)

func Producer(n int) {
	defer func() {
		close(channel)
		wg.Done()
	}()
	for i := 0; i <= n; i++ {

		channel <- i
		fmt.Println("Writing ", i)
	}
}

func Consumer() {
	// read this from channle
	defer wg.Done()
	for data := range channel {
		fmt.Println("Read", data)
		fmt.Println(data)
	}
}

func main() {
	fmt.Println("Producer/Concumer")

	wg.Add(2)
	go Producer(100)
	go Consumer()
	wg.Wait()

}
