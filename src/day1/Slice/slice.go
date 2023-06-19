package main

import "fmt"

func main() {
	fmt.Println("Demo Slice...")

	var sliceInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(len(sliceInt))

	fmt.Println(cap(sliceInt))

	fmt.Println(sliceInt)

	// for i, v := range sliceInt {
	// 	fmt.Println(i, v)
	// }

	// for _, v := range sliceInt {
	// 	fmt.Println(v)
	// }

	sliceInt = append(sliceInt, 11)

	fmt.Println(len(sliceInt))

	fmt.Println(cap(sliceInt))

	fmt.Println(sliceInt)

	sliceInt = append(sliceInt, 12)
	sliceInt = append(sliceInt, 13)
	sliceInt = append(sliceInt, 14)
	fmt.Println(sliceInt)
	fmt.Println(len(sliceInt))
	fmt.Println(cap(sliceInt))
}
