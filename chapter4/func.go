package main

import "fmt"

func main() {

	collection := []float64{98, 93, 77, 82,
		83}

	if result, ok := average(collection...); ok {
		fmt.Println(result)
	}

}

func average(collection ...float64) (r float64, ok bool) {

	var total float64

	for _, v := range collection {
		total += v
	}

	return total / float64(len(collection)), true
}
