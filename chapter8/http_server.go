package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "text/html")

	io.WriteString(res,
		"<html><body>Hello, world</body></html>",
	)
}

func main() {

	fmt.Println("Service listening on port 9000")
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9000", nil)

}
