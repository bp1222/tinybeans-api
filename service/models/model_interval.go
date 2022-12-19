package models

// Interval - An interval of rows or columns. If either the start or end is null or omitted, the interval is unbounded in that direction.
type Interval struct {

	// The first index of the interval, inclusive
	Start *int32 `json:"start,omitempty"`

	// The last index of the interval, inclusive
	End *int32 `json:"end,omitempty"`
}
