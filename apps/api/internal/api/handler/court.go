package handler

import (
	"net/http"
	"paddletraffic/internal/api/response"
	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"
	"paddletraffic/internal/validator"
	"strconv"

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
		r.Delete("/{courtId}", h.Delete)
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

func (h *CourtHandler) Delete(w http.ResponseWriter, r *http.Request) {
	courtIdStr := chi.URLParam(r, "courtId")
	courtId, err := strconv.ParseInt(courtIdStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid court ID")
		return
	}

	if err := validator.ValidateCourtID(courtId); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	if err := h.service.Delete(r.Context(), courtId); err != nil {
		response.InternalError(w, "Failed to delete court")
		return
	}

	response.NoContent(w)
}
