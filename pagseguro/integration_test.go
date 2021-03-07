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
		expectedReferenceId := "ex-00001"

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
