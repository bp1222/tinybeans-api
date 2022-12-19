package models

// DeleteRowsRequest - Delete rows from the sheet
type DeleteRowsRequest struct {

	Interval Interval `json:"interval"`

	// Force the deletion of links, xbrl, footnotes, etc
	Force bool `json:"force,omitempty"`
}
