package handler

import (
	"context"
	"net/http"
	"time"

	"paddletraffic/internal/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{pool: pool}
}

func (h *HealthHandler) RegisterRoutes(r chi.Router) {
	r.Route("/v1/health", func(r chi.Router) {
		r.Get("/", h.Health)
		r.Get("/ready", h.Ready)
	})
}

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response.OK(w, HealthResponse{
		Status: "ok",
	})
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
		response.JSON(w, http.StatusServiceUnavailable, HealthResponse{
			Status:   "unavailable",
			Database: "unreachable",
			Error:    err.Error(),
		})
		return
	}

	response.OK(w, HealthResponse{
		Status:   "ok",
		Database: "connected",
	})
}
