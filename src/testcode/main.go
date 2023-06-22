package main

import "fmt"

func WorkWithSlice(ints *[]int) {
	fmt.Println(*ints)
	fmt.Println("WorkWithSlice::len:", len(*ints))
	fmt.Println("WorkWithSlice::cap:", cap(*ints))

	for _, v := range *ints {
		fmt.Println(v)
	}
	*ints = append(*ints, 5)
	*ints = append(*ints, 6)
	fmt.Println("WorkWithSlice::len:", len(*ints))
	fmt.Println("WorkWithSlice::cap:", cap(*ints))
}

func main() {
	fmt.Println("Demo: Slice to function")

	var slice = make([]int, 5, 5)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	WorkWithSlice(&slice)

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
}
