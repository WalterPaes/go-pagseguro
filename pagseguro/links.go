package pagseguro

// Links struct containing link information related to the resource
type Links struct {
	Rel   string `json:"rel,omitempty"`
	Href  string `json:"href,omitempty"`
	Media string `json:"media,omitempty"`
	Type  string `json:"type,omitempty"`
}
