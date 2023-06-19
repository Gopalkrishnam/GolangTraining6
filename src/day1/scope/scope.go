package main

import "fmt"

// global
// local
// block

var global int = 100

var (
	a, b, c = 10, 100.01, "test"
)

func main() {
	var localmain = 1000

	var (
		a string
		b int
		c float32
	)

	fmt.Println("Scope of varaible")
	fmt.Println(localmain)
	fmt.Println(global)
	{
		var block_var = 1000
		fmt.Println(block_var)
	}
	fmt.Println(block_var)
}
