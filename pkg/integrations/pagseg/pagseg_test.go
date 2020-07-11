package pagseg

import (
	"testing"
)

var (
	chargeRequest = ChargeRequest{
		ReferenceID: ReferenceId,
		Description: "Testing",
		Amount: amount{
			Value:    1000,
			Currency: "BRL",
		},
		PaymentMethod: paymethod{
			Type:         CreditCard,
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

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != PaymentId {
			t.Errorf("It was expected '%v' and got '%v'", PaymentId, response.ID)
		}
	})

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
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.Capture(PaymentId, amountRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != PaymentId {
			t.Errorf("It was expected '%v' and got '%v'", PaymentId, response.ID)
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
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.GetCharge(PaymentId)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != PaymentId {
			t.Errorf("It was expected '%v' and got '%v'", PaymentId, response.ID)
		}
	})
}

func TestCancelAndRefund(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMock{})

		response, err := pagseg.CancelAndRefund(PaymentId, amountRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper()
		if response.ID != PaymentId {
			t.Errorf("It was expected '%v' and got '%v'", PaymentId, response.ID)
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		if response.Status != statusCanceled {
			t.Errorf("It was expected '%v' and got '%v'", statusCanceled, response.Status)
		}
	})
}