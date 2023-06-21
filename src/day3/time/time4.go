package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time package")

	currTime := time.Now()
	fmt.Println("Finding time since")
	d := time.Since(currTime)

	fmt.Println(d)
	t1 := time.Date(2024, 06, 21, 14, 23, 0, 0, &time.Location{})
	d = time.Until(t1)

	fmt.Println(d)
}
