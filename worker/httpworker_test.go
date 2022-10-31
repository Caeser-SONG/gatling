package worker

import (
	"fmt"
	"testing"
)

func TestCheckInput(t *testing.T) {
	//input := HttpInput{}
	input := &HttpInput{
		Method: "GET",
		Url:    "http://www.baidu.com",
		Header: nil,
		Data:   nil,
	}
	a := CheckHttpInput(input)
	fmt.Println(input)

	if !a {
		fmt.Println("no")
		t.Fail()
	}
}
