package main

import "fmt"

func main() {
	fmt.Println("Demo slice copy")

	var slicesrc []int = []int{1, 2, 3, 4, 5}

	slicedst := slicesrc

	fmt.Println(slicesrc)
	fmt.Println(len(slicesrc))
	fmt.Println(cap(slicesrc))
	slicesrc = append(slicesrc, 6)
	fmt.Println("after append", slicesrc)
	fmt.Println(len(slicesrc))
	fmt.Println(cap(slicesrc))
	fmt.Println(slicedst)
	fmt.Println(len(slicedst))
	fmt.Println(cap(slicedst))
}
