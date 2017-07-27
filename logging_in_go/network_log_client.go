package main

import (
	"log"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "localhost:1902")

	if err != nil {
		println(err)
		return
	}
	defer con.Close()

	format := log.Ldate | log.Lshortfile
	logger := log.New(con, "example ", format)

	//should there be a panic downstream, this closure will be activated,
	//and the reason for the panic logged + added benefit of program not crashing
	//although irrelevant in this case, since main thread will finish executing any way
	defer func() {
		if r := recover(); r != nil {
			logger.Println("Recovered from panic: cause", r)
		}
	}()

	logger.Println("This is a message")
	logger.Panicln("This is a panic message")

}
