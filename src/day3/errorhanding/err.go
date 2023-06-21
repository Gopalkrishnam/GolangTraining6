package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("Received system error")
		err = fmt.Errorf("Divide by zero, can not proceed further..%w", err)
		return 0, err
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling creating error")

	res, err := Divide(10, 0)

	if err != nil {
		fmt.Println("We have received error", err)
	} else {
		fmt.Println("Result is ", res)
	}

	res, err = Divide(100, 5)

	if err != nil {
		fmt.Println("We have received error", err)
	} else {
		fmt.Println("Result is ", res)
	}
}
