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

// Create creates a new court with its location and returns a CourtSummary in a single database call
func (r *CourtRepository) Create(ctx context.Context, courtCreate dto.CourtCreate) (dto.CourtSummary, error) {
	// Map DTO to database params
	params, err := mapper.CourtCreateDTOToParams(courtCreate)
	if err != nil {
		return dto.CourtSummary{}, err
	}

	// Create court and location, get back full data
	row, err := r.queries.CreateCourt(ctx, params)
	if err != nil {
		return dto.CourtSummary{}, err
	}

	// Map the result to CourtSummary
	return mapper.CreateCourtRowToCourtSummary(row)
}

// GetAll retrieves all courts with their location data in a single query
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
