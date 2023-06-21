package main

import "fmt"

type MyFirstInterface interface {
	SayHello()
}

type FirstStrcut struct {
	Message string
}

func (f FirstStrcut) SayHello() {
	fmt.Println("Hello from first", f.Message)
}

type SecondStruct struct {
	Value int
}

func (s SecondStruct) SayHello() {
	fmt.Println("Hello from second", s.Value)
}

func main() {
	fmt.Println("Demo:Interface....")

	// var intf MyFirstInterface = FirstStrcut{Message: "Team Golang"}

	// intf.SayHello()

	var intf MyFirstInterface

	var first FirstStrcut

	var second SecondStruct

	intf = first
	intf.SayHello()

	intf = second
	intf.SayHello()

}
