package main

import (
	"fmt"
	"os"
  "golang-book/chapter8/math"
)

func main() {

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Unable to open file", err)
		return
	}

	defer file.Close()

	//get the stats of the file

	stats, err := file.Stat()

	if err != nil {
		fmt.Println("Unable to read file stat", err)
		return
	}

	//read the file contents
	bs := make([]byte, stats.Size())

	i, err := file.Read(bs)

	if err != nil {
		fmt.Println("Unable to read file content", err)
		return
	}

	str := string(bs)

	fmt.Println(str, i)

  //Testing math package
  numbers := []float64 {4, 0.1, 1}
  fmt.Println(math.Average(numbers...))
}
