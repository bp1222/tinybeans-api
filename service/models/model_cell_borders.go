package models

// CellBorders - The borders applied to a cell
type CellBorders struct {

	Top Border `json:"top,omitempty"`

	Bottom Border `json:"bottom,omitempty"`

	Left Border `json:"left,omitempty"`

	Right Border `json:"right,omitempty"`
}
