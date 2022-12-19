package models

// ResizeRowsToFitRequest - Auto-size rows to fit content
type ResizeRowsToFitRequest struct {

	// The intervals of rows to resize
	Intervals []Interval `json:"intervals"`
}
