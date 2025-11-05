package handler

import (
	"net/http"
	"paddletraffic/internal/api/response"
	"paddletraffic/internal/dto"
	"paddletraffic/internal/service"
	"paddletraffic/internal/validator"

	"github.com/go-chi/chi/v5"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/reports", func(r chi.Router) {
		r.Post("/{courtId}", h.Create)
	})
}

func (h *ReportHandler) Create(w http.ResponseWriter, r *http.Request) {
	reportCreate, err := DecodeJSON[dto.ReportCreate](r)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.ValidateReportCreate(reportCreate); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	reportSummary, err := h.service.Create(r.Context(), reportCreate)
	if err != nil {
		response.InternalError(w, "Failed to create report")
		return
	}

	response.Created(w, reportSummary)
}
