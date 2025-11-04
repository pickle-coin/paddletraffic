package repository

import (
	"context"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"
	"paddletraffic/internal/mapper"
)

type reportRepositoryImpl struct {
	queries *db.Queries
}

func NewReportRepository(queries *db.Queries) ReportRepository {
	return &reportRepositoryImpl{queries: queries}
}

func (r *reportRepositoryImpl) Create(ctx context.Context, reportCreate dto.ReportCreate) (dto.ReportSummary, error) {
	params, err := mapper.ReportCreateDTOToParams(reportCreate)
	if err != nil {
		return dto.ReportSummary{}, err
	}

	row, err := r.queries.CreateReport(ctx, params)
	if err != nil {
		return dto.ReportSummary{}, err
	}

	return mapper.CreateReportRowToReportSummary(row)
}
