package books

import "fmt"

type Books struct {
	BookID int    `json:"bookId"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

var (
	books = make([]Books, 3, 10)
)

func init() {
	books[0] = Books{BookID: 1, Title: "Golang", Author: "Google", Price: 100}
	books[1] = Books{BookID: 2, Title: "Golang Book1", Author: "Google Author1", Price: 101}
	books[2] = Books{BookID: 3, Title: "Golang Book 2", Author: "Google Author2", Price: 102}
}

func GetAllBooks() []Books {
	lstBooks := books
	return lstBooks
}

func GetBook(bookId int) (b Books, e error) {
	for _, b := range books {
		if b.BookID == bookId {
			return b, nil
		}
	}
	e = fmt.Errorf("book not found")
	return
}

func DeleteBook(bookId int) error {

	for i, b := range books {
		if bookId == b.BookID {
			books = append(books[:i], books[i:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func InsertBook(title, author string, price int) int {
	maxBookID := books[len(books)-1].BookID
	maxBookID++
	bk := Books{BookID: maxBookID, Title: title, Author: author, Price: price}

	books = append(books, bk)
	return maxBookID
}

func UpdateBook(bkid int, title, author string, price int) error {

	for i, b := range books {
		if bkid == b.BookID {
			books[i].Title = title
			books[i].Author = author
			books[i].Price = price
			return nil
		}
	}
	return fmt.Errorf("book not found")
}
