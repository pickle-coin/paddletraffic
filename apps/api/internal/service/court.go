package service

import (
	"context"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/repository"
)

type CourtService struct {
	repo *repository.CourtRepository
}

func NewCourtService(repo *repository.CourtRepository) *CourtService {
	return &CourtService{repo: repo}
}

func (s *CourtService) Create(ctx context.Context, params db.CreateCourtParams) (db.Court, error) {
	return s.repo.Create(ctx, params)
}

func (s *CourtService) GetAll(ctx context.Context) ([]db.Court, error) {
	return s.repo.GetAll(ctx)
}
