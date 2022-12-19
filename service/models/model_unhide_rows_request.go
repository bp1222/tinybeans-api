package models

// UnhideRowsRequest - Unhide rows in the sheet
type UnhideRowsRequest struct {

	// The intervals of rows to unhide
	Intervals []Interval `json:"intervals"`
}
