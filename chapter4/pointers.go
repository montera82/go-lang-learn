package main

import "fmt"

func main() {

	number := new(int)

	*number = 50

	fmt.Println(*number)

	zero(number)

	// &number is how i read the pointer of a variable
	// new keyword can also be used
	fmt.Println(*number)
}

func zero(number *int) {

	//dereferencing here i.e i am saying put 0 as a value into number's location ( number here being a point)
	*number = 0
}
