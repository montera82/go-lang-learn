package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	v := Hello()

	if v != "hello" {
		t.Errorf("Expected hello but go %s", v)
	}
	//fmt.Println(hello())
}
