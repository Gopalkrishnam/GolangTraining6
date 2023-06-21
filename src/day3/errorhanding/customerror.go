package main

import "fmt"

type CustorError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e CustorError) Error() string {
	return fmt.Sprintf("Received Error Code :%d\n Received Error Message %s", e.ErrorCode, e.ErrorMessage)
}

func Divide(a, b int) (res int, err error) {
	if a == 0 {
		err = CustorError{ErrorCode: 1, ErrorMessage: "Invalid input"}
		return
	}
	if b == 0 {
		err = CustorError{ErrorCode: 2, ErrorMessage: "Divide by zero"}
		return
	}
	res = a / b
	return
}

func main() {
	fmt.Println("Custom error")

	res, err := Divide(0, 10)
	if err != nil {
		fmt.Println(err)
		c := err.(CustorError)
		fmt.Println(c.ErrorCode, c.ErrorMessage)
	} else {
		fmt.Println(res)
	}

	res, err = Divide(10, 0)
	if err != nil {
		fmt.Println(err)
		c := err.(CustorError)
		fmt.Println(c.ErrorCode, c.ErrorMessage)
	} else {
		fmt.Println(res)
	}

	res, err = Divide(10, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
