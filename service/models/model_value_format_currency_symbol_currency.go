package models

// ValueFormatCurrencySymbolCurrency - An ISO currency format
type ValueFormatCurrencySymbolCurrency struct {

	// The ISO currency identifier
	Code string `json:"code,omitempty"`

	// How to display the currency. CODE simply displays the ISO currency code while SYMBOL displays the corresponding currency symbol. For codes where we support two different symbols, SYMBOL and SYMBOL2 display as follows:   | code | SYMBOL | SYMBOL2 |   | ---- | ------ | ------- |   | INR  | ₹      | Rs      |   | RSD  | дин    | din     |   | UAH  | ₴      | грн     | 
	Display string `json:"display,omitempty"`
}
