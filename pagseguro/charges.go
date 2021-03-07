package pagseguro

type Charge struct {
	ID               string `json:"id,omitempty"`
	ReferenceID      string `json:"reference_id,omitempty"`
	Description      string `json:"description,omitempty"`
	Status           string `json:"status,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	*Amount          `json:"amount,omitempty"`
	*PaymentMethod   `json:"payment_method,omitempty"`
	PaymentResponse  *Response `json:"payment_response,omitempty"`
	*Recurring       `json:"recurring,omitempty"`
	NotificationUrls []string `json:"notification_urls,omitempty"`
	Links            []Links  `json:"links,omitempty"`
	Message          string   `json:"message,omitempty"`
	Title            string   `json:"title"`
	Detail           string   `json:"detail"`
	ErrorMessages    []struct {
		Code          string `json:"code,omitempty"`
		Description   string `json:"description,omitempty"`
		ParameterName string `json:"parameter_name,omitempty"`
		Message       string `json:"message,omitempty"`
	} `json:"error_messages,omitempty"`
}

type Capture struct {
	*Amount `json:"amount,omitempty"`
}
