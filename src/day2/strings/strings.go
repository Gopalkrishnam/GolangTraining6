package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings package...")

	// Case conversion

	str := "welcome to go training"

	fmt.Println("Upper case", strings.ToUpper(str))

	fmt.Println("Lower case", strings.ToLower(str))

	fmt.Println("title", strings.ToTitle(str))

	path := "home/user/bin"

	split := strings.Split(path, "/")

	for _, v := range split {
		fmt.Println(v)
	}
	// fmt.Println("lower spl:", strings.ToLowerSpecial(str))

	// fmt.Println("upper spl", strings.ToUpperSpecial(str))
}
