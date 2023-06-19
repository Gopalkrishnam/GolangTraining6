package main

import "fmt"

/*
var <variable_name> <data type> = assignment
*/

func main() {
	var first int = 0
	fmt.Println("Demo variable declaration")

	fmt.Println("Value of variable first is ", first)

	first = 100

	fmt.Println("Value of variable first after assignment", first)

	var second int = 1000

	fmt.Println("Value of variable second ", second)
}
