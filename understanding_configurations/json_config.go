package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("config.json")

	if nil != err {
		panic(err)
	}

	defer file.Close()
	conf := new(config)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(conf)

	if nil != err {
		panic(err)
	}

	fmt.Println(conf.Path)
}
