package pagseguro

// Recurring is a struct that has recurrence information
type Recurring struct {
	// Type indicates if the charge is a recurrence
	Type string `json:"type,omitempty"`
}
