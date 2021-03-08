package tests

import (
	"GoPagSeguro/pagseguro"
	"encoding/json"
	"errors"
	"strings"
)

var (
	chargeID         = "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"
	referenceID      = "76c35c0b-34d5-4ecc-af9d-3b2cecab033c"
	statusAuthorized = "AUTHORIZED"
	statusPaid       = "PAID"
	statusCanceled   = "CANCELED"

	charge = &pagseguro.Charge{
		ID:          chargeID,
		ReferenceID: referenceID,
		Status:      statusAuthorized,
	}
)

type HttpClientSuccess struct{}

func (c *HttpClientSuccess) Get(path string, _ map[string]string) ([]byte, error) {
	var bodyResponse []byte
	paths := strings.Split(path, "/")

	switch paths[len(paths)-1] {
	case "charges":
		charges := []*pagseguro.Charge{charge}
		bodyResponse, _ = json.Marshal(charges)
	case chargeID:
		bodyResponse, _ = json.Marshal(charge)
	default:
		return nil, errors.New("invalid path")
	}

	return bodyResponse, nil
}

func (c *HttpClientSuccess) Post(path string, _ interface{}) ([]byte, error) {
	var bodyResponse []byte
	paths := strings.Split(path, "/")

	switch paths[len(paths)-1] {
	case "charges":
		bodyResponse, _ = json.Marshal(charge)
	case "capture":
		charge.Status = statusPaid
		bodyResponse, _ = json.Marshal(charge)
	case "cancel":
		charge.Status = statusCanceled
		bodyResponse, _ = json.Marshal(charge)
	default:
		return nil, errors.New("invalid path")
	}

	return bodyResponse, nil
}

type HttpClientFail struct{}

func (c *HttpClientFail) Get(_ string, _ map[string]string) ([]byte, error) {
	return nil, errors.New("an error")
}

func (c *HttpClientFail) Post(_ string, _ interface{}) ([]byte, error) {
	return nil, errors.New("an error")
}
