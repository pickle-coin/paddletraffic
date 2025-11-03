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

func (r *reportRepositoryImpl) Create(ctx context.Context, reportCreate dto.CourtStatus) (dto.CourtStatus, error) {
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

func (r *reportRepositoryImpl) GetAll(ctx context.Context, params dto.PaginationParams) (dto.Paginated[dto.ReportSummary], error) {
	totalCount, err := r.queries.CountReports(ctx)
	if err != nil {
		return dto.Paginated[dto.ReportSummary]{}, err
	}

	rows, err := r.queries.GetAllReports(ctx, db.GetAllReportsParams{
		Limit:  int32(params.Limit()),
		Offset: int32(params.Offset()),
	})
	if err != nil {
		return dto.Paginated[dto.ReportSummary]{}, err
	}

	summaries := make([]dto.ReportSummary, 0, len(rows))
	for _, row := range rows {
		summary, err := mapper.GetAllReportsRowToReportSummary(row)
		if err != nil {
			return dto.Paginated[dto.ReportSummary]{}, err
		}
		summaries = append(summaries, summary)
	}

	return dto.NewPaginated(summaries, params.Page, params.PageSize, int(totalCount)), nil
}
