package pagseguro

type Charge struct {
	ID               string `json:"id,omitempty"`
	ReferenceID      string `json:"reference_id,omitempty"`
	Description      string `json:"description,omitempty"`
	Status           string `json:"status,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	Amount           `json:"amount,omitempty"`
	PaymentMethod    `json:"payment_method,omitempty"`
	PaymentResponse  Response `json:"payment_response,omitempty"`
	NotificationUrls []string `json:"notification_urls,omitempty"`
	Links            []Links  `json:"links,omitempty"`
}
