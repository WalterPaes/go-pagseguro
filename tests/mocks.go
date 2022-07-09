package tests

import "github.com/WalterPaes/GoPagSeguro/pagseguro"

var boletoCharge = pagseguro.NewBoletoCharge(
	"12345",
	"Teste",
	"BRL",
	1000,
	&pagseguro.Boleto{
		DueDate: "2024-12-31",
		InstructionLines: &pagseguro.BoletoInstructionLines{
			Line1: "Pagamento Teste3 processado para DESC Fatura",
			Line2: "Via PagSeguro Teste3",
		},
		Holder: &pagseguro.Holder{
			Name: "Waltin",
			TaxID: "27908347096",
			Email: "walter@teste.com",
			Address: &pagseguro.Address{
				Street: "Rua Teste",
				Number: "524545",
				Locality: "Umarizal",
				City: "Belém",
				Region: "Pará",
				RegionCode: "PA",
				Country: "Brasil",
				PostalCode: "88025010",
			},
		},
	},
)

var boletoChargeResponse = pagseguro.Charge{
	ID: "CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD",
	ReferenceID: "12345",
	PaymentMethod: &pagseguro.PaymentMethod{
		Type: pagseguro.BOLETO,
		Boleto: &pagseguro.Boleto{
			ID: "6EA2EB96-CD24-4956-84AB-F4558B6C2097",
			Barcode: "03399853012970000000200726101017777550000005100",
			FormattedBarcode: "03399.85301 29700.000002 00726.101017 7 77550000005100",
			DueDate: "2024-12-31",
		},
	},
}

var creditCardCharge = pagseguro.NewCardCharge(
	"12345",
	"Teste Card 01",
	"BRL",
	108701,
	2,
	true,
	"My store",
	&pagseguro.Card{
		Number: "4111111111111111",
		ExpMonth: "12",
		ExpYear: "2030",
		SecurityCode: "123",
		Holder: &pagseguro.Holder{
			Name: "Waltin",
		},
	},
)

var creditCardChargeResponse = pagseguro.Charge{
	ID: "CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD",
	ReferenceID: "12345",
	Status: "AUTHORIZED",
	PaymentMethod: &pagseguro.PaymentMethod{
		Type: pagseguro.CREDITCARD,
	},
}