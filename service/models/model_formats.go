package models

// Formats - Formats on a cell
type Formats struct {

	TextFormat TextFormat `json:"text_format,omitempty"`

	ValueFormat ValueFormat `json:"value_format,omitempty"`

	CellFormat CellFormat `json:"cell_format,omitempty"`
}
