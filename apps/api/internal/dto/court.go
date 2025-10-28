package dto

type CourtCreate struct {
	Name       string   `json:"name"`
	CourtCount int32    `json:"courtCount"`
	Location   Location `json:"location"`
}

type Location struct {
	AddressLine string      `json:"addressLine"`
	Region      *string     `json:"region,omitempty"`
	PostalCode  *string     `json:"postalCode,omitempty"`
	CountryCode string      `json:"countryCode"`
	Timezone    string      `json:"timezone"`
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     *string     `json:"placeId,omitempty"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type CourtSummary struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	CourtCount int32    `json:"courtCount"`
	Location   Location `json:"location"`
}
