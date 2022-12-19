package models

// SheetUpdatedResponse - A response object indicating the sheet update was successful
type SheetUpdatedResponse struct {

	// The HTTP status code for the response
	Status int32 `json:"status,omitempty"`

	// A code that may be used to reliably detect a specific response condition
	Reason string `json:"reason,omitempty"`

	// Human readable descriptions of the response condition
	Messages []string `json:"messages,omitempty"`

	Context UpdateResponseContext `json:"context,omitempty"`
}
