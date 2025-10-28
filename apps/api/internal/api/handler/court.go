package handler

import (
	"paddletraffic/internal/api/response"
	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"
	"paddletraffic/internal/validator"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CourtHandler struct {
	service *service.CourtService
}

func NewCourtHandler(service *service.CourtService) *CourtHandler {
	return &CourtHandler{service: service}
}

func (h *CourtHandler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/courts", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
	})
}

func (h *CourtHandler) Create(w http.ResponseWriter, r *http.Request) {
	courtCreate, err := DecodeJSON[dto.CourtCreate](r)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.ValidateCourtCreate(courtCreate); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	courtSummary, err := h.service.Create(r.Context(), courtCreate)
	if err != nil {
		response.InternalError(w, "Failed to create court")
		return
	}

	response.Created(w, courtSummary)
}

func (h *CourtHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	params := ParsePaginationParams(r)

	paginatedCourts, err := h.service.GetAll(r.Context(), params)
	if err != nil {
		response.InternalError(w, "Failed to fetch courts")
		return
	}

	response.OK(w, paginatedCourts)
}
