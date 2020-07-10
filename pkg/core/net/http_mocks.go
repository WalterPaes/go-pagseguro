package net

import "net/http"

type ConnectorMock struct{}

func (c ConnectorMock) DoGet() (http.Response, error) {
	response := http.Response{
		StatusCode: 200,
	}
	return response, nil
}

func (c ConnectorMock) DoPost() (http.Response, error) {
	response := http.Response{
		StatusCode: 200,
	}
	return response, nil
}
