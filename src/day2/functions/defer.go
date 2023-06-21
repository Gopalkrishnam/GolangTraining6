package main

import "fmt"

func printMessageDefered(msg string) {
	defer fmt.Println("Deferring this print ", msg)
	fmt.Println("Hello...")
}

func main() {
	fmt.Println("demo: defer")

	// defer fmt.Println("I want to defer this statement")

	// fmt.Println("After defer statement")

	// defer fmt.Println("This is my second defer")

	defer printMessageDefered("Good Morning")

	fmt.Println("Message defered")

	defer printMessageDefered("Good Evening")

}

/*
demo: defer
Message defered
Hello
Good Evening
Hello
Good Morning
*/
