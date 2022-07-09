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

type Config struct {
	Url, Token string
}

type Integration struct {
	Http ApiClient
}

func NewIntegration(client ApiClient) (*Integration, error) {
	return &Integration{
		Http: client,
	}, nil
}

func (i *Integration) BoletoCharge(boleto *boletoCharge) (*Charge, error) {
	response, errResponse := i.Http.Post(chargesEndpoint, boleto.Charge)
	return i.parseToCharge("GenerateBoleto", response, errResponse)
}

func (i *Integration) CardCharge(card *cardCharge) (*Charge, error) {
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
