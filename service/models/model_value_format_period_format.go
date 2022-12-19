package models

// ValueFormatPeriodFormat - Options for formatting a duration string. Valid for PERIOD
type ValueFormatPeriodFormat struct {

	// Method of displaying the period value
	Display string `json:"display"`

	// The separator to use between denominations if multiple are displayed
	Separator string `json:"separator,omitempty"`

	// Precision to use when rounding decimal numbers for display. Renders with automatic precision if null.
	Precision *int32 `json:"precision,omitempty"`

	// Render a label after each denomination
	ShowLabels bool `json:"show_labels,omitempty"`

	// Render the numbers as words instead of digits
	ShowNumbersAsWords bool `json:"show_numbers_as_words,omitempty"`

	// Capitalize the first word
	CapitalizeFirstWord bool `json:"capitalize_first_word,omitempty"`
}
