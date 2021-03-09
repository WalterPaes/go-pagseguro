package pagseguro

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	chargesEndpoint = "/charges"
	captureEndpoint = "/capture"
	cancelEndpoint  = "/cancel"
)

type Integration struct {
	Http HttpClient
}

func NewIntegration(url, token, version string) (*Integration, error) {
	httpClient, err := NewHttpClient(
		url,
		map[string]string{
			"Authorization":     token,
			"x-api-version":     version,
			"x-idempotency-key": "",
		})

	if err != nil {
		return nil, err
	}

	return &Integration{
		Http: httpClient,
	}, nil
}

func (i *Integration) GenerateBoleto(boleto *BoletoCharge) (*Charge, error) {
	response, errResponse := i.Http.Post(chargesEndpoint, boleto.Charge)
	return i.parseToCharge("GenerateBoleto", response, errResponse)
}

func (i *Integration) Authorization(card *CardCharge) (*Charge, error) {
	response, errResponse := i.Http.Post(chargesEndpoint, card.Charge)
	return i.parseToCharge("Authorization", response, errResponse)
}

func (i *Integration) Capture(chargeID string, amount *Amount) (*Charge, error) {
	endpoint := fmt.Sprintf("%s/%s%s", chargesEndpoint, chargeID, captureEndpoint)

	data := struct {
		*Amount `json:"amount"`
	}{
		amount,
	}

	response, errResponse := i.Http.Post(endpoint, data)
	return i.parseToCharge("Capture", response, errResponse)
}

func (i *Integration) GetCharge(chargeID string) (*Charge, error) {
	response, errResponse := i.Http.Get(chargesEndpoint+"/"+chargeID, nil)
	return i.parseToCharge("GetCharge", response, errResponse)
}

func (i *Integration) RefundAndCancel(chargeID string, amount *Amount) (*Charge, error) {
	endpoint := fmt.Sprintf("%s/%s%s", chargesEndpoint, chargeID, cancelEndpoint)

	data := struct {
		*Amount `json:"amount"`
	}{
		amount,
	}

	response, errResponse := i.Http.Post(endpoint, data)
	return i.parseToCharge("RefundAndCancel", response, errResponse)
}

func (i *Integration) GetChargesByReferenceId(referenceID string) ([]Charge, error) {
	response, errResponse := i.Http.Get(chargesEndpoint, map[string]string{
		"reference_id": referenceID,
	})

	if errResponse != nil {
		log.Println("[PAGSEG:GetChargesByReferenceId] Response Error: " + errResponse.Error())
		return nil, errResponse
	}

	var c []Charge
	err := json.Unmarshal(response, &c)
	if err != nil {
		log.Println("[PAGSEG:GetChargesByReferenceId] Unmarshaling error: " + err.Error())
		return nil, err
	}

	return c, nil
}

func (i *Integration) parseToCharge(methodName string, response []byte, errResponse error) (*Charge, error) {
	var c *Charge
	err := json.Unmarshal(response, &c)
	if err != nil {
		log.Println("[PAGSEG:" + methodName + "] Unmarshaling error: " + err.Error())
		return nil, err
	}

	if errResponse != nil {
		log.Println("[PAGSEG:" + methodName + "] Error: " + errResponse.Error())
	}

	return c, errResponse
}
