package pagseguro

type Card struct {
	Number       string  `json:"number,omitempty"`
	ExpMonth     string  `json:"exp_month,omitempty"`
	ExpYear      string  `json:"exp_year,omitempty"`
	SecurityCode string  `json:"security_code,omitempty"`
	Holder       *Holder `json:"holder,omitempty"`
}

type Holder struct {
	Name    string `json:"name,omitempty"`
	TaxID   string `json:"tax_id,omitempty"`
	Email   string `json:"email,omitempty"`
	Address *struct {
		Country    string `json:"country,omitempty"`
		Region     string `json:"region,omitempty"`
		RegionCode string `json:"region_code,omitempty"`
		City       string `json:"city,omitempty"`
		PostalCode string `json:"postal_code,omitempty"`
		Street     string `json:"street,omitempty"`
		Number     string `json:"number,omitempty"`
		Locality   string `json:"locality,omitempty"`
	} `json:"address,omitempty"`
}
