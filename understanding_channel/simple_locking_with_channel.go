package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool)

	for i := 0; i < 1; i++ {
		go worker(i, lock)
	}

	time.Sleep(10 * time.Second)

}

func worker(i int, lock chan bool) {
	fmt.Printf("%d wants the lock \n", i)
	lock <- true
	fmt.Printf("%d has the lock \n", i)

	time.Sleep(time.Millisecond * 500)

	fmt.Printf("%d is about to release the lock\n", i)
	<-lock

	return
}
