package main

import "fmt"

func main() {

	var a = 100
	var b = 100

	var c = 500

	var res = (a == b) && (a == c)

	fmt.Println(res)

	res = (a == b) || (a == c)

	fmt.Println(res)

	res = !(a == b)

	fmt.Println(res)

	
}
