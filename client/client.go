package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Send() error
}

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

func (h *HttpClient) Send() error {
	req, err := http.NewRequest(h.method, h.url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp.ContentLength)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("request status is not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	//	var data interface{}

	//json.Unmarshal(body, &data)
	fmt.Println(string(body))
	return nil
}
