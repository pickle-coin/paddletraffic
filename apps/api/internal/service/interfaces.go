package service

import (
	"context"
	"paddletraffic/internal/dto"
)

type ReportService interface {
	Create(ctx context.Context, reportCreate dto.ReportCreate) (dto.ReportSummary, error)
}
