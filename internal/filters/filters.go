package filters

import "math"

type Filters struct {
	Page     int64
	PageSize int64
}

func NewFilters(page, pageSize int64, pageDefault, pageSizeDefault int64) Filters {
	if page == 0 {
		page = pageDefault
	}
	if pageSize == 0 {
		pageSize = pageSizeDefault
	}
	return Filters{
		Page:     page,
		PageSize: pageSize,
	}
}

func (f Filters) Limit() int64 {
	return f.PageSize
}

func (f Filters) Offset() int64 {
	return (f.Page - 1) * f.PageSize
}

// Metadata Define a new Metadata struct for holding the pagination metadata.
type Metadata struct {
	CurrentPage  int64 `json:"current_page,omitempty"`
	PageSize     int64 `json:"page_size,omitempty"`
	FirstPage    int64 `json:"first_page,omitempty"`
	LastPage     int64 `json:"last_page,omitempty"`
	TotalRecords int64 `json:"total_records,omitempty"`
}

func CalculateMetadata(totalRecords, page, pageSize int64) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int64(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
