package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {

	handler := newHander()

	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt, os.Kill)

	go listenForShutdown(ch)

	go manners.ListenAndServe(":8089", handler)

	var r string
	fmt.Scanln(&r)

}

type handler struct{}

func newHander() *handler {

	return new(handler)
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()

	name := query.Get("name")

	if name == "" {
		name = "Edwin"
	}

	fmt.Fprint(res, "Hello", name)
}

func listenForShutdown(ch <-chan os.Signal) {

	<-ch
	manners.Close()
}
