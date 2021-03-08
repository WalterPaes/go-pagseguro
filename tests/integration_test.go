package tests

import (
	"GoPagSeguro/pagseguro"
	"testing"
)

func getIntegration(t *testing.T, httpClient pagseguro.HttpClient) *pagseguro.Integration {
	t.Helper()
	return &pagseguro.Integration{Http: httpClient}
}

func TestIntegration_GenerateBoleto(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		boleto := pagseguro.NewBoletoCharge(referenceID, "", &pagseguro.Amount{}, &pagseguro.Boleto{})

		newCharge, err := integration.GenerateBoleto(boleto)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if newCharge.ID == "" {
			t.Errorf("An ID was expected")
		}

		if newCharge.ReferenceID != boleto.Charge.ReferenceID {
			t.Errorf("Expected: %s, Got: %s", boleto.Charge.ReferenceID, newCharge.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		_, err := integration.GenerateBoleto(&pagseguro.BoletoCharge{})
		if err == nil {
			t.Fatal("Errors was expected")
		}
	})
}

func TestIntegration_Authorization(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		card := pagseguro.NewCardCharge(
			referenceID,
			"Teste",
			1,
			false,
			nil,
			&pagseguro.Amount{},
			&pagseguro.Card{},
		)

		newCharge, err := integration.Authorization(card)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if newCharge.ID == "" {
			t.Errorf("An ID was expected")
		}

		if newCharge.ReferenceID != card.Charge.ReferenceID {
			t.Errorf("Expected: %s, Got: %s", card.Charge.ReferenceID, newCharge.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		card := pagseguro.NewCardCharge("", "", 0, false, nil, &pagseguro.Amount{}, &pagseguro.Card{})

		_, err := integration.Authorization(card)
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestIntegration_Capture(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		status := "PAID"
		amount := &pagseguro.Amount{}

		newCapture, err := integration.Capture(chargeID, amount)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if newCapture.ID != chargeID {
			t.Errorf("Expected: %s, Got: %s", chargeID, newCapture.ID)
		}

		if newCapture.Status != status {
			t.Errorf("Expected: %s, Got: %s", status, newCapture.Status)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		_, err := integration.Capture("", &pagseguro.Amount{})
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestIntegration_GetCharge(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		charge, err := integration.GetCharge(chargeID)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if charge.ID != chargeID {
			t.Errorf("Expected: %s, Got: %s", chargeID, charge.ID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		_, err := integration.GetCharge("")
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestIntegration_RefundAndCancel(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		status := "CANCELED"
		amount := &pagseguro.Amount{
			Value:    100,
			Currency: "BRL",
		}

		charge, err := integration.RefundAndCancel(chargeID, amount)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if charge.ID != chargeID {
			t.Errorf("Expected: %s, Got: %s", chargeID, charge.ID)
		}

		if charge.Status != status {
			t.Errorf("Expected: %s, Got: %s", status, charge.Status)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		_, err := integration.RefundAndCancel("", &pagseguro.Amount{})
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}

func TestIntegration_GetChargesByReferenceId(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientSuccess{})

		charge, err := integration.GetChargesByReferenceId(referenceID)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if charge[0].ReferenceID != referenceID {
			t.Errorf("Expected: %s, Got: %s", referenceID, charge[0].ReferenceID)
		}

		if charge[0].ID != chargeID {
			t.Errorf("Expected: %s, Got: %s", chargeID, charge[0].ID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		integration := getIntegration(t, &HttpClientFail{})

		_, err := integration.GetChargesByReferenceId("")
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}
