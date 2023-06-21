package main

import (
	"flag"
	"fmt"
)

/*
	message
	count

	./main -message hello -count 4
*/

func main() {
	fmt.Println("Commandline tool")

	fmt.Println("Reading arguments")
	// fileterning out first argument as it is name of executable
	var message string
	var count int
	flag.StringVar(&message, "message", "", "We will print your message")
	flag.IntVar(&count, "count", 0, "Specify count")
	flag.Parse()
	if message == "" && count == 0 {
		flag.Usage()
		return
	}

	fmt.Println("Message is ", message, " count is", count)

	for i := 0; i < count; i++ {
		fmt.Println(message)
	}

}
