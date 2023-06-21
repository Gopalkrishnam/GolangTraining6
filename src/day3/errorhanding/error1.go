// package main

// import "fmt"

// func Divide(a, b int) int {
// 	fmt.Println("In Divide")
// 	if b == 0 {
// 		fmt.Println("throwing panic")
// 		panic("Divide by zero....")
// 	}
// 	return a / b
// }

// func main() {

// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("We have received error")
// 		}
// 	}()
// 	fmt.Println("Error handing")

// 	Divide(10, 0)

// 	Divide(10, 100)
// }

package main

import "fmt"

func Divide(a, b int) int {
	fmt.Println("In Divide")
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println(rec)
		}
		fmt.Println("In defer")
	}()
	return a / b
}

func main() {

	fmt.Println("Error handing")

	Divide(10, 0)

	Divide(10, 100)
}
