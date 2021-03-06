package pagseguro

type Amount struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
	Summary  struct {
		Total    int `json:"total"`
		Paid     int `json:"paid"`
		Refunded int `json:"refunded"`
	} `json:"summary"`
}
