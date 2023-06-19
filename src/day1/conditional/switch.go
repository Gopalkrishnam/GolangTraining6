package main

import "fmt"

func main() {

	var weekday int

	fmt.Println("Please enter weekday")

	fmt.Scanf("%d", &weekday)
	/*
		switch weekday {
		case 0:
			fmt.Println("Sunday")
		case 1:
			fmt.Println("Monday")
		case 2, 3, 4, 5:
			fmt.Println("Weekday..")
		case 6:
			fmt.Println("Weekend")
		default:
			fmt.Println("Unhandled")
		}
	*/
	switch weekday {
	case 6:
		fallthrough
	case 0:
		fmt.Println("Weekend")
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Weekdays")
	default:
		fmt.Println("Unhandled")

	}
}
