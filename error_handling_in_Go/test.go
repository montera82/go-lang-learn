package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println(msg)

	panic(errors.New("Ouch..panic at the end of the main function"))

	msg := "hooss"
}
