// Implementing fanout and fanin

package main

import (
	"fmt"
	"sync"
)

const (
	Fanout = 3
)

var (
	wg = sync.WaitGroup{}
)

func Generator(n int) chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func Square(ch chan int) chan int {
	wg.Add(1)
	channel := make(chan int)
	swg := sync.WaitGroup{}
	fn := func() {
		defer swg.Done()
		for d := range ch {
			channel <- d * d
		}
	}
	swg.Add(Fanout)
	for i := 0; i < Fanout; i++ {

		go fn()
	}

	go func() {
		defer wg.Done()
		swg.Wait()
		close(channel)
	}()
	return channel

}

func Display(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range ch {
			fmt.Println(d)
		}
	}()
}

func main() {
	// ch := Generator(100)
	// sch := Square(ch)
	// Display(sch)

	//Display(Square(Generator(100)))
	Display(Square(Square(Generator(100))))
	wg.Wait()
}
