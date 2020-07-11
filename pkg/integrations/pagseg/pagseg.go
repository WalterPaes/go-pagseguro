package pagseg

import (
	"GoPagSeguro/pkg/core/net"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	ChargeEndpoint  = "/charges"
	CaptureEndpoint = ChargeEndpoint + "/%v/capture"
	GetEndpoint     = ChargeEndpoint + "/%v"
	CancelEndpoint  = ChargeEndpoint + "/%v/cancel"
)

type PagSeguro interface {
	CreateCharge(request ChargeRequest) (*ChargeResponse, error)
	Capture(id string, request AmountRequest) (*ChargeResponse, error)
	GetCharge(id string) (*ChargeResponse, error)
	CancelAndRefund(id string, request AmountRequest) (*ChargeResponse, error)
}

type pagseguro struct {
	HttpRequester net.Requester
	baseUrl       string
	headers       map[string]string
}

func NewPagSeguro(requester net.Requester) PagSeguro {
	return &pagseguro{
		HttpRequester: requester,
		baseUrl:       os.Getenv("PAGSEG_URL"),
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + os.Getenv("PAGSEG_TOKEN"),
			"x-api-version": "1.0",
		},
	}
}

func (pg pagseguro) CreateCharge(request ChargeRequest) (*ChargeResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, statusCode, err := pg.doHttpRequest(http.MethodPost, ChargeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusCreated {
		return nil, errors.New(response.PaymentResponse.Message)
	}

	return response, nil
}

func (pg pagseguro) Capture(id string, request AmountRequest) (*ChargeResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(CaptureEndpoint, id)
	response, statusCode, err := pg.doHttpRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusCreated {
		return nil, errors.New(response.PaymentResponse.Message)
	}

	return response, nil
}

func (pg pagseguro) GetCharge(id string) (*ChargeResponse, error) {
	endpoint := fmt.Sprintf(GetEndpoint, id)
	response, statusCode, err := pg.doHttpRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(response.PaymentResponse.Message)
	}

	return response, nil
}

func (pg pagseguro) CancelAndRefund(id string, request AmountRequest) (*ChargeResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(CancelEndpoint, id)
	response, statusCode, err := pg.doHttpRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusCreated {
		return nil, errors.New(response.PaymentResponse.Message)
	}

	return response, nil
}

func (pg pagseguro) doHttpRequest(method, endpoint string, payload []byte) (*ChargeResponse, int, error) {
	statusCode := http.StatusInternalServerError
	req, err := http.NewRequest(
		method,
		pg.baseUrl+endpoint,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, statusCode, err
	}

	resp, err := pg.HttpRequester.SetHeaders(pg.headers).DoRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	body, err := pg.HttpRequester.ReadBody(resp.Body)
	if err != nil {
		return nil, statusCode, err
	}

	var response *ChargeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, statusCode, err
	}

	return response, resp.StatusCode, err
}