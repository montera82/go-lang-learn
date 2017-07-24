package main

import (
	"time"
)

func main() {

	until := time.After(400 * time.Millisecond)
	msg := make(chan bool)

	go send(msg)

	for i := 0; i < 20; i++ {
		select {
		case <-msg:
			println("Received message")
		case <-until:
			println("closed and timeout")
			return
		default:
			println("**yawn**")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(c chan<- bool) {

	time.Sleep(120 * time.Millisecond)

	c <- true
	close(c)
	println("sent and closed")

}
