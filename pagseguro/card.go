package pagseguro

type Card struct {
	Number       string  `json:"number,omitempty"`
	ExpMonth     string  `json:"exp_month,omitempty"`
	ExpYear      string  `json:"exp_year,omitempty"`
	SecurityCode string  `json:"security_code,omitempty"`
	Holder       *Holder `json:"holder,omitempty"`
}

type Holder struct {
	Name    string   `json:"name,omitempty"`
	TaxID   string   `json:"tax_id,omitempty"`
	Email   string   `json:"email,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type CardCharge struct {
	Charge
}

func NewCardCharge(refId, description string, installments int, capture bool, amount *Amount, card *Card) *CardCharge {
	return &CardCharge{Charge{
		ReferenceID: refId,
		Description: description,
		Amount:      amount,
		PaymentMethod: &PaymentMethod{
			Type:         CREDITCARD,
			Installments: installments,
			Capture:      capture,
			Card:         card,
		},
	}}
}
