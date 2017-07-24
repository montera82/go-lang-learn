package main

import (
	"os"
	"time"
)

func main() {

	doneCh := time.After(time.Second * 5)
	echoCh := make(chan []byte)

	go readStdin(echoCh)

	for {
		select {
		case <-doneCh:
			println("\ntimeout")
			os.Exit(0)
		case buf := <-echoCh:
			os.Stdout.Write(buf)

		}
	}
}

func readStdin(c chan<- []byte) {

	data := make([]byte, 1024)
	i, _ := os.Stdin.Read(data)

	if i > 0 {
		c <- data
	}
}
