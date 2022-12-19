package models

// HideRowsRequest - Hide rows in the sheet
type HideRowsRequest struct {

	// The intervals of rows to hide
	Intervals []Interval `json:"intervals"`

	// Force the hiding of footnotes
	Force bool `json:"force,omitempty"`
}
