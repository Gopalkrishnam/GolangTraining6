package main

import "net/http"

func handleRoot1(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Received request to server1"))
}

func CreateHttpService(addr string) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handleRoot1)
	go func() {
		http.ListenAndServe(addr, handler)
	}()
}

func handleRoot2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Received request to server2"))
}

func main() {

	CreateHttpService(":3040")
	http.HandleFunc("/", handleRoot2)
	http.ListenAndServe(":4050", nil)
}
