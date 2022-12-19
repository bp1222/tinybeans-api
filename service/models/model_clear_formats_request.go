package models

// ClearFormatsRequest - Clear formats from ranges.
type ClearFormatsRequest struct {

	// The ranges to clear formats
	Ranges []Range `json:"ranges"`

	// List of TextFormat fields to clear. Use \"*\" to clear all fields.
	TextFormatFields []string `json:"text_format_fields,omitempty"`

	// List of ValueFormat fields to clear. Use \"*\" to clear all fields.
	ValueFormatFields []string `json:"value_format_fields,omitempty"`

	// List of CellFormat fields to clear. Use \"*\" to clear all fields.
	CellFormatFields []string `json:"cell_format_fields,omitempty"`
}
