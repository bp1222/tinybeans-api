package models

// ValueFormat - Value Formats. Fields that are omitted will be ignored.
type ValueFormat struct {

	// The value format type of the content. Setting this property will clear any other ValueFormat properties that are not valid for the new value format type.
	ValueFormatType string `json:"value_format_type,omitempty"`

	// The scale cell values are entered in. Valid for AUTOMATIC, ACCOUNTING, CURRENCY, and NUMBER.
	EnteredIn string `json:"entered_in,omitempty"`

	// The scale cell values are displayed in. Valid for AUTOMATIC, ACCOUNTING, CURRENCY, and NUMBER.
	ShownIn string `json:"shown_in,omitempty"`

	Precision ValueFormatPrecision `json:"precision,omitempty"`

	CurrencySymbol ValueFormatCurrencySymbol `json:"currency_symbol,omitempty"`

	// Render numbers with a percent symbol. Valid for PERCENT.
	PercentSymbol string `json:"percent_symbol,omitempty"`

	// Where to render the symbol relative to the value. All values valid for ACCOUNTING and CURRENCY. Left values valid for NUMBER. Right values valid for PERCENT.
	SymbolAlign string `json:"symbol_align,omitempty"`

	// Include a leading zero for decimal numbers with no whole number part. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	ShowLeadingZero bool `json:"show_leading_zero,omitempty"`

	// Render the thousands separator. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	ShowThousandsSeparator bool `json:"show_thousands_separator,omitempty"`

	// Render zero as a dash instead of a digit. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	UseDashesForZeros bool `json:"use_dashes_for_zeros,omitempty"`

	// Render parentheses around the number instead of a negative symbol. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	UseParensForNegatives bool `json:"use_parens_for_negatives,omitempty"`

	// Render the number as words instead of digits. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	ShowNumbersAsWords bool `json:"show_numbers_as_words,omitempty"`

	// Capitalize the first word of the numbers as words. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT only when show_numbers_as_words is true.
	CapitalizeFirstWord bool `json:"capitalize_first_word,omitempty"`

	// The value to use for zero when displaying numbers as words. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT only when show_numbers_as_words is true.
	DisplayZeroAs string `json:"display_zero_as,omitempty"`

	// Render the sign on values rounded to zero. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	ShowSignRoundedZero bool `json:"show_sign_rounded_zero,omitempty"`

	// Render the positive sign on numbers greater than zero. Valid for ACCOUNTING, CURRENCY, NUMBER, and PERCENT.
	ShowPositiveSign bool `json:"show_positive_sign,omitempty"`

	// Uppercase all characters in the date string. Valid for DATE.
	DateUppercaseAll bool `json:"date_uppercase_all,omitempty"`

	// Use month abbreviations instead of full month names. Valid for DATE.
	DateAbbreviateMonth bool `json:"date_abbreviate_month,omitempty"`

	// Format to use when rendering the date. Valid for DATE.
	DateFormatString string `json:"date_format_string,omitempty"`

	PeriodFormat ValueFormatPeriodFormat `json:"period_format,omitempty"`

	// Custom prefix value to render in the cell. Valid for ACCOUNTING, CURRENCY, NUMBER, PERCENT, and DATE.
	Prefix string `json:"prefix,omitempty"`

	// Custom suffix value to render in the cell. Valid for ACCOUNTING, CURRENCY, NUMBER, PERCENT, and DATE.
	Suffix string `json:"suffix,omitempty"`
}
