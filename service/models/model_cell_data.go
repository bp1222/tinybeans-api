package models

// CellData - Data in a cell
type CellData struct {

	// String, numeric, or boolean value of the cell. If the cell is a formula, this value will be the formula string.
	Value *interface{} `json:"value,omitempty"`

	// String, numeric, or boolean value result value of the cell. If the cell is a formula, this value will be the calculated result.
	CalculatedValue *interface{} `json:"calculated_value,omitempty"`

	Formats Formats `json:"formats,omitempty"`
}
