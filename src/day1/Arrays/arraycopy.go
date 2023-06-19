package main

import "fmt"

func main() {
	fmt.Println("Demo array copy")

	arraySrc := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arraySrc)

	arrayDst := arraySrc

	fmt.Println(arrayDst)

	arrayDst[0] = 100

	fmt.Println(arraySrc, arrayDst)

	arraydstref := &arraySrc

	fmt.Println(arraySrc, arraydstref)

	arraydstref[0] = 1000

	fmt.Println(arraydstref)
	fmt.Println(arraySrc)

	arraySrc[1] = 3000

	fmt.Println(arraydstref)
	fmt.Println(arraySrc)

}
