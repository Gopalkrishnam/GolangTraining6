package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")

	var welcomestring = "Namste"

	switch welcomestring {
	case "Hello":
		fmt.Println("English")
	case "Namste":
		fmt.Println("Hindi")
	}
}
