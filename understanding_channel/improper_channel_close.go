package main

import (
	"time"
)

func main() {

	until := time.After(5 * time.Second)
	msg := make(chan string)

	go send(msg)

	for {
		select {
		case msg := <-msg:
			println(msg)
		case <-until:
			close(msg)
			println("closed and timeout")

			time.Sleep(500 * time.Millisecond)

			return
		}
	}

}

func send(c chan<- string) {

	for {
		c <- "hello"
		time.Sleep(500 * time.Millisecond)
	}

}
