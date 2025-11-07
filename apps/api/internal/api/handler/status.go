package handler

import (
	"net/http"
	"strconv"

	"paddletraffic/internal/api/response"
	"paddletraffic/internal/service"

	"github.com/go-chi/chi/v5"
)

type StatusHandler struct {
	service *service.StatusService
}

func NewStatusHandler(service *service.StatusService) *StatusHandler {
	return &StatusHandler{service: service}
}

func (h *StatusHandler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/status", func(r chi.Router) {
		r.Get("/", h.GetStatusBatch)
		r.Get("/{courtId}", h.GetStatusByCourtID)
	})
}

func (h *StatusHandler) GetStatusByCourtID(w http.ResponseWriter, r *http.Request) {
	courtIDStr := chi.URLParam(r, "courtId")
	courtID, err := strconv.ParseInt(courtIDStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid courtId parameter")
		return
	}

	status, err := h.service.GetStatusByCourtID(r.Context(), courtID)
	if err != nil {
		response.InternalError(w, "Failed to fetch court status")
		return
	}

	response.OK(w, status)
}

func (h *StatusHandler) GetStatusBatch(w http.ResponseWriter, r *http.Request) {
	courtIDsParam := r.URL.Query()["courtIds"]
	if len(courtIDsParam) == 0 {
		response.BadRequest(w, "courtIds query parameter is required")
		return
	}

	courtIDs := make([]int64, 0, len(courtIDsParam))
	for _, idStr := range courtIDsParam {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.BadRequest(w, "Invalid courtId in courtIds array")
			return
		}
		courtIDs = append(courtIDs, id)
	}

	statuses, err := h.service.GetStatusBatch(r.Context(), courtIDs)
	if err != nil {
		response.InternalError(w, "Failed to fetch court statuses")
		return
	}

	response.OK(w, statuses)
}
