package controller

import (
	"encoding/json"
	"net/http"

	"paddletraffic/internal/database/generated/db"
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

func (c *CourtController) Create(w http.ResponseWriter, r *http.Request) {
	var params db.CreateCourtParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	court, err := c.service.Create(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(court)
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
