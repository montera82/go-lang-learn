package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Eduardo"
	}

	fmt.Fprintf(res, "Hello %s", name)
}

func Goodbye(res http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path, "/")

	name := path[2]

	if name == "" {
		name = "Eduardo"
	}

	fmt.Fprintf(res, "Hello %s", name)
}

func Homepage(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	fmt.Fprint(res, "Welcome to the homepage")
}
