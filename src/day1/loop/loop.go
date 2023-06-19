package main

import "fmt"

func main() {
	fmt.Println("Loop")

	for a := 0; a < 100; a++ {
		fmt.Println(a)
	}
	fmt.Println(a)
	i := 10

	for i < 100 {
		fmt.Println(i)
		i += 2
	}

	for {
		fmt.Println("Hello")
	}

}
