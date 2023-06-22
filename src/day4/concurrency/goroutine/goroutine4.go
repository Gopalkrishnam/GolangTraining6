package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go routines in loop")
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			//fmt.Println("Go routine number", i)
			for j := n; j <= n+10; j++ {
				fmt.Println("from:", n, j)
			}
		}(i)
	}
	wg.Wait()
}
