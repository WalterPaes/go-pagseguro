package pagseg

import (
	"GoPagSeguro/pkg/core/net"
	"bytes"
	"encoding/json"
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
	// Reference: https://dev.pagseguro.uol.com.br/v4.0/reference/cobranca-criando-uma-cobranca#cobrando-cartao-em-um-passo
	CreateCharge(request ChargeRequest) (*ChargeResponse, error)

	// Reference: https://dev.pagseguro.uol.com.br/v4.0/reference/cobranca-capturando-uma-cobranca
	Capture(id string, request AmountRequest) (*ChargeResponse, error)

	// Reference: https://dev.pagseguro.uol.com.br/v4.0/reference/cobranca-consultando-uma-cobranca
	GetCharge(id string) (*ChargeResponse, error)

	// Reference: https://dev.pagseguro.uol.com.br/v4.0/reference/cobranca-reembolsando-uma-cobranca#devolvendo-uma-cobranca-paga-pap
	CancelAndRefund(id string, request AmountRequest) (*ChargeResponse, error)
}

type pagseguro struct {
	HttpRequester net.Requester
	baseUrl       string
	headers       map[string]string
}

// Struct Constructor
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

	response, err := pg.doHttpRequest(http.MethodPost, ChargeEndpoint, payload)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pg pagseguro) Capture(id string, request AmountRequest) (*ChargeResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(CaptureEndpoint, id)
	response, err := pg.doHttpRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pg pagseguro) GetCharge(id string) (*ChargeResponse, error) {
	endpoint := fmt.Sprintf(GetEndpoint, id)
	response, err := pg.doHttpRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pg pagseguro) CancelAndRefund(id string, request AmountRequest) (*ChargeResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(CancelEndpoint, id)
	response, err := pg.doHttpRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pg pagseguro) doHttpRequest(method, endpoint string, payload []byte) (*ChargeResponse, error) {
	req, err := http.NewRequest(
		method,
		pg.baseUrl+endpoint,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	resp, err := pg.HttpRequester.SetHeaders(pg.headers).DoRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := pg.HttpRequester.ReadBody(resp.Body)
	if err != nil {
		return nil, err
	}

	var response *ChargeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}
