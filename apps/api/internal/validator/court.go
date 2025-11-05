package validator

import (
	"fmt"

	"paddletraffic/internal/dto"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateCourtCreate(courtCreate dto.CourtCreate) error {
	if courtCreate.Name == "" {
		return ValidationError{Field: "name", Message: "is required"}
	}

	if courtCreate.CourtCount <= 0 {
		return ValidationError{Field: "courtCount", Message: "must be greater than 0"}
	}

	if err := ValidateLocation(courtCreate.Location); err != nil {
		return err
	}

	return nil
}

func ValidateLocation(location dto.Location) error {
	if location.AddressLine == "" {
		return ValidationError{Field: "location.addressLine", Message: "is required"}
	}

	if location.CountryCode == "" {
		return ValidationError{Field: "location.countryCode", Message: "is required"}
	}

	if len(location.CountryCode) != 2 {
		return ValidationError{Field: "location.countryCode", Message: "must be 2 characters (ISO 3166-1 alpha-2)"}
	}

	if location.Timezone == "" {
		return ValidationError{Field: "location.timezone", Message: "is required"}
	}

	if err := ValidateCoordinates(location.Coordinates); err != nil {
		return err
	}

	return nil
}

func ValidateCoordinates(coords dto.Coordinates) error {
	if coords.Lat < -90 || coords.Lat > 90 {
		return ValidationError{Field: "location.coordinates.lat", Message: "must be between -90 and 90"}
	}

	if coords.Lon < -180 || coords.Lon > 180 {
		return ValidationError{Field: "location.coordinates.lon", Message: "must be between -180 and 180"}
	}

	return nil
}

func ValidateCourtID(id int64) error {
	if id <= 0 {
		return ValidationError{Field: "id", Message: "must be positive"}
	}

	return nil
}
