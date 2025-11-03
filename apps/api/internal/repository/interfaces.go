package repository

import (
	"context"

	"paddletraffic/internal/dto"
)

// Making an interface for the repository to make it easier to mock the repository for testing
type CourtRepository interface {
	Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error)
	GetAll(ctx context.Context, params dto.PaginationParams) (dto.Paginated[dto.CourtSummary], error)
}

type ReportRepository interface {
	Create(ctx context.Context, reportCreate dto.ReportCreate) (dto.CourtStatus, error)
}
