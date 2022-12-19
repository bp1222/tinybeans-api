package models

// ApplyFormatsRequest - Apply formats to a list of ranges
type ApplyFormatsRequest struct {

	// The ranges to format
	Ranges []Range `json:"ranges"`

	TextFormat TextFormat `json:"text_format,omitempty"`

	ValueFormat ValueFormat `json:"value_format,omitempty"`

	CellFormat CellFormat `json:"cell_format,omitempty"`
}
