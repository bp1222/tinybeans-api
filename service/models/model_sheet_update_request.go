package models

// SheetUpdateRequest - A request to update a sheet in some way. A single field on the SheetUpdateRequest should be set.
type SheetUpdateRequest struct {

	EditCells EditCellsRequest `json:"edit_cells,omitempty"`

	EditRange EditRangeRequest `json:"edit_range,omitempty"`

	ApplyFormats ApplyFormatsRequest `json:"apply_formats,omitempty"`

	ClearFormats ClearFormatsRequest `json:"clear_formats,omitempty"`

	InsertRows InsertRowsRequest `json:"insert_rows,omitempty"`

	InsertColumns InsertColumnsRequest `json:"insert_columns,omitempty"`

	DeleteRows DeleteRowsRequest `json:"delete_rows,omitempty"`

	DeleteColumns DeleteColumnsRequest `json:"delete_columns,omitempty"`

	HideRows HideRowsRequest `json:"hide_rows,omitempty"`

	HideColumns HideColumnsRequest `json:"hide_columns,omitempty"`

	UnhideRows UnhideRowsRequest `json:"unhide_rows,omitempty"`

	UnhideColumns UnhideColumnsRequest `json:"unhide_columns,omitempty"`

	ResizeRows ResizeRowsRequest `json:"resize_rows,omitempty"`

	ResizeRowsToFit ResizeRowsToFitRequest `json:"resize_rows_to_fit,omitempty"`

	ResizeColumns ResizeColumnsRequest `json:"resize_columns,omitempty"`

	ResizeColumnsToFit ResizeColumnsToFitRequest `json:"resize_columns_to_fit,omitempty"`

	MergeRanges MergeRangesRequest `json:"merge_ranges,omitempty"`

	UnmergeRanges UnmergeRangesRequest `json:"unmerge_ranges,omitempty"`

	ApplyBorders ApplyBordersRequest `json:"apply_borders,omitempty"`

	ClearBorders ClearBordersRequest `json:"clear_borders,omitempty"`
}
