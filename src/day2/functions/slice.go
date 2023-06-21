package main

import "fmt"

func SliceAsArgument(slice []int) {
	fmt.Println(slice)

	slice[0] = 100

	slice[len(slice)-1] = 100000

	fmt.Println(slice)
	slice = append(slice, 100)
	fmt.Println(slice)
}

func main() {
	fmt.Println("Demo slice as argument")

	slice := make([]int, 5, 5)

	fmt.Println(slice)

	SliceAsArgument(slice)
	fmt.Println("After calling function", slice)
}
