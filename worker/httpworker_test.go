package worker

import (
	"fmt"
	"testing"
)

func TestCheckInput(t *testing.T) {
	input := HttpInput{}
	a := CheckHttpInput(input)
	if !a {
		fmt.Println("no")
		t.Fail()
	}
}
