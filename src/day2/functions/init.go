package main

import "fmt"

var (
	a int
)

func init() {
	fmt.Println("Calling init")
	a = 100
}
func init() {
	fmt.Println("Calling init1")
}
func init() {
	fmt.Println("Calling init2")
}
func init() {
	fmt.Println("Calling init3")
}

func main() {
	fmt.Println("Demo init")
}
