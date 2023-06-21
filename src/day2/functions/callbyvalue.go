package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var a, b = 100, 200
	fmt.Println(a, b)
	// a, b = swap(a, b)
	// fmt.Println(a, b)

	// a, b = SwapByValue(a, b)

	SwapByRef(&a, &b)
	fmt.Println(a, b)

}

func SwapByValue(a, b int) (int, int) {
	c := a
	a = b
	b = c

	return a, b
}

func SwapByRef(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
