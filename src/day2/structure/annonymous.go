package main

import "fmt"

func main() {
	var a = struct {
		int
		age int
		int8
		int16
		int32
		int64
		string
	}{int: 1, string: "Hello"}

	fmt.Println(a)

	fmt.Println(a.int, a.string)

	var b = struct {
		name string
		age  int
	}{name: "Anand", age: 40}

	fmt.Println(b.name, b.age)
	fmt.Println(b)
}
