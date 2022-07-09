package tests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/WalterPaes/GoPagSeguro/pagseguro"
	"github.com/golang/mock/gomock"
)

func getIntegration(t *testing.T) (*pagseguro.Integration, *MockApiClient) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockApiClient(ctrl)
	i, _ := pagseguro.NewIntegration(m)
	return i, m
}

func TestIntegration_GenerateBoleto(t *testing.T) {
	integration, mock := getIntegration(t)
	response, _ := json.Marshal(boletoChargeResponse)
	mock.
		EXPECT().
		Post("/charges", boletoCharge.Charge).
		Return(response, nil)

	charge, err := integration.BoletoCharge(boletoCharge)
	if err != nil {
		t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
	}
	
	t.Run("Assert charge ID", func(t *testing.T) {
		if charge.ID != boletoChargeResponse.ID {
			t.Errorf("Was expected '%s', but got '%s'", charge.ID, boletoChargeResponse.ID)
		}
	})

	t.Run("Assert charge ReferenceID", func(t *testing.T) {
		if charge.ReferenceID != boletoChargeResponse.ReferenceID {
			t.Errorf("Was expected '%s', but got '%s'", charge.ReferenceID, boletoChargeResponse.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		mock.
			EXPECT().
			Post("/charges", boletoCharge.Charge).
			Return(nil, errors.New("an error"))

		_, err := integration.BoletoCharge(boletoCharge)
		if err == nil {
			t.Fatal("Errors was expected")
		}
	})
}

func TestIntegration_CardCharge(t *testing.T) {
	integration, mock := getIntegration(t)
	response, _ := json.Marshal(creditCardChargeResponse)
	mock.
		EXPECT().
		Post("/charges", creditCardCharge.Charge).
		Return(response, nil)

	charge, err := integration.CardCharge(creditCardCharge)
	if err != nil {
		t.Fatalf("ERRORS WAS NOT EXPECTED: %s", err.Error())
	}
	
	t.Run("Assert charge ID", func(t *testing.T) {
		if charge.ID != creditCardChargeResponse.ID {
			t.Errorf("Was expected '%s', but got '%s'", charge.ID, creditCardChargeResponse.ID)
		}
	})

	t.Run("Assert charge ReferenceID", func(t *testing.T) {
		if charge.ReferenceID != creditCardChargeResponse.ReferenceID {
			t.Errorf("Was expected '%s', but got '%s'", charge.ReferenceID, creditCardChargeResponse.ReferenceID)
		}
	})

	t.Run("ERROR", func(t *testing.T) {
		mock.
			EXPECT().
			Post("/charges", creditCardCharge.Charge).
			Return(nil, errors.New("an error"))

		_, err := integration.CardCharge(creditCardCharge)
		if err == nil {
			t.Fatal("Errors was expected")
		}
	})
}
