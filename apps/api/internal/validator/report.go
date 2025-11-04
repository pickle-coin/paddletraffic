package validator

import (
	"paddletraffic/internal/dto"
)

func ValidateReportCreate(reportCreate dto.ReportCreate) error {
	// if courtCreate.Name == "" {
	// 	return ValidationError{Field: "name", Message: "is required"}
	// }

	// if courtCreate.CourtCount <= 0 {
	// 	return ValidationError{Field: "courtCount", Message: "must be greater than 0"}
	// }

	// if err := ValidateLocation(courtCreate.Location); err != nil {
	// 	return err
	// }

	return nil
}
