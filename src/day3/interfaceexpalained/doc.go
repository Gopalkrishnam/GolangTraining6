package main

import "fmt"

type Document struct {
	Author  string
	Content string
}

func (d Document) Print() {
	fmt.Println("Printing Document")
	fmt.Println("Author is :", d.Author)
	fmt.Println("Content is:", d.Content)
}

func (d Document) PageSetup() {
	fmt.Println("Page Setup")

}

func (d Document) String() string {
	return fmt.Sprintf("Document details are \n Author:%s \n Content:%s", d.Author, d.Content)
}

func (d Document) Preview() {
	fmt.Println("We are showining preview")
	fmt.Println("Author is :", d.Author)
	fmt.Println("Content is:", d.Content)
}
