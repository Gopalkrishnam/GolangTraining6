package main

import "fmt"

func echo() {
	fmt.Println("echo has been called...")
}

func main() {
	fmt.Println("Demo function")

	echo()

	echo()

	echo()

	printMessage("Hello")
	printMessage("Hi")
	printMessage("Welcome")
}

func printMessage(msg string) {
	fmt.Println("Message to be printed is", msg)
}
