package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Varaible...")

	var a int = 100

	fmt.Println(a)

	var b = 100

	fmt.Println(b)

	fmt.Printf("%T\n", b)

	fmt.Println("type of variable b is", reflect.TypeOf(b))
}
