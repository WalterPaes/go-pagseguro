package pagseguro

// Response has response information from payment provider
type Response struct {
	// Code indicating the reason for the payment authorization response
	// List: https://dev.pagseguro.uol.com.br/v4.0/reference/cobranca-providers#section-motivos-de-negada
	Code string `json:"code,omitempty"`
	// Message is a friendly message describing the reason for not approval or authorization the charge
	Message string `json:"message,omitempty"`
	// Reference is NSU of the authorization, if the payment is approved by the Issuer
	Reference string `json:"reference,omitempty"`
}
