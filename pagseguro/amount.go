package pagseguro

type Amount struct {
	Value    int    `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
	Summary  *struct {
		Total    int `json:"total,omitempty"`
		Paid     int `json:"paid,omitempty"`
		Refunded int `json:"refunded,omitempty"`
	} `json:"summary,omitempty"`
}
