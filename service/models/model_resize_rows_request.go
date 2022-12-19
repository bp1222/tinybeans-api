package models

// ResizeRowsRequest - Resize rows to the specified size
type ResizeRowsRequest struct {

	// The new size for the rows, in points
	Size int32 `json:"size"`

	// The intervals of rows to resize
	Intervals []Interval `json:"intervals"`
}
