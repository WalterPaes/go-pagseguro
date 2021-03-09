package pagseguro

// Address has data of payment responsible
type Address struct {
	// Country of payment responsible
	Country string `json:"country,omitempty"`
	// Region of payment responsible
	Region string `json:"region,omitempty"`
	// RegionCode of payment responsible
	RegionCode string `json:"region_code,omitempty"`
	// City of payment responsible
	City string `json:"city,omitempty"`
	// PostalCode of payment responsible
	PostalCode string `json:"postal_code,omitempty"`
	// Street of payment responsible
	Street string `json:"street,omitempty"`
	// Number of payment responsible
	Number string `json:"number,omitempty"`
	// Locality of payment responsible
	Locality string `json:"locality,omitempty"`
}
