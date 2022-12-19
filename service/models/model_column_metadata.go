package models

// ColumnMetadata - Metadata about a column
type ColumnMetadata struct {

	// The width of the column, in points
	Size int32 `json:"size,omitempty"`

	// Whether the column is hidden
	Hidden bool `json:"hidden,omitempty"`
}
