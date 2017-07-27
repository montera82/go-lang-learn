package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("Can't divide by zero")

func main() {

	println("Divide 1 by 0")
	_, err := preDevideCheck(1, 0)

	if nil != err {
		fmt.Println(err)

	}

	println("Divide 2 by 0")
	divide(2, 0)

}

func preDevideCheck(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
