package models

// UnhideColumnsRequest - Unhide columns in the sheet
type UnhideColumnsRequest struct {

	// The intervals of columns to unhide
	Intervals []Interval `json:"intervals"`
}
