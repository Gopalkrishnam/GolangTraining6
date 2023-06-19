package main

import "fmt"

func main() {
	fmt.Println("Slice continued...")

	var slice1 = make([]int, 10, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = 100 * (i + 1)
	}

	fmt.Println(len(slice1))

	fmt.Println(cap(slice1))

	fmt.Println(slice1)

	//slice1 = append(slice1, 100)

	fmt.Println(len(slice1))

	fmt.Println(cap(slice1))

	fmt.Println(slice1)

	//slice1[0] = 1
	fmt.Println(slice1)

	fmt.Println(slice1[0:5])
	fmt.Println(slice1[6:])

	fmt.Println(slice1)
	//slice1 = append(slice1[:5], slice1[6:]...)

	slice1 = append(slice1[:5], slice1[4:]...)

	slice1[5] = 10000

	fmt.Println(slice1)
}
