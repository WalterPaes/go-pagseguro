package pagseguro

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient interface {
	Get(path string, params map[string]string) ([]byte, error)
	Post(path string, payload interface{}) ([]byte, error)
}

type Client struct {
	client  *http.Client
	url     string
	headers map[string]string
}

func NewHttpClient(url string, headers map[string]string) (*Client, error) {
	if url == "" {
		return nil, errors.New("URL are required")
	}

	return &Client{
		client:  &http.Client{},
		url:     url,
		headers: headers,
	}, nil
}

func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
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

	body, err := c.do(request)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (c *Client) Post(path string, payload interface{}) ([]byte, error) {
	url := c.url + path

	requestPayload, err := json.Marshal(payload)
	log.Println("[HTTP_REQUEST:POST] Payload: " + string(requestPayload))
	if err != nil {
		log.Println("[HTTP_REQUEST:POST] Marshaling Payload Error: " + err.Error())
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestPayload))
	if err != nil {
		log.Println("[HTTP_REQUEST:POST] Generate Request Error: " + err.Error())
		return nil, err
	}

	body, err := c.do(request)
	return body, err
}

func (c *Client) do(request *http.Request) ([]byte, error) {
	for k, v := range c.headers {
		request.Header.Set(k, v)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.client.Do(request)

	if err != nil {
		log.Println("[HTTP_REQUEST:DO] Request Error: " + err.Error())
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	log.Println("[HTTP_REQUEST:DO] Response body: " + string(data))
	if err != nil {
		log.Println("[HTTP_REQUEST:DO] Read body Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode > http.StatusCreated {
		if len(data) > 0 {
			return data, errors.New(response.Status)
		}
		return nil, errors.New(response.Status)
	}

	return data, err
}
