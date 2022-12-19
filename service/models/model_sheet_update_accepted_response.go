package models

// SheetUpdateAcceptedResponse - A response object indicating the asynchronous sheet update was accepted
type SheetUpdateAcceptedResponse struct {

	// The HTTP status code for the response
	Status int32 `json:"status,omitempty"`

	// A code that may be used to reliably detect a specific response condition
	Reason string `json:"reason,omitempty"`

	// Human readable descriptions of the response condition
	Messages []string `json:"messages,omitempty"`

	Context UpdateResponseContext `json:"context,omitempty"`
}
