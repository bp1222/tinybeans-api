package models

// ClearBordersRequest - Clears borders in ranges
type ClearBordersRequest struct {

	// The ranges to clear borders
	Ranges []Range `json:"ranges"`
}
