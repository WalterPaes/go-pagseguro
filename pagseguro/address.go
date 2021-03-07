package pagseguro

type Address struct {
	Country    string `json:"country,omitempty"`
	Region     string `json:"region,omitempty"`
	RegionCode string `json:"region_code,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	Locality   string `json:"locality,omitempty"`
}
