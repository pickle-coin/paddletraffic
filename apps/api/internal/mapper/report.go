package mapper

import (
	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"
)

func ReportCreateDTOToParams(reportDTO dto.ReportCreate) (db.CreateReportParams, error) {
	params := db.CreateReportParams{
		CourtID:        reportDTO.CourtID,
		CourtsOccupied: reportDTO.CourtsOccupied,
		GroupsWaiting:  reportDTO.GroupsWaiting,
	}

	return params, nil
}
