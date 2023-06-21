package main

import "fmt"

type Obj struct {
}

func main() {

	fmt.Println("Demo empty interface...")

	// var dymanic interface{}

	// dymanic = "Hello"

	// dymanic = 10

	// dymanic = 10.5

	// dymanic = Obj{}

	Add(10, 20)

	Add("Hi", "Hello")

	Add(10.1, 15.5)

	Add(10, 15.5)
}

func Add(a interface{}, b interface{}) {
	// areferred := a.(int)
	// breferred := b.(int)

	// fmt.Println(areferred + breferred)
	fmt.Println(a, b)

	switch a.(type) {
	case int:
		areferred := a.(int)
		breferred := b.(int)
		fmt.Println(areferred + breferred)
	case float32:
		areferred := a.(float32)
		breferred := b.(float32)
		fmt.Println(areferred + breferred)

	case float64:

		areferred := a.(float64)
		breferred := b.(float64)
		fmt.Println(areferred + breferred)
	case string:
		areferred := a.(string)
		breferred := b.(string)
		fmt.Println(areferred + breferred)
	case Obj:
	}
}
