package pagseguro

import (
	"log"
	"os"
	"testing"
)

func getIntegration(t *testing.T) *Integration {
	t.Helper()
	integration, err := NewIntegration(os.Getenv("PAGSEG_URL"), os.Getenv("PAGSEG_TOKEN"))
	if err != nil {
		log.Fatal("Integration error", err.Error())
	}
	return integration
}

func TestIntegration_GenerateBoleto(t *testing.T) {
	integration := getIntegration(t)

	t.Run("SUCCESS", func(t *testing.T) {
		expectedReferenceId := "ex-00483001"

		boleto := NewBoletoCharge(
			expectedReferenceId,
			"Motivo da cobrança",
			&Amount{
				Value:    1000,
				Currency: "BRL",
			}, &Boleto{
				DueDate: "2021-05-08",
				InstructionLines: &BoletoInstructionLines{
					Line1: "Pagamento processado para DESC Fatura",
					Line2: "Via PagSeguro",
				},
				Holder: &Holder{
					Name:  "Waltin Junin",
					TaxID: "88600742072",
					Email: "waltin@junin.com",
					Address: &Address{
						Country:    "Brasil",
						Region:     "São Paulo",
						RegionCode: "SP",
						City:       "São Paulo",
						PostalCode: "01452002",
						Street:     "Avenida Brigadeiro Faria Lima",
						Number:     "1384",
						Locality:   "Pinheiros",
					},
				},
			})

		newCharge, err := integration.GenerateBoleto(boleto)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if newCharge.ID == "" {
			t.Errorf("An ID was expected")
		}

		if newCharge.ReferenceID != expectedReferenceId {
			t.Errorf("Expected: %s, Got: %s", expectedReferenceId, newCharge.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		boleto := NewBoletoCharge("", "", &Amount{}, &Boleto{})

		newCharge, err := integration.GenerateBoleto(boleto)
		if err == nil {
			t.Error("Errors was expected")
		}

		if len(newCharge.ErrorMessages) < 1 {
			t.Errorf("Errors was expected")
		}
	})
}

func TestIntegration_Authorization(t *testing.T) {
	integration := getIntegration(t)

	t.Run("SUCCESS", func(t *testing.T) {
		// TODO: remove soon
		t.Skip()
		expectedReferenceId := "jr-10101"

		card := NewCardCharge(
			expectedReferenceId,
			"Teste",
			1,
			false,
			&Amount{
				Value:    1000,
				Currency: "BRL",
			},
			&Card{
				Number:       "4111111111111111",
				ExpMonth:     "03",
				ExpYear:      "2026",
				SecurityCode: "123",
				Holder: &Holder{
					Name: "Waltin Junin",
				},
			},
		)

		newCharge, err := integration.Authorization(card)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if newCharge.ID == "" {
			t.Errorf("An ID was expected")
		}

		if newCharge.ReferenceID != expectedReferenceId {
			t.Errorf("Expected: %s, Got: %s", expectedReferenceId, newCharge.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		card := NewCardCharge("", "", 0, false, &Amount{}, &Card{})

		newCharge, err := integration.Authorization(card)
		if err == nil {
			t.Error("Errors was expected")
		}

		if len(newCharge.ErrorMessages) < 1 {
			t.Errorf("Errors was expected")
		}
	})
}

func TestIntegration_Capture(t *testing.T) {
	integration := getIntegration(t)

	t.Run("SUCCESS", func(t *testing.T) {
		chargeID := "CHAR_D0292102-5E22-4F5A-9C4F-52C22F9E978B"
		status := "PAID"
		capture := &Capture{
			&Amount{
				Value:    1000,
				Currency: "BRL",
			},
		}

		newCapture, err := integration.Capture(chargeID, capture)
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
		capture, err := integration.Capture("abc", &Capture{})
		if err == nil {
			t.Error("Errors was expected")
		}

		if len(capture.ErrorMessages) < 1 {
			t.Errorf("Errors was expected")
		}
	})
}

func TestIntegration_GetCharge(t *testing.T) {
	integration := getIntegration(t)

	t.Run("SUCCESS", func(t *testing.T) {
		chargeID := "CHAR_D0292102-5E22-4F5A-9C4F-52C22F9E978B"

		charge, err := integration.GetCharge(chargeID)
		if err != nil {
			t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
		}

		if charge.ID != chargeID {
			t.Errorf("Expected: %s, Got: %s", chargeID, charge.ID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		capture, err := integration.GetCharge("CHAR_A0000000-0A00-0A0A-0A0A-00A00A0A000A")
		if err == nil {
			t.Error("Errors was expected")
		}

		if len(capture.ErrorMessages) < 1 {
			t.Errorf("Errors was expected")
		}
	})
}

func TestIntegration_GetChargeByReferenceId(t *testing.T) {
	integration := getIntegration(t)

	t.Run("SUCCESS", func(t *testing.T) {
		referenceID := "jr-00003"
		chargeID := "CHAR_D0292102-5E22-4F5A-9C4F-52C22F9E978B"

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
		_, err := integration.GetChargesByReferenceId("abc")
		if err == nil {
			t.Error("Errors was expected")
		}
	})
}
