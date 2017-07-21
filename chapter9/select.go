package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "From 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "From 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Println("Message 1", msg)
			case msg := <-c2:
				fmt.Println("Message 2", msg)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
