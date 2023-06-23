package transaction

import (
	"fmt"
	"time"
)

type Transaction struct {
	TID       int
	BookID    int
	SubID     int
	IssueDate time.Time
	DueDate   time.Time
}

var (
	transaction = make([]Transaction, 0, 10)
)

func IssueBook(BkID, SubID int) {
	tid := 0
	if len(transaction) == 0 {
		tid = 1
	} else {
		tid = transaction[len(transaction)-1].TID + 1
	}
	issueDate := time.Now()
	dueDate := issueDate.Add(15 * 24 * time.Hour)
	trans := Transaction{TID: tid, BookID: BkID, SubID: SubID, IssueDate: issueDate, DueDate: dueDate}
	transaction = append(transaction, trans)
}

func GetAll() []Transaction {
	trans := transaction
	return trans
}

func Get(tid int) (tr Transaction, e error) {
	for _, t := range transaction {
		if tid == t.TID {
			tr = t
			return
		}
	}
	e = fmt.Errorf("transaction not found")
	return
}
