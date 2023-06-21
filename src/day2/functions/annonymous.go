package main

import "fmt"

func main() {
	fmt.Println("Annonymous function")

	func(msg string) {
		fmt.Println("Hello ", msg)
	}("Mavenir")

	fun := func(a, b int) int {
		return a + b
	}

	ans := fun(10, 20)

	fmt.Println(ans)
}
