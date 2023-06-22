package main

import (
	"fmt"
	"sync"
)

var (
	global = 0
	mutex  = sync.Mutex{}
)

func main() {
	fmt.Println("Contigency")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 25000; i++ {
			mutex.Lock()
			global++
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 25000; i++ {
			mutex.Lock()
			global++
			mutex.Unlock()
		}
	}()
	wg.Wait()

	fmt.Println("Value of global", global)
}
