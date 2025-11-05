package mapper

import (
	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"
)

func ReportCreateDTOToParams(reportDTO dto.ReportCreate) (db.CreateReportParams, error) {
	params := db.CreateReportParams{
		CourtID:        *reportDTO.CourtID,
		CourtsOccupied: *reportDTO.CourtsOccupied,
		GroupsWaiting:  *reportDTO.GroupsWaiting,
	}

	return params, nil
}

func CreateReportRowToReportSummary(row db.Report) (dto.ReportSummary, error) {
	summary := dto.ReportSummary{
		ID:             row.ID,
		CourtID:        row.CourtID,
		CourtsOccupied: row.CourtsOccupied,
		GroupsWaiting:  row.GroupsWaiting,
	}
	return summary, nil
}
