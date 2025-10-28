package mapper

import (
	"fmt"
	"strconv"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/dto"

	"github.com/jackc/pgx/v5/pgtype"
)

func CourtCreateDTOToParams(courtDTO dto.CourtCreate) (db.CreateCourtParams, error) {
	params := db.CreateCourtParams{
		Name:        courtDTO.Name,
		CourtCount:  courtDTO.CourtCount,
		AddressLine: courtDTO.Location.AddressLine,
		CountryCode: courtDTO.Location.CountryCode,
		Timezone:    courtDTO.Location.Timezone,
	}

	lat := pgtype.Numeric{}
	if err := lat.Scan(fmt.Sprintf("%f", courtDTO.Location.Coordinates.Lat)); err != nil {
		return params, fmt.Errorf("invalid latitude: %w", err)
	}
	params.Lat = lat

	lon := pgtype.Numeric{}
	if err := lon.Scan(fmt.Sprintf("%f", courtDTO.Location.Coordinates.Lon)); err != nil {
		return params, fmt.Errorf("invalid longitude: %w", err)
	}
	params.Lon = lon

	if courtDTO.Location.Region != nil {
		params.Region = pgtype.Text{String: *courtDTO.Location.Region, Valid: true}
	}

	if courtDTO.Location.PostalCode != nil {
		params.PostalCode = pgtype.Text{String: *courtDTO.Location.PostalCode, Valid: true}
	}

	if courtDTO.Location.PlaceID != nil {
		params.PlaceID = pgtype.Text{String: *courtDTO.Location.PlaceID, Valid: true}
	}

	return params, nil
}

func CreateCourtRowToCourtSummary(row db.CreateCourtRow) (dto.CourtSummary, error) {
	return buildCourtSummary(
		row.CourtID, row.CourtName, row.CourtCount,
		row.AddressLine, row.CountryCode, row.Timezone,
		row.Lat, row.Lon, row.Region, row.PostalCode, row.PlaceID,
	)
}

func GetAllCourtsRowToCourtSummary(row db.GetAllCourtsRow) (dto.CourtSummary, error) {
	return buildCourtSummary(
		row.CourtID, row.CourtName, row.CourtCount,
		row.AddressLine, row.CountryCode, row.Timezone,
		row.Lat, row.Lon, row.Region, row.PostalCode, row.PlaceID,
	)
}

func buildCourtSummary(
	courtID int64, courtName string, courtCount int32,
	addressLine, countryCode, timezone string,
	lat, lon pgtype.Numeric,
	region, postalCode, placeID pgtype.Text,
) (dto.CourtSummary, error) {
	latFloat, err := numericToFloat64(lat)
	if err != nil {
		return dto.CourtSummary{}, fmt.Errorf("invalid latitude in database: %w", err)
	}

	lonFloat, err := numericToFloat64(lon)
	if err != nil {
		return dto.CourtSummary{}, fmt.Errorf("invalid longitude in database: %w", err)
	}

	summary := dto.CourtSummary{
		ID:         courtID,
		Name:       courtName,
		CourtCount: courtCount,
		Location: dto.Location{
			AddressLine: addressLine,
			CountryCode: countryCode,
			Timezone:    timezone,
			Coordinates: dto.Coordinates{
				Lat: latFloat,
				Lon: lonFloat,
			},
		},
	}

	if region.Valid {
		summary.Location.Region = &region.String
	}

	if postalCode.Valid {
		summary.Location.PostalCode = &postalCode.String
	}

	if placeID.Valid {
		summary.Location.PlaceID = &placeID.String
	}

	return summary, nil
}

func numericToFloat64(n pgtype.Numeric) (float64, error) {
	if !n.Valid {
		return 0, fmt.Errorf("numeric value is null")
	}

	// Get the value as a driver.Value (interface{}) and convert
	val, err := n.Value()
	if err != nil {
		return 0, fmt.Errorf("failed to get numeric value: %w", err)
	}

	// The value is typically returned as a string
	switch v := val.(type) {
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse numeric string: %w", err)
		}
		return f, nil
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("unexpected numeric type: %T", v)
	}
}
