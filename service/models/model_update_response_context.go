package models

// UpdateResponseContext - Contextual information for the Response
type UpdateResponseContext struct {

	// The revision created by the request
	Revision int32 `json:"revision,omitempty"`

	// The url of the async job that was kicked off by the request
	Location string `json:"location,omitempty"`
}
