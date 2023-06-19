package main

import "fmt"

func main() {
	var a = 9900
	var b = "hello"

	switch {

	case a < 100:
		fmt.Println("Less than 100")

	case a == 100:
		fmt.Println("Both are equal")
	case a == 99:
		fmt.Println("a is equal to 99")
	case b == "hello":
		fmt.Println("b is hello")
	}
}
