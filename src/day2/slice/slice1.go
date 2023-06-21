package main

import "fmt"

func main() {
	fmt.Println("Demo slice copy")
	var slicesrc = make([]int, 5, 6)

	for i := 0; i < len(slicesrc); i++ {
		slicesrc[i] = i
	}

	slicedst := slicesrc

	fmt.Println(slicesrc)
	fmt.Println(len(slicesrc))
	fmt.Println(cap(slicesrc))
	fmt.Println(slicedst)
	fmt.Println(len(slicedst))
	fmt.Println(cap(slicedst))

	fmt.Println("Append to slicesrc")
	slicesrc = append(slicesrc, 5)

	fmt.Println(slicesrc)
	fmt.Println(len(slicesrc))
	fmt.Println(cap(slicesrc))
	fmt.Println(slicedst)
	fmt.Println(len(slicedst))
	fmt.Println(cap(slicedst))

	//deleting element from slice

	idx := 3

	slicesrc = append(slicesrc[:idx], slicesrc[idx+1:]...)
	fmt.Println("After delete")
	fmt.Println(slicesrc)

}
