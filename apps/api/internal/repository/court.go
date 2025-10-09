package repository

import (
	"context"

	"paddletraffic/internal/database/generated/db"
)

type CourtRepository struct {
	queries *db.Queries
}

func NewCourtRepository(queries *db.Queries) *CourtRepository {
	return &CourtRepository{queries: queries}
}

func (r *CourtRepository) Create(ctx context.Context, params db.CreateCourtParams) (db.Court, error) {
	return r.queries.CreateCourt(ctx, params)
}

func (r *CourtRepository) GetAll(ctx context.Context) ([]db.Court, error) {
	return r.queries.GetAllCourts(ctx)
}
