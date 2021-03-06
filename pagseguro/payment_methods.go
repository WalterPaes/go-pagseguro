package pagseguro

type PaymentMethod struct {
	Type         string `json:"type,omitempty"`
	Installments int    `json:"installments,omitempty"`
	Capture      bool   `json:"capture"`
	*Card        `json:"card,omitempty"`
	*Boleto      `json:"boleto,omitempty"`
}
