package repository

import (
	"context"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"
	"paddletraffic/internal/mapper"
)

type CourtRepository struct {
	queries *db.Queries
}

func NewCourtRepository(queries *db.Queries) *CourtRepository {
	return &CourtRepository{queries: queries}
}

func (r *CourtRepository) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error) {
	params, err := mapper.CourtCreateDTOToParams(courtCreate)
	if err != nil {
		return dto.CourtSummary{}, err
	}

	row, err := r.queries.CreateCourt(ctx, params)
	if err != nil {
		return dto.CourtSummary{}, err
	}

	return mapper.CreateCourtRowToCourtSummary(row)
}

func (r *CourtRepository) GetAll(ctx context.Context, params dto.PaginationParams) (dto.Paginated[dto.CourtSummary], error) {
	totalCount, err := r.queries.CountCourts(ctx)
	if err != nil {
		return dto.Paginated[dto.CourtSummary]{}, err
	}

	rows, err := r.queries.GetAllCourts(ctx, db.GetAllCourtsParams{
		Limit:  int32(params.Limit()),
		Offset: int32(params.Offset()),
	})
	if err != nil {
		return dto.Paginated[dto.CourtSummary]{}, err
	}

	summaries := make([]dto.CourtSummary, 0, len(rows))
	for _, row := range rows {
		summary, err := mapper.GetAllCourtsRowToCourtSummary(row)
		if err != nil {
			return dto.Paginated[dto.CourtSummary]{}, err
		}
		summaries = append(summaries, summary)
	}

	return dto.NewPaginated(summaries, params.Page, params.PageSize, int(totalCount)), nil
}
