package pagseguro

const (
	BOLETO     = "BOLETO"
	CREDITCARD = "CREDIT_CARD"
)

// PaymentMethod struct has payment method information
type PaymentMethod struct {
	// Type indicates the payment method used to charge (CREDIT_CARD, BOLETO)
	Type           string `json:"type,omitempty"`
	// Installments is number os installments
	Installments   int    `json:"installments,omitempty"`
	// Capture indicates if the charge should be just pre-authorized or capture automatically
	Capture        bool   `json:"capture"`
	SoftDescriptor string `json:"soft_descriptor,omitempty"`
	// Card has credit card information
	*Card          `json:"card,omitempty"`
	// Boleto has data to generate boleto
	*Boleto        `json:"boleto,omitempty"`
}
