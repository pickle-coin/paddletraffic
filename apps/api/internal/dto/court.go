package dto

// CourtCreate represents the request body for creating a new court
// Matches the OpenAPI CourtCreate schema
type CourtCreate struct {
	Name       string   `json:"name"`
	CourtCount int32    `json:"courtCount"`
	Location   Location `json:"location"`
}

// Location represents a physical place that may host pickleball
// Matches the OpenAPI Location schema
type Location struct {
	AddressLine string      `json:"addressLine"`
	Region      *string     `json:"region,omitempty"`
	PostalCode  *string     `json:"postalCode,omitempty"`
	CountryCode string      `json:"countryCode"`
	Timezone    string      `json:"timezone"`
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     *string     `json:"placeId,omitempty"`
}

// Coordinates represents latitude and longitude
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// CourtSummary represents the response for a created court
// Matches the OpenAPI CourtSummary schema
type CourtSummary struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	CourtCount int32    `json:"courtCount"`
	Location   Location `json:"location"`
}
