package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		// error case
		return 0, false
	}
	c := a / b
	return c, true
}

func main() {
	fmt.Println("Demo: Divide")

	// fmt.Println(Divide(10, 20))
	// fmt.Println(Divide(100, 10))

	// fmt.Println(Divide(100, 0))

	res, eval := Divide(10, 20)
	if eval {
		fmt.Println(res)
	} else {
		fmt.Println("Exception")
	}

	res, eval = Divide(100, 10)
	if eval {
		fmt.Println(res)
	} else {
		fmt.Println("Exception")
	}

	res, eval = Divide(100, 0)
	if eval {
		fmt.Println(res)
	} else {
		fmt.Println("Exception")
	}

	area, circ := Circle(10)
	fmt.Println(area, circ)
}

func Circle(radius int) (float32, float32) {
	area := 3.14159 * float32(radius) * float32(radius)
	circ := 2 * 3.14159 * float32(radius)

	return area, circ
}
