package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type CourtService struct {
	repo          repository.CourtRepository
	statusService *StatusService
}

func NewCourtService(repo repository.CourtRepository, statusService *StatusService) *CourtService {
	return &CourtService{
		repo:          repo,
		statusService: statusService,
	}
}

func (s *CourtService) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error) {
	courtSummary, err := s.repo.Create(ctx, courtCreate)
	if err != nil {
		return dto.CourtSummary{}, err
	}

	// Create initial status with 0 values for the new court
	_, err = s.statusService.InsertStatus(ctx, courtSummary.ID, 0, 0)
	if err != nil {
		// TODO: Consider whether we want to rollback the court creation or handle this differently
		return courtSummary, err
	}

	return courtSummary, nil
}

func (s *CourtService) GetAll(ctx context.Context, params dto.PaginationParams) (dto.Paginated[dto.CourtSummary], error) {
	return s.repo.GetAll(ctx, params)
}
