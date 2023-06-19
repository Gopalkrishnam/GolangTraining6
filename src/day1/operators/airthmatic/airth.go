package main

import "fmt"

func main() {
	fmt.Println("Demo: Airthmatic operators...")
	/*
		+
		-
		*
		/
		%
	*/

	var a, b = 100, 200

	var result = a + b

	fmt.Println("Result is ", result)

	var res1 = a - b

	fmt.Println("Res1 is ", res1)

	var res2 = a * b

	fmt.Println(res2)

	res2 = a / b
	fmt.Println(res2)

	a = 10
	b = 3

	res2 = a % b
	fmt.Println(res2)
}
