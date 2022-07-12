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

func BenchmarkSend(b *testing.B) {
	h := NewHttpClient("Get", "https://baidu.com", nil, nil)
	for i := 0; i < b.N; i++ {
		h.Send()
	}

}
