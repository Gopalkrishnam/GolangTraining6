package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go routine with closure...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(a int) {
		defer wg.Done()
		for i := 0; i < a; i++ {
			fmt.Println("from Go routine 1", i)
		}
	}(100)
	wg.Add(1)
	go func(a int) {
		defer wg.Done()
		for i := 0; i < a; i++ {
			fmt.Println("from Go routine 2", i)
		}
	}(100)
	wg.Wait()
}
