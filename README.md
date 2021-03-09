[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<h1 align="center">
GoPagSeguro
</h1>

<h4 align="center">
  This Lib was developed for implement some features of the new PagSeguro's api using Go Lang
</h4>

## Usage
As a library
```
go get "github.com/WalterPaes/GoPagSeguro/pagseguro"
```
Import in project
```go
import "github.com/WalterPaes/GoPagSeguro/pagseguro"
```

### Create a New Integration
```
integration := pagseguro.NewIntegration(
    {PAGSEG_URL}, 
    {PAGSEG_TOKEN}, 
    {PAGSEG_API_VERSION}
)
```

### Generating a Boleto Charge
```
// Create a boleto charge struct to payload
boleto := pagseguro.NewBoletoCharge(
"ex-10001", 
"Charge reason", 
&pagseguro.Amount{
	Value:    100,
	Currency: "BRL",
}, &pagseguro.Boleto{
    DueDate:          "2021-05-01",
    InstructionLines: &pagseguro.BoletoInstructionLines{
        Line1: "Instruction One",
        Line2: "Instruction Two",
    },
    Holder: &pagseguro.Holder{
        Name:    "Walter Paes",
        TaxID:   "65763916093",
        Email:   "email@email.com",
        Address: &pagseguro.Address{
            Country:    "Brasil",
            Region:     "São Paulo",
            RegionCode: "SP",
            City:       "São Paulo",
            PostalCode: "88025011",
            Street:     "Rua Teste",
            Number:     "123",
            Locality:   "Pinheiros",
        },
    },
}, nil)

// Call function to do request
newCharge, err := integration.GenerateBoleto(boleto)
```

### Create a credit card Authorization
```
// Create a credit card charge struct to payload
card := pagseguro.NewCardCharge(
    "ex-10001",
    "charge reason",
    1,
    false,
    nil,
    &pagseguro.Amount{
        Value:    1000,
        Currency: "BRL",
    },
    &pagseguro.Card{
        Number:       "5306941322840724",
        ExpMonth:     "11",
        ExpYear:      "2022",
        SecurityCode: "123",
        Holder:       &pagseguro.Holder{
            Name:    "Walter Paes",
        },
}, nil)

// Call function to do request
newCharge, err := integration.Authorization(card)
```

### Do a Capture after Pre-Authorization
```
chargeID := "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"
amount := &pagseguro.Amount{
    Value: 1000,
    Currency: "BRL",
}

// Call function to do request
newCapture, err := integration.Capture(chargeID, amount)
```

### Get a charge
```
chargeID := "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"

// Call function to do request
newCapture, err := integration.GetCharge(chargeID)
```

### Refund and Cancel
```
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
