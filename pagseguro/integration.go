package pagseguro

import (
	"encoding/json"
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
	//os.Setenv("PAGSEG_URL", "https://sandbox.api.pagseguro.com/")
	//os.Setenv("PAGSEG_TOKEN", "0D09FA5117204ECDB75924FFD20A11CC")
	//os.Setenv("PAGSEG_KEY", "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr+ZqgD892U9/HXsa7XqBZUayPquAfh9xx4iwUbTSUAvTlmiXFQNTp0Bvt/5vK2FhMj39qSv1zi2OuBjvW38q1E374nzx6NNBL5JosV0+SDINTlCG0cmigHuBOyWzYmjgca+mtQu4WczCaApNaSuVqgb8u7Bd9GCOL4YJotvV5+81frlSwQXralhwRzGhj/A57CGPgGKiuPT+AOGmykIGEZsSD9RKkyoKIoc0OS8CPIzdBOtTQCIwrLn2FxI83Clcg55W8gkFSOS6rWNbG5qFZWMll6yl02HtunalHmUlRUL66YeGXdMDC2PuRcmZbGO5a/2tbVppW6mfSWG3NPRpgwIDAQAB")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
}

func NewIntegration(url, token string) (*Integration, error) {
	httpClient, err := NewHttpClient(
		url,
		map[string]string{
			"Authorization":     "Bearer " + token,
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

func (i *Integration) Authorization(charge Charge) (*Charge, error) {
	response, errResponse := i.http.Post(chargesEndpoint, charge)
	if response == nil && errResponse != nil {
		log.Println("[PAGSEG:AUTHORIZATION] Error: " + errResponse.Error())
		return nil, errResponse
	}

	var c *Charge
	err := json.Unmarshal(response, &c)
	if err != nil {
		log.Println("[PAGSEG:AUTHORIZATION] Unmarshaling error: " + err.Error())
		return nil, err
	}

	if errResponse != nil {
		log.Println("[PAGSEG:AUTHORIZATION] Error: " + errResponse.Error())
	}

	return c, errResponse
}

func (i *Integration) Capture(charge Charge) (*Charge, error) {
	payload, _ := json.Marshal(charge)

	response, errResponse := i.http.Post(chargesEndpoint, payload)

	var c *Charge
	err := json.Unmarshal(response, &c)
	if err != nil {
		log.Println("[PAGSEG:CAPTURE] Unmarshaling error: " + err.Error())
		return nil, err
	}

	if errResponse != nil {
		log.Println("[PAGSEG:CAPTURE] Error: " + errResponse.Error())
	}

	return c, errResponse
}

func (i *Integration) GenerateBoleto(charge Charge) (*Charge, error) {
	return nil, nil
}
