package models

// UnmergeRangesRequest - Unmerge merges that intersect the provided ranges
type UnmergeRangesRequest struct {

	// The ranges to unmerge
	Ranges []Range `json:"ranges"`
}
