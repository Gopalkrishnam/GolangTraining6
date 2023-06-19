package main

import "fmt"

func main() {
	fmt.Println("Demo: Variable")

	var a, b, c = 10, 10.1, "hello"

	var teststr string

	fmt.Println("Value of teststr ", teststr)

	fmt.Println(a, b, c)

	fmt.Printf("%T\n%T\n%T\n", a, b, c)
}
