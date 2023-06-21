package main

import "fmt"

func Iterate(strslice []string, fn func(string)) {
	for _, val := range strslice {
		fn(val)
	}
}

func main() {
	fmt.Println("Higher order function")

	slice := []string{"Hello", "Mavenir", "How", "Are", "you"}

	fn := func(str string) {
		fmt.Println("#", str)
	}

	Iterate(slice, fn)

	Iterate(slice, func(str string) {
		fmt.Println("!!!", str)
	})
}
