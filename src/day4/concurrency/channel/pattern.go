//
//
//task1  task2 task3
//

// reading ->  parsing request -> processing -> writing back

// Generate -> Squre -> Read Display result
//
//	Squre
//
// generator Square  Display
//
//	Square
package main

import (
	"fmt"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

func Generator(n int) chan int {

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

func Squre(ch chan int) chan int {
	// read data from generator channle
	//square number
	// write to channel to display
	channel := make(chan int)

	go func() {
		defer wg.Done()
		defer close(channel)
		for d := range ch {
			channel <- d * d
		}
	}()
	return channel
}

func Display(ch chan int) {
	// read from channl and display
	go func() {
		defer wg.Done()
		for d := range ch {
			fmt.Println(d)
		}
	}()
}

func main() {
	fmt.Println("Demo pipeline")
	wg.Add(3)
	ch := Generator(100)

	sch := Squre(ch)

	Display(sch)

	//Display(Squre(Generator(100)))
	//Display(Squre(Squre(Generator(100))))
	wg.Wait()
}
