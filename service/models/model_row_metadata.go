package models

// RowMetadata - Metadata about a row
type RowMetadata struct {

	// The height of the row, in points
	Size int32 `json:"size,omitempty"`

	// Whether the row is hidden
	Hidden bool `json:"hidden,omitempty"`

	// Whether the row is filtered
	Filtered bool `json:"filtered,omitempty"`
}
