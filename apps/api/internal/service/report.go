package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type ReportService struct {
	repo repository.CourtRepository
}

func NewReportService(repo repository.CourtRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.ReportSummary, error) {
	return s.repo.Create(ctx, courtCreate)
}
