package models

// HideColumnsRequest - Hide columns in the sheet
type HideColumnsRequest struct {

	// The intervals of columns to hide
	Intervals []Interval `json:"intervals"`

	// Force the hiding of footnotes
	Force bool `json:"force,omitempty"`
}
