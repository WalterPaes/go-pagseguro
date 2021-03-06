package pagseguro

type PaymentMethod struct {
	Type           string `json:"type,omitempty"`
	Installments   int    `json:"installments,omitempty"`
	Capture        bool   `json:"capture"`
	SoftDescriptor string `json:"soft_descriptor,omitempty"`
	*Card          `json:"card,omitempty"`
	*Boleto        `json:"boleto,omitempty"`
}
