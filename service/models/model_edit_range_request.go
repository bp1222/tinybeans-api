package models

// EditRangeRequest - Edit all of the cells in a contiguous range
type EditRangeRequest struct {

	Range Range `json:"range"`

	// Row-major ordered two-dimensional array of cell values
	Values [][]*interface{} `json:"values"`
}
