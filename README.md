[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<h1 align="center">
GoPagSeguro
</h1>

<h4 align="center">
  This Lib was developed for implement some features of the new PagSeguro's api using Go Lang
</h4>

## Import and Usage
```
go get github.com/WalterPaes/go-pagseguro
```

### Create a New Integration
```go
config := pagseguro.Config{
	Url: "PAGSEG_URL",
	Token: "PAGSEG_TOKEN",
}
client := pagseguro.NewClient(config)
integration := pagseguro.NewIntegration(client)
```

### Generating a Boleto Charge
```go
// Create a boleto charge struct to payload
boletoCharge := pagseguro.NewBoletoCharge(
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

// Call function to do request
newCharge, err := integration.BoletoCharge(boleto)
```

### Create a credit card Authorization
```go
// Create a credit card charge struct to payload
creditCardCharge = pagseguro.NewCardCharge(
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

// Call function to do request
newCharge, err := integration.CardCharge(card)
```

### Do a Capture after Pre-Authorization
```go
chargeID := "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"
amount := &pagseguro.Amount{
    Value: 1000,
    Currency: "BRL",
}

// Call function to do request
newCapture, err := integration.Capture(chargeID, amount)
```

### Get a charge
```go
chargeID := "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"

// Call function to do request
newCapture, err := integration.GetCharge(chargeID)
```

### Refund and Cancel
```go
chargeID := "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"
amount := &pagseguro.Amount{
    Value:    1000,
    Currency: "BRL",
}

// Call function to do request
charge, err := integration.RefundAndCancel(chargeID, amount)
```

## :rocket: Technologies

This project was developed with the following technologies:

-  [Go](https://golang.org/)
-  [GoLand](https://www.jetbrains.com/go/?gclid=EAIaIQobChMI5-ug_OvG6gIVBgiRCh0GGARZEAAYASAAEgKOSPD_BwE)
-  [PagSeguro v4.0 API](https://dev.pagseguro.uol.com.br/v4.0/reference/nova-plataforma)

Made by [Walter Junior](https://www.linkedin.com/in/walter-paes/)
