package math

import "fmt"

func Add(a, b int, ints ...int) int {
	s := a + b

	for _, v := range ints {
		s += v
	}
	return s
}

func Sub(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}

func Multiply(a, b int) int {
	return a * b
}

func add(a, b int) {
	fmt.Println(a + b)
}
