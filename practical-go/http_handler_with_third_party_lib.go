package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":9009", router))
}

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	fmt.Fprint(res, "Welcome")
}

func Hello(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	fmt.Fprintf(res, "Hello %s", params.ByName("name"))
}
