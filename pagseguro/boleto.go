package pagseguro

type Boleto struct {
	DueDate          string                  `json:"due_date,omitempty"`
	InstructionLines *BoletoInstructionLines `json:"instruction_lines,omitempty"`
	*Holder          `json:"holder,omitempty"`
}

type BoletoInstructionLines struct {
	Line1 string `json:"line_1,omitempty"`
	Line2 string `json:"line_2,omitempty"`
}

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
