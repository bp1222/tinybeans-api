package models

// ValueFormatCurrencySymbol - Render numbers with a currency symbol. Valid for ACCOUNTING and CURRENCY. Either generic or currency should be set, but not both.
type ValueFormatCurrencySymbol struct {

	// A generic currency symbol
	Generic string `json:"generic,omitempty"`

	Currency ValueFormatCurrencySymbolCurrency `json:"currency,omitempty"`
}
