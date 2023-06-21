package main

/*
	{
		"name": "anand",
		"age": 35,
		"address": "HSR Banaglore"
	}
*/

import "fmt"

func main() {
	fmt.Println("Demo: strings")

	var firstStr string = "Welcome \r\nto golang \ntraining"

	var secondstr string = `hello and welcome to go training batch 5 "This is very intresting subject
	We have wonderful concepts to learn !@##$%^^"`

	fmt.Println(secondstr)

	fmt.Println(firstStr)

	jsonstr := "{\n\t\"name\": \"anand\",\n\t\"age\": 35,\n\t\"address\": \"HSR Banaglore\"\n}"

	fmt.Println(jsonstr)

	jsonwithbacktick := `{
		"name": "anand",
		"age": 35,
		"address": "HSR Banaglore"
	}`

	fmt.Println(jsonwithbacktick)
}
