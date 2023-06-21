package main

import (
	"fmt"
	"strconv"
)

func MyCustomPrintln(format string, inputs ...interface{}) {
	fmt.Println(format, inputs)

	s := format

	for _, val := range inputs {
		switch val.(type) {
		case int:
			s += strconv.Itoa(val.(int))

		case string:
			s += val.(string)
		}
	}

	fmt.Println(s)

}

func main() {

	fmt.Println("Variadic")

	MyCustomPrintln("Hello", 1, "hi", "xyz", 10)
}
