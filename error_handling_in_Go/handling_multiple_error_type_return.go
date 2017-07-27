package main

import (
	"errors"
	"math/rand"
)

var ErrTimeout error = errors.New("Time out error")
var ErrRejected error = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(32))

func main() {

	_, err := sendRequest()

	for err == ErrTimeout {
		println("timeout...retrying")
		_, err = sendRequest()

	}

	if err != nil {
		println("request was rejected")
	}
	println("request was successful !")

}

func sendRequest() (string, error) {

	switch random.Int() % 3 {
	case 0:
		return "", ErrTimeout
	case 1:
		return "", ErrRejected
	default:
		return "Successful", nil

	}

}
