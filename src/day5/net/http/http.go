package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleRootGet(res http.ResponseWriter, req *http.Request) {
	fmt.Println("in handle RootGet")
}

func handleRootPost(res http.ResponseWriter, req *http.Request) {
	fmt.Println("In handleRoot Post")
}
func handleRootRoute(res http.ResponseWriter, req *http.Request) {
	fmt.Println("root is hit...")
	res.Write([]byte("Root route is hit..."))
	fmt.Println("Method hit is ", req.Method)
	fmt.Println(req.Header)
	res.Write([]byte(`<html><body>
	This content is from web browser
	</body></html>`))
	switch req.Method {
	case "GET":
		handleRootGet(res, req)
	case "POST":
		handleRootPost(res, req)
	case "PUT":
	case "DELETE":
	}
	res.WriteHeader(http.StatusNotFound)

}
func handleEnqiry(res http.ResponseWriter, req *http.Request) {
	fmt.Println("In handle Enquiry")
}
func main() {
	fmt.Println("http server demo..")
	http.HandleFunc("/", handleRootRoute)
	http.HandleFunc("/enquiry", handleEnqiry)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("received error while creating server ", err)
		os.Exit(0)
	}
}
