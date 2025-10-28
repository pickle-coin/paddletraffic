package dto

type Paginated[T any] struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Data       []T `json:"data"`
}

type PaginationParams struct {
	Page     int
	PageSize int
}

const (
	DefaultPage     = 1
	DefaultPageSize = 50
	MaxPageSize     = 100
)

func NewPaginationParams(page, pageSize int) PaginationParams {
	if page < 1 {
		page = DefaultPage
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p PaginationParams) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p PaginationParams) Limit() int {
	return p.PageSize
}

func NewPaginated[T any](data []T, page, pageSize, totalItems int) Paginated[T] {
	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages < 0 {
		totalPages = 0
	}

	return Paginated[T]{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       data,
	}
}
