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

func (s *CourtService) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error) {
	return s.repo.Create(ctx, courtCreate)
}

func (s *CourtService) GetAll(ctx context.Context) ([]dto.CourtSummary, error) {
	return s.repo.GetAll(ctx)
}
