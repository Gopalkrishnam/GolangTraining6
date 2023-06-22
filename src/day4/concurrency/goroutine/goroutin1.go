package main

// Waitgroups
import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup = sync.WaitGroup{}
)

func Generator(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	fmt.Println("Demo: Goroutine...")
	wg.Add(2)
	go Generator(1000000)

	go Generator(1000000)
	wg.Wait()
	//time.Sleep(50 * time.Millisecond)
}
