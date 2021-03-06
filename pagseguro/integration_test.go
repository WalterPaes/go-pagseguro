package pagseguro

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestIntegration_Authorization(t *testing.T) {
	integration, err := NewIntegration(os.Getenv("PAGSEG_URL"), os.Getenv("PAGSEG_TOKEN"))
	if err != nil {
		log.Fatal("A", err.Error())
	}
	
	charge := Charge{
		ReferenceID: "ex-00001",
		Description: "Motivo da cobrança",
		Amount: &Amount{
			Value:    1000,
			Currency: "BRL",
		},
		PaymentMethod: &PaymentMethod{
			Type:           "CREDIT_CARD",
			Installments:   1,
			Capture:        false,
			SoftDescriptor: "MyStore",
			Card: &Card{
				Number:       "4111111111111111",
				ExpMonth:     "03",
				ExpYear:      "2026",
				SecurityCode: "123",
				Holder: &Holder{
					Name: "Waltin Junin",
				},
			},
		},
	}

	c, err := integration.Authorization(charge)
	if err != nil {
		log.Fatal("B: ", err.Error())
	}

	a, _ := json.Marshal(c)
	fmt.Println(string(a))
}
