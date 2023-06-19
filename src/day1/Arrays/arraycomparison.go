package main

import "fmt"

func main() {
	fmt.Println("Array comparison")

	var array1 = [5]int{1, 2, 3, 4, 5}

	var array2 = [5]int{1, 2, 3, 40, 5}

	if array1 == array2 {
		fmt.Println("equal...")
	} else {
		fmt.Println("Not equal...")
	}
}
