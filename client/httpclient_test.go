package client

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	h := NewHttpClient("GET", "http://www.huawei.com", nil, nil)
	data, err := h.Send()
	if err != nil {
		t.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(data)
}

func BenchmarkSend(b *testing.B) {
	h := NewHttpClient("Get", "https://baidu.com", nil, nil)
	for i := 0; i < b.N; i++ {
		h.Send()
	}

}
