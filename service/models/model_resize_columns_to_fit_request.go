package models

// ResizeColumnsToFitRequest - Auto-size columns to fit content
type ResizeColumnsToFitRequest struct {

	// The intervals of columns to resize
	Intervals []Interval `json:"intervals"`
}
