package models

// Range - A range in a sheet. If any field is omitted or null, the range is unbounded in that direction.
type Range struct {

	// The index of the first row of the range, inclusive
	StartRow *int32 `json:"start_row,omitempty"`

	// The index of the last row of the range, inclusive
	StopRow *int32 `json:"stop_row,omitempty"`

	// The index of the first column of the range, inclusive
	StartColumn *int32 `json:"start_column,omitempty"`

	// The index of the last column of the range, inclusive
	StopColumn *int32 `json:"stop_column,omitempty"`
}
