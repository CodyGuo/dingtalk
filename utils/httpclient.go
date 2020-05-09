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

func (c *HttpClient) GetURL() string {
	return c.Url
}

func (c *HttpClient) SetURL(url string) {
	c.Url = url
}

func (c *HttpClient) GetTimeout() time.Duration {
	return c.client.Timeout
}

func (c *HttpClient) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
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
