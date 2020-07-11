package pagseg

import (
	"testing"
)

var (
	chargeRequest = ChargeRequest{
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

	amountRequest = AmountRequest{Amount: amount{Value: 1000}}

	statusPaid = "PAID"
	statusCanceled = "CANCELED"
	successMessage = "SUCESSO"
)

func TestCreateCharge(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.CreateCharge(chargeRequest)
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
		if response.Status != statusPaid {
			t.Errorf("It was expected '%v' and got '%v'", statusPaid, response.Status)
		}
	})

	t.Run("Assert Message", func(t *testing.T) {
		response := paymentHelper()
		if response.PaymentResponse.Message != successMessage {
			t.Errorf("It was expected '%v' and got '%v'", successMessage, response.PaymentResponse.Message)
		}
	})
}

func TestCapture(t *testing.T) {
	id := "CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD"
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.Capture(id, amountRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != id {
			t.Errorf("It was expected '%v' and got '%v'", id, response.ID)
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		if response.Status != statusPaid {
			t.Errorf("It was expected '%v' and got '%v'", statusPaid, response.Status)
		}
	})

	t.Run("Assert Message", func(t *testing.T) {
		response := paymentHelper()
		if response.PaymentResponse.Message != successMessage {
			t.Errorf("It was expected '%v' and got '%v'", successMessage, response.PaymentResponse.Message)
		}
	})
}

func TestGetCharge(t *testing.T) {
	id := "CHAR_A024DA52-C821-4A94-816F-803AD5307823"
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.GetCharge(id)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != id {
			t.Errorf("It was expected '%v' and got '%v'", id, response.ID)
		}
	})
}

func TestCancelAndRefund(t *testing.T) {
	id := "CHAR_be4545a8-8e62-4d44-85fa-66ccaf2329af"
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.CancelAndRefund(id, amountRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != id {
			t.Errorf("It was expected '%v' and got '%v'", id, response.ID)
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		if response.Status != statusCanceled {
			t.Errorf("It was expected '%v' and got '%v'", statusCanceled, response.Status)
		}
	})
}