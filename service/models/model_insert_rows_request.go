package models

// InsertRowsRequest - Insert rows into the sheet
type InsertRowsRequest struct {

	// The index to insert the rows
	Index int32 `json:"index"`

	// The number of rows to insert
	Count int32 `json:"count"`

	InheritFrom InheritFrom `json:"inherit_from"`
}
