package pagseg

import (
	"GoPagSeguro/pkg/core/net"
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
	paymentHelper := func(requester net.Requester) *ChargeResponse {
		pagseg := NewPagSeguro(requester)

		response, err := pagseg.CreateCharge(chargeRequest)
		if err != nil {
			t.Error(err)
		}
		return response
	}

	t.Run("Assert Id", func(t *testing.T) {
		response := paymentHelper(&HttpRequesterMockSuccess{})
		if response.ID != PaymentId {
			t.Errorf("It was expected '%v' and got '%v'", PaymentId, response.ID)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert ReferenceId", func(t *testing.T) {
		response := paymentHelper(&HttpRequesterMockSuccess{})
		if response.ReferenceID != chargeRequest.ReferenceID {
			t.Errorf("It was expected '%v' and got '%v'", chargeRequest.ReferenceID, response.ReferenceID)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper(&HttpRequesterMockSuccess{})
		if response.Status != statusPaid {
			t.Errorf("It was expected '%v' and got '%v'", statusPaid, response.Status)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert Message", func(t *testing.T) {
		response := paymentHelper(&HttpRequesterMockSuccess{})
		if response.PaymentResponse.Message != successMessage {
			t.Errorf("It was expected '%v' and got '%v'", successMessage, response.PaymentResponse.Message)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Bad Request", func(t *testing.T) {
		response := paymentHelper(&HttpRequesterMockFail{})
		if len(response.ErrorMessages) < 1 {
			t.Error("Errors was expected")
		}
	})
}

func TestCapture(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMockSuccess{})

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

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		if response.Status != statusPaid {
			t.Errorf("It was expected '%v' and got '%v'", statusPaid, response.Status)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert Message", func(t *testing.T) {
		response := paymentHelper()
		if response.PaymentResponse.Message != successMessage {
			t.Errorf("It was expected '%v' and got '%v'", successMessage, response.PaymentResponse.Message)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})
}

func TestGetCharge(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMockSuccess{})

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

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})
}

func TestCancelAndRefund(t *testing.T) {
	paymentHelper := func() *ChargeResponse {
		pagseg := NewPagSeguro(&HttpRequesterMockSuccess{})

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

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})

	t.Run("Assert Status", func(t *testing.T) {
		response := paymentHelper()
		if response.Status != statusCanceled {
			t.Errorf("It was expected '%v' and got '%v'", statusCanceled, response.Status)
		}

		if len(response.ErrorMessages) > 0 {
			t.Error("Errors was not expected")
		}
	})
}