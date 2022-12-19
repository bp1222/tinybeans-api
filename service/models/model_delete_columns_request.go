package models

// DeleteColumnsRequest - Delete columns from the sheet
type DeleteColumnsRequest struct {

	Interval Interval `json:"interval"`

	// Force the deletion of links, xbrl, footnotes, etc
	Force bool `json:"force,omitempty"`
}
