package main

import "fmt"

func main() {
	fmt.Println("Demo: Conditional statement")

	var a, b = 100, 1000

	// if a != b {
	// 	fmt.Println("a and b are not equal")
	// } else {
	// 	fmt.Println(" a and b are equal")
	// }

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a and b are equal")
	}

	if (a > b) && (a < b) {

	}

}
