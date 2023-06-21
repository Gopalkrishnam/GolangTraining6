package main

import "fmt"

func main() {
	doc := Document{Author: "Mavenir", Content: `this document is test document
	sjjhjsdhjsd kkjahjshdk asjdhk`}

	sheet := SpreadSheet{rows: 10, column: 5, Author: "Mavenir Systems"}
	PageSetupObject(doc)

	PreviewObject(doc)

	PrintObject(doc)

	PageSetupObject(sheet)

	PreviewObject(sheet)

	PrintObject(sheet)

	fmt.Println("Printing contents for doc and sheet")
	fmt.Println(doc)

	fmt.Println(sheet)

}
