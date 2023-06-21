package main

import "fmt"

func main() {

	fmt.Println("Demo: Maps")

	var mymap map[string]int = map[string]int{"A": 100, "B": 200, "C": 300}

	v := mymap["A"]

	fmt.Println(v)
	fmt.Println(mymap)

	mymap["D"] = 1000

	mymap["A"] = 10000

	fmt.Println(mymap)

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	// deleting element from map
	delete(mymap, "C")
	fmt.Println("After deleting C")
	for key, value := range mymap {
		fmt.Println(key, value)
	}

	for key := range mymap {
		delete(mymap, key)
	}

	fmt.Println(mymap)

	mymap = map[string]int{}
}
