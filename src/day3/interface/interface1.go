package main

import "fmt"

type FirstInterface interface {
	Compute()
	xyz()
}

type SecondInterface interface {
	Compute()
}

type StructImpl struct {
}

func (s StructImpl) Compute() {
	fmt.Println("We are in compute method of StructImpl")
}

func (s StructImpl) xyz() {

}

func main() {
	fmt.Println("Multiple Interface")

	var f FirstInterface = StructImpl{}

	var s SecondInterface = StructImpl{}

	f.Compute()

	s.Compute()

}
