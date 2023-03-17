package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


type FailedReqFunc func() bool

type HttpClient struct {
	method string
	url    string
	data   map[string]interface{}
	header map[string]string
}

func NewHttpClient(method string, url string, data map[string]interface{}, header map[string]string) *HttpClient {
	return &HttpClient{
		method: method,
		url:    url,
		data:   data,
		header: header,
	}
}

func (h *HttpClient) Send() (interface{}, error) {
	req, err := http.NewRequest(h.method, h.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	//fmt.Println(resp.ContentLength)
	// 判断err放后面
	//	没有这个 当resp重定向错误时 resp 不为nil
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil, fmt.Errorf("request status is not 200")
	}
	// todo:是否加内容判断

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return body, nil
}
