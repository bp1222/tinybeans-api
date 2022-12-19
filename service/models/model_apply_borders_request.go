package models

// ApplyBordersRequest - Apply borders to ranges
type ApplyBordersRequest struct {

	// The ranges to apply borders
	Ranges []Range `json:"ranges"`

	Top Border `json:"top,omitempty"`

	Bottom Border `json:"bottom,omitempty"`

	Left Border `json:"left,omitempty"`

	Right Border `json:"right,omitempty"`

	InnerHorizontal Border `json:"inner_horizontal,omitempty"`

	InnerVertical Border `json:"inner_vertical,omitempty"`
}
