package pagseg

import (
	"testing"
)

var chargeRequest = ChargeRequest{
	ReferenceID: "76c35c0b-34d5-4ecc-af9d-3b2cecab033c",
	Description: "Testing",
	Amount: amount{
		Value:    1000,
		Currency: "BRL",
	},
	PaymentMethod: paymethod{
		Type:         "CREDIT_CARD",
		Installments: 1,
		Capture:      true,
		Card: card{
			Number:       "4111111111111111",
			ExpMonth:     "03",
			ExpYear:      "2026",
			SecurityCode: "123",
			Holder: struct {
				Name string `json:"name,omitempty"`
			}{
				Name: "Jose da Silva",
			},
		},
	},
}

func TestCreateCharge(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.Pay(chargeRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert ReferenceId", func(t *testing.T) {
		response := paymentHelper()
		if response.ReferenceID != chargeRequest.ReferenceID {
			t.Errorf("It was expected '%v' and got '%v'", chargeRequest.ReferenceID, response.ReferenceID)
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		expectedStatus := "PAID"
		if response.Status != expectedStatus {
			t.Errorf("It was expected '%v' and got '%v'", expectedStatus, response.Status)
		}
	})

	t.Run("Assert Message", func(t *testing.T) {
		response := paymentHelper()
		expectedMessage := "SUCESSO"
		if response.PaymentResponse.Message != expectedMessage {
			t.Errorf("It was expected '%v' and got '%v'", expectedMessage, response.PaymentResponse.Message)
		}
	})
}
