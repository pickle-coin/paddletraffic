package controller

import (
	"encoding/json"
	"net/http"

	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/service"
)

type CourtController struct {
	service *service.CourtService
}

func NewCourtController(service *service.CourtService) *CourtController {
	return &CourtController{service: service}
}

func (c *CourtController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/courts", c.handleCourts)
}

func (c *CourtController) handleCourts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.GetAll(w, r)
	case http.MethodPost:
		c.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
