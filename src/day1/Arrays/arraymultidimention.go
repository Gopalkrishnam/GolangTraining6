package main

import "fmt"

func main() {
	fmt.Println("Multidimensional array...")

	var multiArray = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(multiArray)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(multiArray[i][j])
		}
	}

	for idx, val := range multiArray {
		//fmt.Println(idx, val)
		for i, v := range val {
			fmt.Println(idx, i, v)
		}
	}
}
