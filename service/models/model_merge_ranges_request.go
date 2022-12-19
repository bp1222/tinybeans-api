package models

// MergeRangesRequest - Merge ranges
type MergeRangesRequest struct {

	// The ranges to merge
	Ranges []Range `json:"ranges"`

	// How cells should be merged
	MergeType string `json:"merge_type,omitempty"`

	// Force the merge through links, xbrl, footnotes, etc
	Force bool `json:"force,omitempty"`
}
