package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Demo: string package...")
	str := "Hello "

	str1 := "and welcome"

	str3 := str + str1
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str3)

	str4 := fmt.Sprint(str, str1, str3)

	fmt.Println(str4)

	str = strings.ToUpper(str)

	a := "this is test code"

	b := 10

	c := 22.55

	str5 := fmt.Sprintf("%s_%d_%f", a, b, c)

	fmt.Println(str5)

	fmt.Println(str)
}
