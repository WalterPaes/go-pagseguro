package pagseguro

type Charge struct {
	// ID PagSeguro charge identifier
	ID string `json:"id,omitempty"`
	// ReferenceID is own identifier assigned to charge
	ReferenceID string `json:"reference_id,omitempty"`
	// Description is charge description
	Description string `json:"description,omitempty"`
	// Status is charge status (AUTHORIZED, PAID, WAITING, DECLINED, CANCELED)
	Status string `json:"status,omitempty"`
	// CreatedAt date and time that charge was created
	CreatedAt string `json:"created_at,omitempty"`
	// PaidAt date and time that charge was paid
	PaidAt string `json:"paid_at,omitempty"`
	// Amount has amount information to be charge
	*Amount `json:"amount,omitempty"`
	// PaymentMethod struct has payment method information
	*PaymentMethod `json:"payment_method,omitempty"`
	// PaymentResponse is response of charge
	PaymentResponse *Response `json:"payment_response,omitempty"`
	// Recurring is a struct that has recurrence information
	*Recurring `json:"recurring,omitempty"`
	// NotificationUrls that will be notified of any changes to the charge
	NotificationUrls []string `json:"notification_urls,omitempty"`
	// Links struct containing link information related to the resource
	Links            []Links  `json:"links,omitempty"`
	Message          string   `json:"message,omitempty"`
	Title            string   `json:"title,omitempty"`
	Detail           string   `json:"detail,omitempty"`
	ErrorMessages    []struct {
		Code          string `json:"code,omitempty"`
		Description   string `json:"description,omitempty"`
		ParameterName string `json:"parameter_name,omitempty"`
		Message       string `json:"message,omitempty"`
	} `json:"error_messages,omitempty"`
}

// BoletoCharge is a specific struct to boleto charge
type boletoCharge struct {
	*Charge
}

// NewBoletoCharge is a function to create a instance of BoletoCharge
func NewBoletoCharge(
	referenceId string, 
	description string, 
	currency string,
	amountValue int, 
	boleto *Boleto, 
	notificationUrls... string,
) *boletoCharge {
	boletoCharge := &boletoCharge{}
	boletoCharge.Charge = &Charge{
		ReferenceID: referenceId,
		Description: description,
		Amount:      &Amount{
			Value: amountValue,
			Currency: currency,
		},
		PaymentMethod: &PaymentMethod{
			Type:   BOLETO,
			Boleto: boleto,
		},
		NotificationUrls: notificationUrls,
	} 
	return boletoCharge
}

// CardCharge is a specific struct to card charge
type cardCharge struct {
	*Charge
}

// NewCardCharge is a function to create a instance of CardCharge
func NewCardCharge(
	referenceId string,
	description string,
	currency string,
	amountValue int,
	installments int,
	capture bool,
	softDescriptor string,
	card *Card,
	notificationUrls... string,
) *cardCharge {
	cardCharge := &cardCharge{}
	cardCharge.Charge = &Charge{
		ReferenceID: referenceId,
		Description: description,
		Amount:      &Amount{
			Value: amountValue,
			Currency: currency,
		},
		PaymentMethod: &PaymentMethod{
			Type:         CREDITCARD,
			Installments: installments,
			Capture:      capture,
			SoftDescriptor: softDescriptor,
			Card:         card,
		},
		NotificationUrls: notificationUrls,
	}
	return cardCharge
}
