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

func (r *CourtRepository) GetAll(ctx context.Context) ([]dto.CourtSummary, error) {
	rows, err := r.queries.GetAllCourts(ctx)
	if err != nil {
		return nil, err
	}

	summaries := make([]dto.CourtSummary, 0, len(rows))
	for _, row := range rows {
		summary, err := mapper.GetAllCourtsRowToCourtSummary(row)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
