package models

// InsertColumnsRequest - Insert columns into the sheet
type InsertColumnsRequest struct {

	// The index to insert the columns
	Index int32 `json:"index"`

	// The number of columns to insert
	Count int32 `json:"count"`

	InheritFrom InheritFrom `json:"inherit_from"`
}
