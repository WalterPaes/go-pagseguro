package pagseguro

type Response struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Reference string `json:"reference,omitempty"`
}
