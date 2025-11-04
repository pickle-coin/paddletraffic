package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) Create(ctx context.Context, reportCreate dto.ReportCreate) (dto.ReportSummary, error) {
	return s.repo.Create(ctx, reportCreate)
}
