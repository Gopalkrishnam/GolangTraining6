package main

import "fmt"

func SliceByValue(slice []int) {
	fmt.Println("In SliceByValue")
	fmt.Println(slice)
	slice = append(slice, 10)
	slice = append(slice, 20)

}

func sliceByRef(slice *[]int) {
	fmt.Println("In SliceByValue")
	fmt.Println(*slice)
	fmt.Println("Before :", cap(*slice))
	fmt.Println("Before:", len(*slice))
	*slice = append(*slice, 10)
	*slice = append(*slice, 20)

	fmt.Println("After :", cap(*slice))
	fmt.Println("Afert:", len(*slice))

	for _, v := range *slice {
		fmt.Println(v)
	}

	for i := 0; i < len(*slice); i++ {
		fmt.Println((*slice)[i])
		(*slice)[i] = (*slice)[i] * 10
	}

	for i, v := range *slice {
		fmt.Println(v)
		(*slice)[i] = v * 100
	}

}

func main() {
	fmt.Println("Demo slice by reference")

	var slc = make([]int, 5, 5)
	fmt.Println(cap(slc))
	fmt.Println(len(slc))

	for i := 0; i < len(slc); i++ {
		slc[i] = i + 1
	}
	fmt.Println(slc)

	SliceByValue(slc)
	fmt.Println(slc)
	fmt.Println(cap(slc))
	fmt.Println(len(slc))

	sliceByRef(&slc)

	fmt.Println(slc)
	fmt.Println(cap(slc))
	fmt.Println(len(slc))

}
