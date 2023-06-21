package main

import "fmt"

func ArrayAsArg(arr [5]int) {
	fmt.Println(arr)

	arr[0] = 100
	arr[4] = 200
	fmt.Println(arr)
}

func main() {
	fmt.Println("Function array as argument")

	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	ArrayAsArg(array)
	fmt.Println(array)
	fmt.Println("Calling by reference")
	fmt.Println(array)
	ArrayByRef(&array)
	fmt.Println(array)
}

func ArrayByRef(arr *[5]int) {
	fmt.Println("ArrayByRef", arr)

	arr[0] = 100
	arr[4] = 200
	fmt.Println(arr)
}
