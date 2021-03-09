package pagseguro

// Boleto has data to generate boleto
type Boleto struct {
	// DueDate of boleto
	DueDate string `json:"due_date,omitempty"`
	// InstructionLines of boleto
	InstructionLines *BoletoInstructionLines `json:"instruction_lines,omitempty"`
	// Holder information
	*Holder `json:"holder,omitempty"`
}

// BoletoInstructionLines of boleto
type BoletoInstructionLines struct {
	// Line1 is first line about boleto
	Line1 string `json:"line_1,omitempty"`
	// Line2 is second line about boleto
	Line2 string `json:"line_2,omitempty"`
}

// BoletoCharge is a specific struct to boleto charge
type BoletoCharge struct {
	*Charge
}

func NewBoletoCharge(refId, description string, amount *Amount, boleto *Boleto, notificationUrls []string) *BoletoCharge {
	return &BoletoCharge{&Charge{
		ReferenceID: refId,
		Description: description,
		Amount:      amount,
		PaymentMethod: &PaymentMethod{
			Type:   BOLETO,
			Boleto: boleto,
		},
		NotificationUrls: notificationUrls,
	}}
}
