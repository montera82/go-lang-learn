package main

import (
	"fmt"
	hermes "golang-book/unit_tests_in_go/fake_external_library"
)

func main() {

	m := hermes.Message{
		Email:   "noc@test.com",
		Subject: "sample subject",
	}
	problem := []byte("Critical error!")

	Alert(m, problem)
}

func Alert(m hermes.Message, problem []byte) {

	resp, _ := m.Send()
	fmt.Println(resp)
}
