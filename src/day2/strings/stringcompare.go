package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Demo string compare")

	str1 := "hello111"

	str2 := "hello"

	if str1 == str2 {
		fmt.Println("Equal...")
	} else {
		fmt.Println("Not equal")
	}

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("Equal..")
	} else {
		fmt.Println("Not equal..")
	}

	if reflect.DeepEqual(str1, str2) {
		fmt.Println("Equal ......")
	} else {
		fmt.Println("Not equal.......")
	}


}
