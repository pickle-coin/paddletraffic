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

type StatusRepository interface {
	GetStatusByCourtID(ctx context.Context, courtID int64) (dto.CourtStatus, error)
	GetStatusBatch(ctx context.Context, courtIDs []int64) ([]dto.CourtStatus, error)
	InsertStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error)
	UpdateStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error)
}
