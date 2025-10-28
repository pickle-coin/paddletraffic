package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	page := dto.DefaultPage
	pageSize := dto.DefaultPageSize

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil {
			page = p
		}
	}

	if pageSizeParam := r.URL.Query().Get("pageSize"); pageSizeParam != "" {
		if ps, err := strconv.Atoi(pageSizeParam); err == nil {
			pageSize = ps
		}
	}

	params := dto.NewPaginationParams(page, pageSize)

	paginatedCourts, err := c.service.GetAll(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedCourts)
}
