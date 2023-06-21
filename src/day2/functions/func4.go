package main

import "fmt"

func Circle(radius float32) (area float32, circ float32) {

	area = 3.14159 * radius * radius
	circ = 2 * 3.14159 * radius
	return
}

func main() {
	fmt.Println("Demo named return")

	area, circ := Circle(100)
	fmt.Println(area, circ)

	area, _ = Circle(10)

	_, circ = Circle(10)

	fmt.Println(area, circ)
}
