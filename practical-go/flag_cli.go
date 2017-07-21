package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "supply to specify the name to greet")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "use spanish language")
	flag.BoolVar(&spanish, "s", false, "use spanish language")

}

func main() {

	flag.Parse()

	switch spanish {
	case true:
		fmt.Printf("Holla %s! \n", *name)
	default:
		fmt.Printf("Hello %s! \n", *name)

	}
}
