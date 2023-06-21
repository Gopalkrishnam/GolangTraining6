package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Short declaration operator :=")

	var a = 100

	b := 100
	fmt.Println("a and b are ", a, b)

	x, y := 10, 10.1

	fmt.Println(x, y)

	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))

}
