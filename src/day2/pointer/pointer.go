package main

import "fmt"

func main() {
	var integer int = 100

	fmt.Println(integer)

	var ptrInteger *int

	ptrInteger = &integer

	fmt.Println(ptrInteger)
	fmt.Println(*ptrInteger)
}
