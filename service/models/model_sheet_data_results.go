package models

// SheetDataResults - Contains data from a sheet
type SheetDataResults struct {

	Range Range `json:"range,omitempty"`

	// Merged ranges that intersect with the request range
	Merges []Range `json:"merges,omitempty"`

	// Metadata about the rows in the request range
	RowMetadata []RowMetadata `json:"row_metadata,omitempty"`

	// Metadata about the columns in the request range
	ColumnMetadata []ColumnMetadata `json:"column_metadata,omitempty"`

	// Cell data in row-major order
	Cells [][]CellData `json:"cells,omitempty"`
}
