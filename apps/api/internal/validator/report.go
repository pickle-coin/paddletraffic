package validator

import (
	"paddletraffic/internal/dto"
)

func ValidateReportCreate(reportCreate dto.ReportCreate) error {
	if reportCreate.CourtID == nil {
		return ValidationError{Field: "court_id", Message: "must include field"}
	}
	if reportCreate.CourtsOccupied == nil {
		return ValidationError{Field: "courts_occupied", Message: "must include field"}
	}
	if reportCreate.GroupsWaiting == nil {
		return ValidationError{Field: "groups_waiting", Message: "must include field"}
	}

	if *reportCreate.CourtsOccupied < 0 {
		return ValidationError{Field: "courts_occupied", Message: "must be positive"}
	}

	if *reportCreate.GroupsWaiting < 0 {
		return ValidationError{Field: "groups_waiting", Message: "must be positive"}
	}

	// more validation occurs in paddletraffic/internal/queries/reports.sql
	// e.g.
	// assert courts occupied <= total_courts
	// if courts_occupied < total_courts then assert groups_waiting == 0

	return nil
}
