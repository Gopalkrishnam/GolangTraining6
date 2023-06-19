package main

import "fmt"

func main() {
	fmt.Println("Demo: Comparison operators")

	var a, b = 100, 200

	var c = a == b //equality

	fmt.Println(c)

	c = a != b
	fmt.Println(c)

	c = a < b
	fmt.Println(c)

	c = a > b
	fmt.Println(c)

	c = a <= b
	fmt.Println(c)

	c = a >= b
	fmt.Println(c)

	var x, y = 10, 1.1

	c = x != y
}
