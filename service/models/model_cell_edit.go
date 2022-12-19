package models

// CellEdit - A single cell edit
type CellEdit struct {

	// The row of the cell to edit
	Row int32 `json:"row"`

	// The column of the cell to edit
	Col int32 `json:"col"`

	// String, numeric, or boolean value
	Value *interface{} `json:"value"`
}
