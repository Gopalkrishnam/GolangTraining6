package main

import "fmt"

func main() {
	fmt.Println("Slice continued")

	var slice1 = []int{}

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 2)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 3)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 4)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1 = append(slice1, 5)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

}
