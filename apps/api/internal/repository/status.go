package repository

import (
	"context"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"
)

type statusRepositoryImpl struct {
	queries *db.Queries
}

func NewStatusRepository(queries *db.Queries) StatusRepository {
	return &statusRepositoryImpl{queries: queries}
}

func (r *statusRepositoryImpl) GetStatusByCourtID(ctx context.Context, courtID int64) (dto.CourtStatus, error) {
	row, err := r.queries.GetCourtStatus(ctx, courtID)
	if err != nil {
		return dto.CourtStatus{}, err
	}

	return dto.CourtStatus{
		CourtID:        row.CourtID,
		CourtsOccupied: row.CourtsOccupied,
		GroupsWaiting:  row.GroupsWaiting,
		LastReport:     nil, // TODO: Implement from reports table
	}, nil
}

func (r *statusRepositoryImpl) GetStatusBatch(ctx context.Context, courtIDs []int64) ([]dto.CourtStatus, error) {
	rows, err := r.queries.GetCourtStatusBatch(ctx, courtIDs)
	if err != nil {
		return nil, err
	}

	// Build result array with only the statuses that exist
	statuses := make([]dto.CourtStatus, 0, len(rows))
	for _, row := range rows {
		statuses = append(statuses, dto.CourtStatus{
			CourtID:        row.CourtID,
			CourtsOccupied: row.CourtsOccupied,
			GroupsWaiting:  row.GroupsWaiting,
			LastReport:     nil, // TODO: Implement from reports table
		})
	}

	return statuses, nil
}

func (r *statusRepositoryImpl) InsertStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error) {
	row, err := r.queries.InsertCourtStatus(ctx, db.InsertCourtStatusParams{
		CourtID:        courtID,
		CourtsOccupied: courtsOccupied,
		GroupsWaiting:  groupsWaiting,
	})
	if err != nil {
		return dto.CourtStatus{}, err
	}

	return dto.CourtStatus{
		CourtID:        row.CourtID,
		CourtsOccupied: row.CourtsOccupied,
		GroupsWaiting:  row.GroupsWaiting,
		LastReport:     nil, // TODO: Implement from reports table
	}, nil
}

func (r *statusRepositoryImpl) UpdateStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error) {
	row, err := r.queries.UpdateCourtStatus(ctx, db.UpdateCourtStatusParams{
		CourtID:        courtID,
		CourtsOccupied: courtsOccupied,
		GroupsWaiting:  groupsWaiting,
	})
	if err != nil {
		return dto.CourtStatus{}, err
	}

	return dto.CourtStatus{
		CourtID:        row.CourtID,
		CourtsOccupied: row.CourtsOccupied,
		GroupsWaiting:  row.GroupsWaiting,
		LastReport:     nil, // TODO: Implement from reports table
	}, nil
}
