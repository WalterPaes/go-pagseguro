package pagseg

type ChargeRequest struct {
	ReferenceID      string    `json:"reference_id"`
	Description      string    `json:"description"`
	Amount           amount    `json:"amount"`
	PaymentMethod    paymethod `json:"payment_method"`
	NotificationUrls []string  `json:"notification_urls"`
}

type ChargeResponse struct {
	ID               string          `json:"id,omitempty"`
	ReferenceID      string          `json:"reference_id,omitempty"`
	Status           string          `json:"status,omitempty"`
	CreatedAt        string          `json:"created_at,omitempty"`
	PaidAt           string          `json:"paid_at,omitempty"`
	Description      string          `json:"description,omitempty"`
	Amount           amount          `json:"amount,omitempty"`
	PaymentResponse  paymentResponse `json:"payment_response,omitempty"`
	PaymentMethod    paymethod       `json:"payment_method,omitempty"`
	Links            []link          `json:"links,omitempty"`
	NotificationUrls []string   `json:"notification_urls,omitempty"`
}

type AmountRequest struct {
	Amount amount `json:"amount"`
}

type amount struct {
	Value    int     `json:"value"`
	Currency string  `json:"currency,omitempty"`
	Summary  summary `json:"summary,omitempty"`
}

type summary struct {
	Total    int `json:"total,omitempty"`
	Paid     int `json:"paid,omitempty"`
	Refunded int `json:"refunded,omitempty"`
}

type paymethod struct {
	Type         string `json:"type,omitempty"`
	Installments int    `json:"installments,omitempty"`
	Capture      bool   `json:"capture,omitempty"`
	Card         card   `json:"card,omitempty"`
}

type card struct {
	Number       string `json:"number,omitempty"`
	ExpMonth     string `json:"exp_month,omitempty"`
	ExpYear      string `json:"exp_year,omitempty"`
	SecurityCode string `json:"security_code,omitempty"`
	Brand        string `json:"brand,omitempty"`
	FirstDigits  string `json:"first_digits,omitempty"`
	LastDigits   string `json:"last_digits,omitempty"`
	Holder       struct {
		Name string `json:"name,omitempty"`
	} `json:"holder,omitempty"`
}

type link struct {
	Rel   string `json:"rel,omitempty"`
	Href  string `json:"href,omitempty"`
	Media string `json:"media,omitempty"`
	Type  string `json:"type,omitempty"`
}

type paymentResponse struct {
	Code      interface{} `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Reference string `json:"reference,omitempty"`
}