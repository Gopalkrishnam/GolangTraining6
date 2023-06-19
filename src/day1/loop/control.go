package main

import "fmt"

func main() {
	fmt.Println("Demo: Loop control statements")

testlabel:
	fmt.Println("Here my go to starts")
	i := 0
	for {
		if i == 100 {
			break
		}
		i++
		if i%2 != 0 {
			goto testlabel
		}
		fmt.Println("Value of i is", i)

	}
}
