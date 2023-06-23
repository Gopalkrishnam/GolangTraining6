package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library/books"
	"net/http"
	"os"
)

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Root")
	res.Write([]byte("Welcome to libray"))
}

func handleBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Book")
	switch req.Method {
	case "GET":
		data, _ := ioutil.ReadAll(req.Body)
		mp := make(map[string]interface{})
		fmt.Println(string(data))
		err := json.Unmarshal(data, &mp)
		fmt.Println(err)
		fmt.Println(mp)
		fmt.Println(string(data))
		fmt.Println(req.GetBody)
		bk, err := books.GetBook(int(mp["bookID"].(float64)))
		if err != nil {
			res.Write([]byte("record not found"))
			res.WriteHeader(http.StatusNotFound)
		} else {
			data, _ = json.Marshal(bk)
			res.Write(data)
		}
		//res.WriteHeader(http.StatusOK)
	case "POST":
		books.InsertBook("Abc", "xyz", 102)

	}
}

func handleSubs(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Subs")
}
func handleTransaction(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Transaction")
}

func main() {
	fmt.Println("Library....")
	http.HandleFunc("/", handle)
	http.HandleFunc("/book", handleBook)
	http.HandleFunc("/subs", handleSubs)
	http.HandleFunc("/transaction", handleTransaction)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Received error", err)
		os.Exit(0)
	}
}
