package net

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Requester interface {
	DoRequest(req *http.Request) (*http.Response, error)
	SetHeaders(headers map[string]string) Requester
	ReadBody(body io.Reader) ([]byte, error)
}

type customRequest struct {
	Headers map[string]string
}

func NewHttpRequest() Requester {
	return &customRequest{}
}

func (r *customRequest) SetHeaders(headers map[string]string) Requester {
	r.Headers = headers
	return r
}

func (r customRequest) DoRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}

	for i, v := range r.Headers {
		request.Header.Set(i, v)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return nil, nil
}

func (r customRequest) ReadBody(body io.Reader) ([]byte, error) {
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
