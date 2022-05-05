package client

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	h := NewHttpClient("GET", "http://www.huawei.com", nil, nil)
	e := h.Send()
	fmt.Println(e)
}
