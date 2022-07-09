package pagseguro

// Amount has amount information to be charge
type Amount struct {
	// Value amount to be charge in cents
	Value    int    `json:"value,omitempty"`
	// Currency is ISO three-letter currency code, in capital letters.
	// For now, only the Brazilian Real is supported (“BRL”)
	Currency string `json:"currency,omitempty"`
	// Summary is a summary of charge amounts
	*Summary `json:"summary,omitempty"`
}

type Summary struct {
	// Total value of charge
	Total int `json:"total,omitempty"`
	// Total paid
	Paid int `json:"paid,omitempty"`
	// Total refunded
	Refunded int `json:"refunded,omitempty"`
}