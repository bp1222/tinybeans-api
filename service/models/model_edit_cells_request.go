package models

// EditCellsRequest - Edit a list of cells
type EditCellsRequest struct {

	// The cells to edit
	Cells []CellEdit `json:"cells"`
}
