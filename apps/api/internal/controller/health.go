package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthController struct {
	pool *pgxpool.Pool
}

func NewHealthController(pool *pgxpool.Pool) *HealthController {
	return &HealthController{pool: pool}
}

func (h *HealthController) RegisterRoutes(r chi.Router) {
	r.Get("/health", h.Health)
	r.Get("/health/ready", h.Ready)
}

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (h *HealthController) Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *HealthController) Ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	response := HealthResponse{
		Status: "ok",
	}

	if err := h.pool.Ping(ctx); err != nil {
		response.Status = "unavailable"
		response.Database = "unreachable"
		response.Error = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Database = "connected"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
