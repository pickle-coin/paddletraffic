package controller

import (
	"encoding/json"
	"net/http"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"

	"github.com/go-chi/chi/v5"
)

type CourtController struct {
	service *service.CourtService
}

func NewCourtController(service *service.CourtService) *CourtController {
	return &CourtController{service: service}
}

func (c *CourtController) RegisterRoutes(r chi.Router) {
	r.Route("/v1/courts", func(r chi.Router) {
		r.Get("/", c.GetAll)
		r.Post("/", c.Create)
	})
}

// Create handles POST /v1/courts - Creates a new court with its location
func (c *CourtController) Create(w http.ResponseWriter, r *http.Request) {
	var courtCreate dto.CourtCreate
	if err := json.NewDecoder(r.Body).Decode(&courtCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if courtCreate.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if courtCreate.CourtCount <= 0 {
		http.Error(w, "courtCount must be greater than 0", http.StatusBadRequest)
		return
	}
	if courtCreate.Location.AddressLine == "" {
		http.Error(w, "location.addressLine is required", http.StatusBadRequest)
		return
	}
	if courtCreate.Location.CountryCode == "" {
		http.Error(w, "location.countryCode is required", http.StatusBadRequest)
		return
	}
	if len(courtCreate.Location.CountryCode) != 2 {
		http.Error(w, "location.countryCode must be 2 characters (ISO 3166-1 alpha-2)", http.StatusBadRequest)
		return
	}
	if courtCreate.Location.Timezone == "" {
		http.Error(w, "location.timezone is required", http.StatusBadRequest)
		return
	}

	courtSummary, err := c.service.Create(r.Context(), courtCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courtSummary)
}

func (c *CourtController) GetAll(w http.ResponseWriter, r *http.Request) {
	courts, err := c.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courts)
}
