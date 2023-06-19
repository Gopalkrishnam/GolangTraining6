package main

import "fmt"

func main() {
	fmt.Println("Demo Arrays...")
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)

	var array2 [5]int32 = [5]int32{0, 1, 2, 3, 4}
	fmt.Println(array2)

	var array3 = [5]int16{1, 2}

	fmt.Println(array3)

	var array4 = [5]int{0: 100, 4: 900, 3: 50, 2: 10, 1: 5}

	fmt.Println(array4)

	var array5 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100: 500}

	fmt.Println(array5)

	fmt.Println(len(array5))

	fmt.Println("Value at index 101", array5[101])

}
