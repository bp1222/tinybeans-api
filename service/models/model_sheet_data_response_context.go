package models

// SheetDataResponseContext - Contextual information for the Response
type SheetDataResponseContext struct {

	// The document revision of the data returned in this response
	Revision int32 `json:"revision,omitempty"`

	// The current page index of this response
	Page int32 `json:"page,omitempty"`

	// The maximum number of cells returned per page in this response
	PerPage int32 `json:"per_page,omitempty"`

	// The sub-region of the request that this page represents
	PagedRegion string `json:"paged_region,omitempty"`

	// The URL of the next page of results
	NextUrl string `json:"next_url,omitempty"`
}
