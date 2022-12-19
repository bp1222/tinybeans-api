package models

// SheetDataResponse - A response object containing the results of a sheet data query
type SheetDataResponse struct {

	// The HTTP status code for the response
	Status int32 `json:"status,omitempty"`

	// A code that may be used to reliably detect a specific response condition
	Reason string `json:"reason,omitempty"`

	// Human readable descriptions of the response condition
	Messages []string `json:"messages,omitempty"`

	Context SheetDataResponseContext `json:"context,omitempty"`

	Results SheetDataResults `json:"results,omitempty"`
}
