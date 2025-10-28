package controller

import (
	"encoding/json"
	"net/http"

	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"
	"paddletraffic/internal/validator"

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

func (c *CourtController) Create(w http.ResponseWriter, r *http.Request) {
	var courtCreate dto.CourtCreate
	if err := json.NewDecoder(r.Body).Decode(&courtCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateCourtCreate(courtCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
