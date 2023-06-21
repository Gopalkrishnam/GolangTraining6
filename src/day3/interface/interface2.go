package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type Object struct {
}

func (o Object) Read() {
	fmt.Println("In read...")
}

func (o Object) Write() {
	fmt.Println("In Write")
}

func main() {
	fmt.Println("Demo Embeding of interface")

	var r Reader = Object{}

	var w Writer = Object{}

	var rw ReadWriter = Object{}

	r.Read()

	w.Write()

	rw.Read()

	rw.Write()
}
