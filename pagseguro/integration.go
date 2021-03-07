package pagseguro

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	chargesEndpoint = "/charges"
	captureEndpoint = "/capture"
	cancelEndpoint  = "/cancel"
)

type Integration struct {
	http HttpClient
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
}

func NewIntegration(url, token string) (*Integration, error) {
	httpClient, err := NewHttpClient(
		url,
		map[string]string{
			"Authorization":     token,
			"x-api-version":     os.Getenv("PAGSEG_API_VERSION"),
			"x-idempotency-key": "",
		})

	if err != nil {
		return nil, err
	}

	return &Integration{
		http: httpClient,
	}, nil
}

func (i *Integration) GenerateBoleto(boleto *BoletoCharge) (*Charge, error) {
	response, errResponse := i.http.Post(chargesEndpoint, boleto.Charge)
	return i.parseToCharge("GenerateBoleto", response, errResponse)
}

func (i *Integration) Authorization(card *CardCharge) (*Charge, error) {
	response, errResponse := i.http.Post(chargesEndpoint, card.Charge)
	return i.parseToCharge("Authorization", response, errResponse)
}

func (i *Integration) Capture(chargeID string, capture *Capture) (*Charge, error) {
	endpoint := fmt.Sprintf("%s/%s%s", chargesEndpoint, chargeID, captureEndpoint)
	response, errResponse := i.http.Post(endpoint, capture)
	return i.parseToCharge("Capture", response, errResponse)
}

func (i *Integration) GetCharge(chargeID string) (*Charge, error) {
	response, errResponse := i.http.Get(chargesEndpoint+"/"+chargeID, nil)
	return i.parseToCharge("GetCharge", response, errResponse)
}

func (i *Integration) GetChargesByReferenceId(referenceID string) ([]Charge, error) {
	response, errResponse := i.http.Get(chargesEndpoint, map[string]string{
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
