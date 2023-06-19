// Defining constants

package main

import "fmt"

// const <name> <data type> = initilization

func main() {
	const c1 int = 100
	fmt.Println("Demo: Constants")
	fmt.Println(c1)

	const c2 = 100

	fmt.Println(c2)

	const c3, c4, c5 = 10, 3.14159, "Hello"
	fmt.Println(c3, c4, c5)
}
