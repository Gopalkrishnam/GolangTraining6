package main

import (
	"fmt"
	"training/emp"
	"training/math"
	"training/math/advanced"

	cl "github.com/fatih/color"
)

func main() {
	fmt.Println(math.Add(10, 20))
	fmt.Println(advanced.Square(10))

	var e = emp.New(10, "Anand")
	e.Name = "Kumar"
	e.EID = 10
	fmt.Println(e)
	cl.Red("Hello")
	cl.Green("Hi")
}
