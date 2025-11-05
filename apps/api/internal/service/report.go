package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type reportServiceImpl struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportServiceImpl{repo: repo}
}

func (s *reportServiceImpl) Create(ctx context.Context, reportCreate dto.ReportCreate) (dto.ReportSummary, error) {
	return s.repo.Create(ctx, reportCreate)
}
