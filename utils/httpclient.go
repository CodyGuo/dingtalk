package utils

import (
	"bytes"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
	Url    string
}

func NewHttpClient(url string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{Timeout: timeout},
		Url:    url,
	}
}

func (c *HttpClient) Request(method string, header map[string]string, body []byte) (response *http.Response, err error) {
	req, err := http.NewRequest(method, c.Url, bytes.NewReader(body))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	return c.client.Do(req)
}
