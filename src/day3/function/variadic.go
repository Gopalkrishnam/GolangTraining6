package main

import "fmt"

func Add(a, b int, ints ...int) int {
	sum := a + b

	for _, v := range ints {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Demo Variadic functions")

	res := Add(10, 20)

	fmt.Println(res)

	res = Add(100, 200, 300, 4000)
	fmt.Println(res)

	res = Add(10, 20)
	fmt.Println(res)
	res = Add(10, 20, 30)
	fmt.Println(res)

}
