package main

import "fmt"

type promptedfields struct {
	int
	string
	int16
	float32
}

func main() {
	fmt.Println("Demo:Prompted Fields")
	var a promptedfields = promptedfields{int: 1, string: "abc", int16: 0, float32: 3.14159}

	fmt.Println(a.int)
	fmt.Println(a.string)
	fmt.Println(a)
}
