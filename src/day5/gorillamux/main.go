package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunctionGET(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("this server uses Gorilla mux GET"))
}

func handleFunctionPOST(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("this server uses Gorilla mux POST....."))
}

func main() {
	fmt.Println("Gorilla mux...")
	r := mux.NewRouter()
	r.HandleFunc("/", handleFunctionGET).Methods("GET")
	r.HandleFunc("/", handleFunctionPOST).Methods("POST")
	http.ListenAndServe(":1020", r)
}
