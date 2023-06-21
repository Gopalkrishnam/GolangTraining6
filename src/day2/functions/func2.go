package main

import "fmt"

// func Add(a, b int) {
// 	fmt.Println(a + b)
// }

func Add(a int, b int) int {
	c := a + b

	return c
}



func main() {
	fmt.Println("Demo Add")
	res := Add(10, 20)
	fmt.Println(res)
	res = Add(30, 40)
	fmt.Println(res)

	res = Add(Add(10, 20), Add(30, 40))
	fmt.Println(res)
}
