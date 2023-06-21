package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	mp := make(map[int]string)

	mp[1] = "Anand"
	mp[2] = "Vijay"
	mp[3] = "Anirudh"

	fmt.Println(mp)

	a, avl := mp[3]

	if avl {
		fmt.Println("element is available")
	} else {
		fmt.Println("Element is not available")
	}

	fmt.Println(a)

	mp = make(map[int]string)
}
