package dto

type ReportCreate struct {
	CourtID        int64 `json:"court_id"`
	CourtsOccupied int32 `json:"courts_occupied"`
	GroupsWaiting  int32 `json:"groups_waiting"`
}

type ReportSummary struct {
	ID             int64 `json:"id"`
	CourtID        int64 `json:"court_id"`
	CourtsOccupied int32 `json:"courts_occupied"`
	GroupsWaiting  int32 `json:"groups_waiting"`
}
