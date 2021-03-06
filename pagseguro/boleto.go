package pagseguro

type Boleto struct {
	DueDate          string `json:"due_date,omitempty"`
	InstructionLines *struct {
		Line1 string `json:"line_1,omitempty"`
		Line2 string `json:"line_2,omitempty"`
	} `json:"instruction_lines,omitempty"`
	*Holder `json:"holder,omitempty"`
}
