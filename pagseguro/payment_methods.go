package pagseguro

type PaymentMethod struct {
	Type         string `json:"type"`
	Installments int    `json:"installments,omitempty"`
	Capture      bool   `json:"capture,omitempty"`
	Card         `json:"card"`
	Boleto       `json:"boleto"`
}
