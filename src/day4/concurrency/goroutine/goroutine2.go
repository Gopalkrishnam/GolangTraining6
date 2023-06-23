package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go routine with closure")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("We are in annonymous function/ this called as closure")
	}()
	wg.Wait()
}
