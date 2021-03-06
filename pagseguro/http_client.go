package pagseguro

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	client  *http.Client
	url     string
	headers map[string]string
	token   string
	body    interface{}
}

func NewHttpClient(url, token string, headers map[string]string) (*HttpClient, error) {
	if url == "" || token == "" {
		return nil, errors.New("URL and Token are required")
	}

	return &HttpClient{
		client:  &http.Client{},
		url:     url,
		token:   token,
		headers: headers,
	}, nil
}

func (c *HttpClient) Get(path string, params map[string]string, body interface{}) (interface{}, error) {
	if body != nil {
		c.body = body
	}

	url := c.url + path
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("[HTTP_REQUEST:GET] Generate Request Error: " + err.Error())
		return nil, err
	}

	query := request.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	request.URL.RawQuery = query.Encode()

	err = c.do(request)
	if err != nil {
		return nil, err
	}

	return c.body, err
}

func (c *HttpClient) Post(path string, payload, body interface{}) (interface{}, error) {
	if body != nil {
		c.body = body
	}

	url := c.url + path

	requestPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("[HTTP_REQUEST:POST] Marshaling Payload Error: " + err.Error())
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestPayload))
	if err != nil {
		log.Println("[HTTP_REQUEST:POST] Generate Request Error: " + err.Error())
		return nil, err
	}

	err = c.do(request)
	if err != nil {
		return nil, err
	}

	return c.body, err
}

func (c *HttpClient) do(request *http.Request) error {
	for k, v := range c.headers {
		request.Header.Set(k, v)
	}

	response, err := c.client.Do(request)
	r, _ := json.Marshal(response)
	log.Println("[HTTP_REQUEST:DO] Http Response: " + string(r))

	if err != nil {
		log.Println("[HTTP_REQUEST:DO] Request Error: " + err.Error())
		return err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[HTTP_REQUEST:DO] Read body Error: " + err.Error())
		return err
	}

	err = json.Unmarshal(data, c.body)
	if err != nil {
		log.Println("[HTTP_REQUEST:DO] Unmarshaling Error: " + err.Error())
		return err
	}

	return err
}
