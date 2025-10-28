package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type CourtService struct {
	repo *repository.CourtRepository
}

func NewCourtService(repo *repository.CourtRepository) *CourtService {
	return &CourtService{repo: repo}
}

// Create creates a new court with its location and returns the full CourtSummary in a single database call
func (s *CourtService) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error) {
	return s.repo.Create(ctx, courtCreate)
}

// GetAll returns all courts with their location data
func (s *CourtService) GetAll(ctx context.Context) ([]dto.CourtSummary, error) {
	return s.repo.GetAll(ctx)
}
