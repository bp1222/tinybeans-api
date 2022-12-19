package models
// InheritFrom : Where to inherit formats from when performing an insertion
type InheritFrom string

// List of InheritFrom
const (
	NONE InheritFrom = "NONE"
	BEFORE InheritFrom = "BEFORE"
	AFTER InheritFrom = "AFTER"
)
