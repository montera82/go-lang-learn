package main

import "fmt"

func main() {

  fmt.Println("Enter a number");

  number := 0
  
  fmt.Scanf("%f", number)
	fmt.Println(half(number))
}

func half(number int) (response int, ok bool) {

	if number%2 == 0 {
		response = number / 2
		ok = true

		return
	} else {
		response = 1
		ok = false

		return
	}

}
