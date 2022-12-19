package models

type InternalServerErrorResponse struct {

	// The HTTP status code for the response
	Status int32 `json:"status,omitempty"`

	// A code that may be used to reliably detect a specific error condition
	Reason string `json:"reason,omitempty"`

	// Human readable descriptions of the error condition
	Messages []string `json:"messages,omitempty"`
}
