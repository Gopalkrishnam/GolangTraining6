package main

import "fmt"

func main() {
	fmt.Println("Demo arrays")
	array := [5]int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Println(i, v)
	}

}
