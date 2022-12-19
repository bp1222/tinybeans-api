package models

// ValueFormatPrecision - Precision to use when rounding numbers for display. Valid for AUTOMATIC, ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
type ValueFormatPrecision struct {

	// Render with automatic precision based on the value in the cell
	Auto bool `json:"auto,omitempty"`

	// Explicit precision value to use. Required unless auto is true.
	Value int32 `json:"value,omitempty"`
}
