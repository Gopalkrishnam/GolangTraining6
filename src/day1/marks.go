package main

import "fmt"

func main() {
	fmt.Println("Demo: marks")

	var s1, s2, s3 int

	fmt.Println("Please enter marks in s1")
	fmt.Scanf("%d", &s1)
	fmt.Println("Please enter marks in s2")
	fmt.Scanf("%d", &s2)
	fmt.Println("Please enter marks in s3")
	fmt.Scanf("%d", &s3)

	var total = s1 + s2 + s3

	fmt.Println("total marks", total)

	//var avg = float64(total) / float64(3)

	var avg = float64(total / 3)

	fmt.Println(avg)
}
