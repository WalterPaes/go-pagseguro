package pagseguro

type Response struct {
	Code          int    `json:"code,omitempty"`
	Message       string `json:"message,omitempty"`
	Reference     string `json:"reference,omitempty"`
	ErrorMessages []struct {
		Code          string `json:"code,omitempty"`
		Description   string `json:"description,omitempty"`
		ParameterName string `json:"parameter_name,omitempty"`
	} `json:"error_messages,omitempty"`
}
