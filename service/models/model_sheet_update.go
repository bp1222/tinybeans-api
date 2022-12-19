package models

type SheetUpdate struct {

	// The list of requests to perform on the sheet
	Requests []SheetUpdateRequest `json:"requests"`
}
