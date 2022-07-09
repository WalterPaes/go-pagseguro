package pagseguro

// Card has credit card information
type Card struct {
	// Number of credit card
	Number string `json:"number,omitempty"`
	// ExpMonth is expiration month of card
	ExpMonth string `json:"exp_month,omitempty"`
	// ExpYear is expiration year of card
	ExpYear string `json:"exp_year,omitempty"`
	// SecurityCode of card
	SecurityCode string `json:"security_code,omitempty"`
	// Holder information
	Holder *Holder `json:"holder,omitempty"`
}

// Holder information
type Holder struct {
	// Name of holder
	Name string `json:"name,omitempty"`
	// TaxID is customer document
	TaxID string `json:"tax_id,omitempty"`
	// Email of customer
	Email string `json:"email,omitempty"`
	// Address has data of payment responsible
	Address *Address `json:"address,omitempty"`
}