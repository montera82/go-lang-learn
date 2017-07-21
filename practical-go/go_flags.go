package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name"  default:"world" description:"a name to say hello to"`
	Spanish bool   `short:"s" long:"spanish" description:"spanish lanquage option"`
}

func main() {
	_, err := flags.Parse(&opts)

	if nil != err {
		panic(err)
		os.Exit(1)
	}

	switch opts.Spanish {
	case false:
		fmt.Printf("Hello %s ! \n", opts.Name)
	default:
		fmt.Printf("Holla %s ! \n", opts.Name)

	}
}
