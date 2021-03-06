package pagseguro

type Boleto struct {
	DueDate          string `json:"due_date"`
	InstructionLines struct {
		Line1 string `json:"line_1"`
		Line2 string `json:"line_2"`
	} `json:"instruction_lines"`
	Holder `json:"holder"`
}
