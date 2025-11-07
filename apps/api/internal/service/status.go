package service

import (
	"context"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/repository"
)

type StatusService struct {
	repo repository.StatusRepository
}

func NewStatusService(repo repository.StatusRepository) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) GetStatusByCourtID(ctx context.Context, courtID int64) (dto.CourtStatus, error) {
	return s.repo.GetStatusByCourtID(ctx, courtID)
}

func (s *StatusService) GetStatusBatch(ctx context.Context, courtIDs []int64) ([]dto.CourtStatus, error) {
	return s.repo.GetStatusBatch(ctx, courtIDs)
}

func (s *StatusService) InsertStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error) {
	return s.repo.InsertStatus(ctx, courtID, courtsOccupied, groupsWaiting)
}

func (s *StatusService) UpdateStatus(ctx context.Context, courtID int64, courtsOccupied int32, groupsWaiting int32) (dto.CourtStatus, error) {
	return s.repo.UpdateStatus(ctx, courtID, courtsOccupied, groupsWaiting)
}
