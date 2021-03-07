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
