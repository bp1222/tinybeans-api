package models

// Border - The type of border that should be applied
type Border struct {

	// The type of border to apply
	Style string `json:"style"`

	// The thickness of the border, in points. Rounded to the nearest hundredth.
	Weight float32 `json:"weight,omitempty"`

	// A hex color code
	Color string `json:"color,omitempty"`
}
