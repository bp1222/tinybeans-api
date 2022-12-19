package models

// CellFormatIndent - Indentation of content in the cell
type CellFormatIndent struct {

	// The size of the indent
	Value float32 `json:"value"`

	// The unit of the size
	Unit string `json:"unit"`
}
