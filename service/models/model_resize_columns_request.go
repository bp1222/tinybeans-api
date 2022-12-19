package models

// ResizeColumnsRequest - Resize columns to the specified size
type ResizeColumnsRequest struct {

	// The new size for the columns, in points
	Size int32 `json:"size"`

	// The intervals of columns to resize
	Intervals []Interval `json:"intervals"`
}
