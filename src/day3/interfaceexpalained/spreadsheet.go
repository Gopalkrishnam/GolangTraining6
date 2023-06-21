package main

import "fmt"

type SpreadSheet struct {
	rows, column int
	Author       string
}

func (s SpreadSheet) Print() {

	fmt.Println("We are printing spread sheet")
	fmt.Println("Rows, Column:", s.rows, s.column)
	fmt.Println("Author is ", s.Author)

}

func (s SpreadSheet) String() string {
	return fmt.Sprintf("Contents of Spread sheet are\nRows:%d\n Column:%d\nAuthor:%s", s.rows, s.column, s.Author)
}

func (s SpreadSheet) Preview() {
	fmt.Println("In Preview")
}

func (s SpreadSheet) PageSetup() {
	fmt.Println("In Page Setup")
}
