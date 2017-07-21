package main

import "fmt"

func main() {

	nextEvenNumber := makeEvenGenerator()
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
	fmt.Println(nextEvenNumber())
}

func makeEvenGenerator() func() uint {

	retUint := uint(0)

	return func() uint {

		retUint += 2

		return retUint

	}

}
