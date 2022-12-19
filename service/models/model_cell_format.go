package models

// CellFormat - Cell Formats. Fields that are omitted will be ignored.
type CellFormat struct {

	Indent CellFormatIndent `json:"indent,omitempty"`

	// The horizontal alignment of the content in the cell
	HorizontalAlign string `json:"horizontal_align,omitempty"`

	// The vertical alignment of the content in the cell
	VerticalAlign string `json:"vertical_align,omitempty"`

	// The text orientation
	TextRotation string `json:"text_rotation,omitempty"`

	// A hex color code
	BackgroundColor string `json:"background_color,omitempty"`

	// The leader dot pattern to show on the cell
	LeaderDots string `json:"leader_dots,omitempty"`

	Borders CellBorders `json:"borders,omitempty"`
}
