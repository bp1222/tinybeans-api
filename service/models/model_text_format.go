package models

// TextFormat - Text formats. Fields that are omitted will be ignored.
type TextFormat struct {

	// Text bold format
	Bold bool `json:"bold,omitempty"`

	// Text italic format
	Italic bool `json:"italic,omitempty"`

	// Text underline format
	Underline bool `json:"underline,omitempty"`

	// Text strikethrough format
	Strikethrough bool `json:"strikethrough,omitempty"`

	// Text font format
	FontFamily string `json:"font_family,omitempty"`

	// Text font size, in points
	FontSize float32 `json:"font_size,omitempty"`

	// A hex color code
	FontColor string `json:"font_color,omitempty"`
}
