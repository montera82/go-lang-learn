package main

import (
	"time"
)

func main() {

	until := time.After(5 * time.Second)
	msg := make(chan string)
	done := make(chan bool)

	go send(msg, done)

	for {
		select {
		case msg := <-msg:
			println(msg)
		case <-until:
			done <- true
			println("timeout")

			time.Sleep(500 * time.Millisecond)

			return
		}
	}

}

func send(c chan<- string, done <-chan bool) {

	for {
		select {
		case <-done:
			println("done")
			close(c)
			return
		default:
			c <- "hello"
			time.Sleep(500 * time.Millisecond)
		}

	}

}
