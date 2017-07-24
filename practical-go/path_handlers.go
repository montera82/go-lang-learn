package main

import (
	"fmt"
	handlers "golang-book/practical-go/example_handlers"
	"log"
	"net/http"
	"path"
)

func main() {
	pr := newPathResolver()
	pr.add("GET /hello", handlers.Hello)
	pr.add("* /goodbye/*", handlers.Goodbye)
	pr.add("GET /", handlers.Homepage)

	fmt.Printf("staring server at localhost:9001..")
	log.Fatal(http.ListenAndServe(":9001", pr))
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {

	p := pathResolver{make(map[string]http.HandlerFunc)}

	return &p
	//return new(pathResolver)
}

func (p *pathResolver) add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	check := req.Method + " " + req.URL.Path

	fmt.Println(check)
	fmt.Println(p)

	for pattern, handler := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handler(res, req)

			return
		} else {
			fmt.Fprint(res, err)
		}

		http.NotFound(res, req)
		return

	}
}
