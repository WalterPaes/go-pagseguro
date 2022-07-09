package pagseguro

// Boleto has data to generate boleto
type Boleto struct {
	// ID of boleto
	ID               string `json:"id,omitempty"`
	// Barcode of boleto
	Barcode          string `json:"barcode,omitempty"`
	// FormattedBarcode of boleto
	FormattedBarcode string `json:"formatted_barcode,omitempty"`
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
