package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homepage)

	http.ListenAndServe(":8888", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Eduardo"
	}

	fmt.Fprintf(res, "Hello %s", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path, "/")

	name := path[2]

	if name == "" {
		name = "Eduardo"
	}

	fmt.Fprintf(res, "Hello %s", name)
}

func homepage(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	fmt.Fprint(res, "Welcome to the homepage")
}
